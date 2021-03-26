package main

import (
	"bytes"
	"context"
	"database/sql"
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"text/template"

	_ "github.com/go-sql-driver/mysql"
)

type Config struct {
	ProjectName        string // 项目名
	ProjectDescription string // 项目描述
	ModuleName         string // 模块名
	TemplateDir        string // 模板路径
	DistDir            string // 生成代码路径
	DB                 struct {
		DriverName     string // 数据库驱动名，目前只支持mysql
		DataSourceName string // 数据库连接配置
	}
	SkinParseBySuffix []string // 文件名包含某个后缀名就不参与模板解析
	ReservedWords     []string // 语言保留字
	BreakerWords      []string // 注释断开词
	InitialismWords   []string // 缩略词
}

var skinParseBySuffix []string
var reservedWords []string
var breakerWords []string
var initialismWords []string

var flagTable = flag.String("table", "", "指定生成的数据库表名")
var flagTruncateDistBeforeGen = flag.Bool("truncateDistBeforeGen", false, "先清空目录，再生成代码")

func main() {
	flag.Parse()
	var config Config
	configName := "configs"
	if os.Getenv("CONFIGS_NAME") != "" {
		configName = os.Getenv("CONFIGS_NAME")
	}
	fmt.Printf("当前环境变量CONFIG_NAME为%s，使用%s.json配置文件\n", configName, configName)
	configName = configName + ".json"

	configDirs := []string{"configs/", "./", "../../configs/"}
	for _, v := range configDirs {
		bs, err := ioutil.ReadFile(v + configName)
		if err != nil {
			continue
		}

		err = json.Unmarshal(bs, &config)
		if err != nil {
			panic(err)
		}
		break
	}
	projectName := config.ProjectName
	projectDescription := config.ProjectDescription
	moduleName := config.ModuleName
	templateDir := config.TemplateDir
	distDir := config.DistDir
	skinParseBySuffix = config.SkinParseBySuffix
	reservedWords = config.ReservedWords
	breakerWords = config.BreakerWords
	initialismWords = config.InitialismWords

	if *flagTruncateDistBeforeGen {
		//fmt.Println("生成代码之前，先清空目录")
		err := os.RemoveAll(distDir)
		if err != nil {
			log.Fatal(err)
		}
	}

	ctx := context.Background()
	db, err := sql.Open(config.DB.DriverName, config.DB.DataSourceName)

	sql := "SELECT TABLE_NAME,TABLE_COMMENT FROM information_schema.TABLES WHERE table_schema=(SELECT DATABASE ()) AND table_type = 'BASE TABLE' order by TABLE_NAME"
	rows, err := db.QueryContext(ctx, sql)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	tables := make([]map[string]interface{}, 0)

	for rows.Next() {
		table := make(map[string]interface{})
		var tableName, tableComment string

		if err := rows.Scan(&tableName, &tableComment); err != nil {
			log.Fatal(err)
		}

		jsonStr := GetBetweenStr(tableComment, "<<<", ">>>")
		_ = json.Unmarshal([]byte(jsonStr), &table)

		table["tableName"] = tableName
		table["tableComment"] = tableComment

		// 指定生成的数据库表名
		if *flagTable != "" && *flagTable != table["tableName"] {
			continue
		}

		tables = append(tables, table)
	}

	tablesStr := make([]string, 0)
	for _, table := range tables {
		tablesStr = append(tablesStr, "\""+table["tableName"].(string)+"\"")
	}
	sql = "SELECT TABLE_NAME,COLUMN_NAME,COLUMN_DEFAULT,IS_NULLABLE,DATA_TYPE,COLUMN_TYPE,COLUMN_KEY,COLUMN_COMMENT,EXTRA,NUMERIC_PRECISION, NUMERIC_SCALE,DATETIME_PRECISION FROM information_schema.COLUMNS WHERE TABLE_NAME in (" + strings.Join(tablesStr, ",") + ") AND table_schema=(SELECT DATABASE ()) ORDER BY TABLE_NAME,ORDINAL_POSITION"
	fmt.Println(sql)

	rows, err = db.QueryContext(ctx, sql)
	if err != nil {
		log.Fatal(err)
	}

	for rows.Next() {
		var tableName, columnName, columnKey, columnComment, extra, columnType, dataType, isNullable string
		var columnDefault, numericPrecision, numericScale, datetimePrecision interface{}
		if err := rows.Scan(&tableName, &columnName, &columnDefault, &isNullable, &dataType, &columnType, &columnKey, &columnComment, &extra, &numericPrecision, &numericScale, &datetimePrecision); err != nil {
			log.Fatal(err)
		}

		for _, table := range tables {
			if table["tableName"] == tableName {
				columns := make(map[string]interface{})
				jsonStr := GetBetweenStr(columnComment, "<<<", ">>>")
				_ = json.Unmarshal([]byte(jsonStr), &columns)

				columns["tableName"] = tableName
				columns["columnName"] = columnName
				columns["isNullable"] = isNullable
				columns["dataType"] = dataType
				columns["columnType"] = columnType
				columns["columnKey"] = columnKey
				columns["columnComment"] = columnComment
				columns["extra"] = extra
				columns["numericPrecision"] = numericPrecision
				columns["numericScale"] = numericScale
				columns["datetimePrecision"] = datetimePrecision
				if columnDefault == nil {
					columns["columnDefault"] = nil
				} else {
					columns["columnDefault"] = string(columnDefault.([]byte))
				}

				if table["columns"] == nil {
					table["columns"] = make([]map[string]interface{}, 0)
				}
				table["columns"] = append(table["columns"].([]map[string]interface{}), columns)

				continue
			}
		}
	}

	if err := rows.Close(); err != nil {
		log.Fatal(err)
	}

	funcMap := template.FuncMap{
		"Upper":      Upper,      // 都转大写 如：user_id -> USER_ID
		"UpperFirst": UpperFirst, // 首字母转大写 如：userId -> UserId
		"UpperAll":   UpperAll,   // 删除断词符，所有都变为大写 如：user_id -> USERID

		"Lower":      Lower,      // 都转小写 如：User_ID -> user_id
		"LowerFirst": LowerFirst, // 首字母转小写 如：UserId -> userId
		"LowerAll":   LowerAll,   // 删除断词符，所有都变为小写 如：user_id -> userid

		"UpperInitialisms":         UpperInitialisms,         // 特殊字符如ID、IP都全大写 如：user_id -> user_ID ; userId -> userID
		"Camel":                    Camel,                    // 小驼峰写法，特殊字符如ID、IP都全大写 如：user_id -> userID ; id_card_no -> IDCardNo
		"CamelWithoutInitialisms":  CamelWithoutInitialisms,  // 小驼峰写法，特殊字符如ID、IP不都全大写 如：user_id -> userId
		"Pascal":                   Pascal,                   // 大驼峰写法，特殊字符如ID、IP都全大写 如：user_id -> UserID
		"PascalWithoutInitialisms": PascalWithoutInitialisms, // 大驼峰写法，特殊字符如ID、IP不都全大写 如：user_id -> UserId

		"Underscore": Underscore, // 转成下杆线 如：userID -> user_id
		"Dash":       Dash,       // 转成横杠 如：user_id -> user-id ; userId -> user-id

		"Title":   strings.Title,   // 空格间的单词首字母都是大写 如：user_id -> User_id ; user id -> User Id
		"ToTitle": strings.ToTitle, // 都转大写 如：user_id -> USER_ID ; user id -> USER ID

		"Breaker": Breaker, // 取断点前的字符串, 如：创建时间，毫秒时间戳 -> 创建时间 ; 名称,英文无空格 -> 名称

		"Add":      Add,
		"Subtract": Subtract,
		"Multiply": Multiply,
		"Divide":   Divide,

		"Contains":         strings.Contains,   // str是否包含substr子字符串
		"ReplaceAll":       strings.ReplaceAll, // 在新的str中查找到所有oldStr，替换为newStr
		"IsGE":             IsGE,               // a是否大于等于b
		"IsNotNil":         IsNotNil,           // a是否不为null
		"IsNumberDataType": IsNumberDataType,   // str对应的sql类型是否是数字型
		"IsStringDataType": IsStringDataType,   // str对应的sql类型是否是字符串型
		"IsReservedWord":   IsReservedWord,     // str对应的字符串是否是保留字中的字符串
	}

	for _, table := range tables {
		err = filepath.Walk(templateDir, func(templatePath string, info os.FileInfo, err error) error {
			fmt.Println("templatePath", templatePath)
			if templatePath == templateDir || info.IsDir() {
				return nil
			}

			outputPath := strings.ReplaceAll(templatePath, "__vertical_bar__", "|")
			outputPath = strings.Replace(outputPath, templateDir, distDir, 1)
			outputPath = strings.Replace(outputPath, ".gotemplate", "", 1)
			outputPath = strings.Replace(outputPath, ".gotmpl", "", 1)
			outputPath = strings.Replace(outputPath, ".gohtml", "", 1)
			t := template.Must(template.New("path").Funcs(funcMap).Parse(outputPath))

			wr := bytes.NewBuffer(nil)
			if err = t.Execute(wr, table); err != nil {
				log.Fatal(err)
				return err
			}
			outputFile := string(wr.Bytes())

			data, err := ioutil.ReadFile(templatePath)
			if err != nil {
				log.Fatal(err)
				return nil
			}

			if err = os.MkdirAll(outputFile[0:strings.LastIndex(outputFile, string(os.PathSeparator))], os.ModePerm); err != nil {
				log.Fatal(err)
				return err
			}

			file, err := os.OpenFile(outputFile, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, os.ModePerm)
			if err != nil {
				log.Fatal(err)
				return err
			}

			// 根据后缀名直接写入，不进行模板解析
			for _, i := range skinParseBySuffix {
				if strings.HasSuffix(templatePath, i) {
					file.Write(data)
					if err := file.Close(); err != nil {
						log.Fatal(err)
					}
					return nil
				}
			}

			t2 := template.Must(template.New("content").Funcs(funcMap).Parse(string(data)))
			if err = t2.Execute(file, map[string]interface{}{"table": table, "tables": tables, "moduleName": moduleName, "projectName": projectName, "projectDescription": projectDescription}); err != nil {
				log.Fatal(err)
			}

			if err := file.Close(); err != nil {
				log.Fatal(err)
			}
			return nil
		})
		if err != nil {
			log.Fatal(err)
		}
	}

	// Rows.Err will report the last error encountered by Rows.Scan.
	if err := rows.Err(); err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%s ", tables)
}

