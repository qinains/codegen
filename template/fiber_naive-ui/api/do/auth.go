package do

import "github.com/spf13/viper"

// 登录实体
type Auth struct {
	ID           int64  `gorm:"column:id;type:BIGINT UNSIGNED;primaryKey;autoIncrement;not null;comment:用户ID"`                              //用户ID
	TenantID     int64  `gorm:"column:tenant_id;type:BIGINT UNSIGNED;index;not null;comment:租户ID"`                                          //租户ID
	LoginName    string `json:"LoginName" gorm:"column:login_name;type:VARCHAR(32);index;not null;default:;comment:登录名"`                    // 登录名
	PasswordHash string `json:"PasswordHash" gorm:"column:password_hash;type:VARCHAR(128);not null;default:;comment:密码"`                    // 密码
	Phone        string `json:"Phone" gorm:"column:phone;type:VARCHAR(16);index;not null;default:;comment:手机号码"`                            // 手机号码
	Nickname     string `json:"Nickname" gorm:"column:nickname;type:VARCHAR(255);not null;default:;comment:昵称"`                             // 昵称
	CreateTime   int64  `json:"CreateTime" gorm:"column:create_time;type:BIGINT UNSIGNED;not null;autoCreateTime:milli;comment:创建时间，毫秒时间戳"` // 创建时间，毫秒时间戳
	UpdateTime   int64  `json:"UpdateTime" gorm:"column:update_time;type:BIGINT UNSIGNED;not null;autoUpdateTime:milli;comment:更新时间，毫秒时间戳"` // 更新时间，毫秒时间戳
}

func (Auth) TableName() string {
	return viper.GetString("jwt.authTableName")
}
