package dto

// 验证码
type Captcha struct {
	CaptchaID string //验证码ID
	Data      string //图片经过Base64编码之后的文本
	Code      string //验证码
}

// 登录实体
type Login struct {
	Captcha
	TenantID  int64  `label:"租户ID"`                                 //租户ID
	LoginName string `validate:"required,min=3,max=64" label:"登录名"` //登录名
	Password  string `validate:"required,min=4,max=64" label:"密码"`  //密码
}

// 注册实体
type Register struct {
	Captcha
	TenantID  int64  `label:"租户ID"`                                  //租户ID
	Phone     string `validate:"required,min=11,max=16" label:"手机号"` //手机号
	LoginName string `validate:"required,min=3,max=64" label:"登录名"`  //登录名
	Password  string `validate:"required,min=4,max=64" label:"密码"`   //密码
}

// Token
type Token struct {
	AccessToken  string //用于访问资源
	RefreshToken string //可用来获取新的AccessToken
}