// IsNotNil 判断是否不为空
func IsNotNil(i interface{}) bool {
	return i != nil
}

// IsNumberDataType 判断sql类型是否是数字
func IsNumberDataType(dataType string) bool {
	return strings.Contains(dataType, "int") || strings.Contains(dataType, "decimal") || strings.Contains(dataType, "float") || strings.Contains(dataType, "double") || strings.Contains(dataType, "real") || strings.Contains(dataType, "bit") || strings.Contains(dataType, "serial") || strings.Contains(dataType, "dec") || strings.Contains(dataType, "fixed") || strings.Contains(dataType, "numeric")
}

// IsStringDataType 判断sql类型是否是字符串
func IsStringDataType(dataType string) bool {
	return strings.Contains(dataType, "varchar") || strings.Contains(dataType, "char") || strings.Contains(dataType, "text")
}

// IsGE 是否大于某个值
func IsGE(i interface{}, len int) bool {
	switch t := i.(type) {
	case []uint8:
		l, _ := strconv.Atoi(string(t))
		return l >= len
	case string:
		l, _ := strconv.Atoi(t)
		return l >= len
	case int:
		return t >= len
	default:
		return false
	}
}

// IsReservedWord 是否是保留字
func IsReservedWord(s string) bool {
	for _, i := range reservedWords {
		if s == i {
			return true
		}
	}
	return false
}

