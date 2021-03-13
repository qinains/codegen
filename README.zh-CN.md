# codegen是什么
codegen是一款根据sql生成项目代码的工具

## 使用
### 直接使用
```
git clone https://github.com/qinains/codegen.git
cd codegen
cp configs/configs-example.json configs/configs.json
# 编辑 configs/configs.json 中的数据库配置等信息
go run main.go
# 在 dist/ 文件夹中生成对应文件
```

### 在项目中使用
```
go get -u github.com/qinains/codegen
# 在项目中创建如 codegen/ 的文件夹，结构如下
# project
# └── codegen
#     ├── configs/configs.json
#     └── template
#         └── default
#             ├── api
#             ├── ...
#             ...
cd project/codegen
codegen
# 在 dist/ 文件夹中生成对应文件
```

## 功能特点
1. 可对项目文件夹进行配置
2. 可编写模板

## 配置
配置文件位于 configs/configs.json 文件中，其中
```
ProjectName         项目名，英文，比如 example
ProjectDescription  项目描述
ModuleName          模块名，比如 example、github.com/qinains/example
TemplateDir         模板目录，比如 tamplate
DistDir             生成代码目录，比如 dist
DB                  数据库配置 
    DriverName      数据库驱动名，目前只支持mysql
    DataSourceName  数据库连接配置，比如 root:root@tcp(127.0.0.1:3306)/example?charset=utf8mb4&parseTime=True&loc=Local
ReservedWords       语言保留字，如 for,if,break
BreakerWords        注释断开词，取断点前的字符串，如：创建时间，毫秒时间戳 -> 创建时间 ; 名称,英文无空格 -> 名称
InitialismWords     通用字，将转化为大写，比如 id -> ID, ip -> IP
```

## 模板编写
默认模板位于 template/default/ 文件夹中，可通过配置文件中的"TemplateDir"配置项修改为其他值

### 模板引擎说明
参考 [https://golang.google.cn/pkg/text/template/](https://golang.google.cn/pkg/text/template/)

### 可用变量
```
.projectName 项目名，英文
.projectDescription 项目描述
.moduleName 模块名
.table 表
    .tableName 表名，比如 user_log
    .tableComment 表描述，比如 用户日志表
    .columns
        .tableName 表名，比如 user_log
        .columnName 列名，比如 login_name,create_time...
        .isNullable 是否可空
        .dataType 数据类型，比如 int,tinyint,varchar...
        .columnType 列类型
        .columnKey 列key
        .columnComment 列描述，比如 登录名，创建时间
        .extra 额外信息
        .numericPrecision 类型大小
        .numericScale 数字大小
        .datetimePrecision 时间类型大小
        .columnDefault 字段默认值
        .其他键，在comment中定义的<<<{key1:value1,key2:value2,...}>>中的key1,key2...
.tables 表列表
    [].table 以上定义的"表"数组
```

### 文件夹编写说明
因为Windows系统文件名中不能包含`|`字符，当遇到使用过滤器文件名的情况，用`__vertical_bar__`代替`|`

### 可使用过滤器和函数
字符串过滤器
```
Upper	                    都转大写 如：user_id -> USER_ID
UpperFirst                  首字母转大写 如：userId -> UserId
UpperAll                    删除断词符，所有都变为大写 如：user_id -> USERID
Lower	                    都转小写 如：User_ID -> user_id
LowerFirst                  首字母转小写 如：UserId -> userId
LowerAll                    删除断词符，所有都变为小写 如：user_id -> userid

UpperInitialisms            特殊字符如ID、IP都全大写 如：user_id -> user_ID ; userId -> userID
Camel	                    小驼峰写法，特殊字符如ID、IP都全大写 如：user_id -> userID ; id_card_no -> IDCardNo
CamelWithoutInitialisms	    小驼峰写法，特殊字符如ID、IP不都全大写 如：user_id -> userId
Pascal	                    大驼峰写法，特殊字符如ID、IP都全大写 如：user_id -> UserID
PascalWithoutInitialisms    大驼峰写法，特殊字符如ID、IP不都全大写 如：user_id -> UserId

Underscore                  转成下杆线 如：userID -> user_id
Dash	                    转成横杠 如：user_id -> user-id ; userId -> user-id

Title	                    空格间的单词首字母都是大写 如：user_id -> User_id ; user id -> User Id
ToTitle                     都转大写 如：user_id -> USER_ID ; user id -> USER ID

Breaker                     取断点前的字符串, 如：创建时间，毫秒时间戳 -> 创建时间 ; 名称,英文无空格 -> 名称
```

函数
```
Contains str substr         str是否包含substr子字符串
IsGE a b                    a是否大于等于b
IsNotNil a                  a是否不为null
IsNumberDataType str        str对应的sql类型是否是数字型
IsStringDataType str        str对应的sql类型是否是字符串
IsReservedWord              str对应的字符串是否是保留字中的字符串
```

### FAQ
1. ".vue" 文件也有 {{}} 模板变量，如何不解析这些变量
将{{}}添加到{{\`\`}}中，比如将{{item.Name}}修改为{{\`{{item.Name}}\`}}
