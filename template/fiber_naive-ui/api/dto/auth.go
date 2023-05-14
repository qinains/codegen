package dto

import (
	"github.com/spf13/viper"
)

// 验证码
type Captcha struct {
	CaptchaID string `json:"captchaID"`                           //验证码ID
	Data      string `json:"data"`                                //图片经过Base64编码之后的文本
	Code      string `json:"code" validate:"captcha" label:"验证码"` //验证码
}

// 登录实体
type Login struct {
	Captcha
	TenantID  int64  `json:"tenantID" label:"租户ID"`                                  //租户ID
	LoginName string `json:"loginName" validate:"required,min=3,max=64" label:"登录名"` //登录名
	Password  string `json:"password" validate:"required,min=4,max=64" label:"密码"`   //密码
}

// 注册实体
type Register struct {
	Captcha
	TenantID  int64  `json:"tenantID" label:"租户ID"`                                                                                          //租户ID
	Phone     string `json:"phone" validate:"required,dbUnique=user:phone&tenant_id->TenantID,min=11,max=16" label:"手机号"`                    //手机号
	LoginName string `json:"loginName" validate:"required,dbUnique=user:login_name->LoginName&tenant_id->TenantID,min=3,max=64" label:"登录名"` //登录名
	Password  string `json:"password" validate:"required,min=4,max=64" label:"密码"`                                                           //密码
}

// Token
type Token struct {
	AccessToken  string `json:"accessToken"`  //用于访问资源
	RefreshToken string `json:"refreshToken"` //可用来获取新的AccessToken
}

// AuthResp
type AuthResp struct {
	ID        int64  `json:"ID"`                    //用户ID
	TenantID  int64  `json:"tenantID"`              //租户ID
	Phone     string `json:"phone" label:"手机号码"`    // 手机号码
	LoginName string `json:"loginName" label:"登录名"` // 登录名
	Nickname  string `json:"nickname" label:"昵称"`   // 昵称
}

func (AuthResp) TableName() string {
	return viper.GetString("jwt.authTableName")
}

type Permission struct {
	ID    int64  `json:"ID"`               // 权限ID
	Label string `json:"label" label:"标签"` // 标签
	Value string `json:"value" label:"值"`  // 值
}

func (Permission) TableName() string {
	return "menu"
}

// AuthInfoResp
type AuthInfoResp struct {
	ID          int64         `json:"ID"`                                                                                                       //用户ID
	TenantID    int64         `json:"tenantID"`                                                                                                 //租户ID
	Phone       string        `json:"phone" label:"手机号码"`                                                                                       // 手机号码
	LoginName   string        `json:"loginName" label:"登录名"`                                                                                    // 登录名
	Nickname    string        `json:"nickname" label:"昵称"`                                                                                      // 昵称
	Permissions *[]Permission `json:"permissions" gorm:"foreignkey:ID;references:ID;association_autoupdate:false;association_autocreate:false"` //权限
}

func (AuthInfoResp) TableName() string {
	return viper.GetString("jwt.authTableName")
}

// 可添加其他“请求”或“返回”的实体，建议以Req、Resp结尾