// Breaker 取断点前的字符串, 如：创建时间，毫秒时间戳 -> 创建时间 ; 名称,英文无空格 -> 名称
func Breaker(s string) string {
	for _, word := range breakerWords {
		if index := strings.Index(s, word); index > -1 {
			s = s[0:index]
		}

	}
	return s
}

func Underscore(s string) string {
	data := make([]byte, 0, len(s)*2)
	j := false
	num := len(s)
	for i := 0; i < num; i++ {
		d := s[i]
		if i > 0 && d >= 'A' && d <= 'Z' && j {
			data = append(data, '_')
		}
		if d != '_' {
			j = true
		}
		data = append(data, d)
	}
	return strings.ToLower(string(data[:]))
}

func Camel(s string) string {
	return UpperInitialisms(LowerFirst(PascalWithoutInitialisms(s)))
}

func CamelWithoutInitialisms(s string) string {
	return LowerFirst(PascalWithoutInitialisms(s))
}

func Pascal(s string) string {
	return UpperInitialisms(PascalWithoutInitialisms(s))
}

func PascalWithoutInitialisms(s string) string {
	if s == "" {
		return ""
	}
	data := make([]byte, 0, len(s))
	j := false
	k := false
	num := len(s) - 1
	for i := 0; i <= num; i++ {
		d := s[i]
		if k == false && d >= 'A' && d <= 'Z' {
			k = true
		}
		if d >= 'a' && d <= 'z' && (j || k == false) {
			d = d - 32
			j = false
			k = true
		}
		if k && d == '_' && num > i && s[i+1] >= 'a' && s[i+1] <= 'z' {
			j = true
			continue
		}
		data = append(data, d)
	}
	return string(data[:])
}

