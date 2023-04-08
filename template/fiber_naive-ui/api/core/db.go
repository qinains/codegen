package core

import (
	"errors"

	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
	"{{$.moduleName}}/initialize"
)

var DB *gorm.DB

func OnInitDB() error {
	driverName := viper.GetString("db.driverName")
	if driverName == "mysql" {
		var err error
		DB, err = gorm.Open(mysql.New(mysql.Config{
			DSN:                       viper.GetString("db.dataSourceName"), // DSN data source name
			DefaultStringSize:         256,                                  // string 类型字段的默认长度
			DisableDatetimePrecision:  true,                                 // 禁用 datetime 精度，MySQL 5.6 之前的数据库不支持
			DontSupportRenameIndex:    true,                                 // 重命名索引时采用删除并新建的方式，MySQL 5.7 之前的数据库和 MariaDB 不支持重命名索引
			DontSupportRenameColumn:   true,                                 // 用 `change` 重命名列，MySQL 8 之前的数据库和 MariaDB 不支持重命名列
			SkipInitializeWithVersion: false,                                // 根据版本自动配置
		}), &gorm.Config{
			Logger: logger.Default.LogMode(logger.Info),
			NamingStrategy: schema.NamingStrategy{
				SingularTable: viper.GetBool("db.singularTable"), // 是否使用单数表名
			},
		})
		if DB != nil {
			gorm.ErrRecordNotFound = errors.New("未找到数据")
		}
		initialize.InitAutoMigrate(DB)
		return err
	}
	return nil
}