// hidden -> hidden; ip_hidden -> IP_hidden; ipHidden -> IPHidden; hiddenId -> hiddenID; ip_id -> IP_ID; ipId -> IPID;
func UpperInitialisms(s string) string {
	if s == "" {
		return ""
	}

	upperS := strings.ToUpper(s)
	newS := s

	maxLength := len(s)
	current := 0
	for _, i := range initialismWords {
		sIndex := strings.Index(upperS, i)
		for sIndex > -1 {
			current = sIndex + len(i)
			if (sIndex > 0 && s[sIndex-1] >= 'A' && s[sIndex-1] <= 'Z') || (current < len(s) && s[current] >= 'a' && s[current] <= 'z') {
				sIndex = strings.Index(upperS[current:maxLength], i)
				if sIndex > -1 {
					sIndex = current + sIndex
				}
				continue
			}
			newS = string([]byte(newS)[0:sIndex]) + i + string([]byte(newS)[current:maxLength])
			sIndex = strings.Index(upperS[current:maxLength], i)
			if sIndex > -1 {
				sIndex = current + sIndex
			}
		}
	}

	return newS
}

func Dash(s string) string {
	return strings.ReplaceAll(Underscore(s), "_", "-")
}
func Humanize(s string) string {
	return ""
}

func Upper(s string) string {
	return strings.ToUpper(s)
}

func UpperAll(s string) string {
	return Upper(Pascal(s))
}

func Lower(s string) string {
	return strings.ToLower(s)
}

func LowerAll(s string) string {
	return Lower(Pascal(s))
}

func UpperFirst(s string) string {
	return strings.ToUpper(string(s[0])) + string(s[1:])
}

func LowerFirst(s string) string {
	return strings.ToLower(string(s[0])) + string(s[1:])
}

// Pluralize returns string s in plural form.
func Pluralize(s string) string {
	return ""
}

// Singularize returns string s in singular form.
func Singularize(s string) string {
	return ""
}

func IsUpper(c byte) bool {
	return 'A' <= c && c <= 'Z'
}

func IsDigit(c byte) bool {
	return '0' <= c && c <= '9'
}

func GetBetweenStr(str, start, end string) string {
	n := strings.Index(str, start)
	if n == -1 {
		n = 0
	} else {
		n = n + len(start)
	}
	str = string([]byte(str)[n:])
	m := strings.LastIndex(str, end)
	if m == -1 {
		m = len(str)
	}
	str = string([]byte(str)[:m])
	return str
}

func Add(a int, b int) int {
	return a + b
}
func Subtract(a int, b int) int {
	return a - b
}
func Multiply(a int, b int) int {
	return a * b
}

func Divide(a int, b int) int {
	return a / b
}
