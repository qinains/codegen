package core

import (
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

// Code 参考HTTP状态码
const (
	CodeOK                = 200 //正常返回
	CodeErrorForm         = 400 //表单/业务错误。如果是业务异常，一般是msg显示异常值，data为null。如果是表单异常，一般msg为空字符串，data包含异常项
	CodeErrorUnauthorized = 401 //当前请求需要用户验证，之后一般需要用户去登录
	CodeErrorForbidden    = 403 //已进行身份验证，但无权限，一般需要用户拥有该权限
	CodeErrorSystem       = 500 //系统错误
)

type Response struct {
	Code int         `json:"code"` // 错误码：200正常,400表单/业务错误(如果是业务异常，一般是msg显示异常值，data为null。如果是表单异常，一般msg为空字符串，data包含异常项),401需要用户验证,403无权限,500系统异常
	Msg  string      `json:"msg"`  // 错误信息
	Data interface{} `json:"data"` // 正常返回内容
}

type ErrorField struct {
	Field string `json:"field"` //错误字段
	Tag   string `json:"tag"`   //错误标记
	Error string `json:"error"` //错误内容
}

type Items struct {
	Total int64       `json:"total"` // 总页码
	List  interface{} `json:"list"`  // 项目列表
}

func Resp(c *fiber.Ctx, code int, msg string, data interface{}) error {
	return c.JSON(Response{Code: code, Msg: msg, Data: data})
}

func OK(c *fiber.Ctx, data interface{}) error {
	return c.JSON(Response{Code: CodeOK, Msg: "", Data: data})
}

func OKMsg(c *fiber.Ctx, msg string) error {
	return c.JSON(Response{Code: CodeOK, Msg: msg})
}

func Error(c *fiber.Ctx, msg string, data interface{}) error {
	return c.JSON(Response{Code: CodeErrorForm, Msg: msg, Data: data})
}

func ErrorBusiness(c *fiber.Ctx, msg string) error {
	return c.JSON(Response{Code: CodeErrorForm, Msg: msg})
}

func ErrorForm(c *fiber.Ctx, data interface{}) error {
	return c.JSON(Response{Code: CodeErrorForm, Data: data})
}

func ErrorFormValidationErrors(c *fiber.Ctx, validate *validator.Validate, err error) error {
	trans := GetTranslator(validate, c.Get("accept-language"))
	var errors []*ErrorField
	for _, err := range err.(validator.ValidationErrors) {
		var element ErrorField
		element.Field = err.Field()
		element.Tag = err.Tag()
		element.Error = err.Translate(trans)
		errors = append(errors, &element)
	}
	return c.JSON(Response{Code: CodeErrorForm, Data: errors})
}

func ErrorSystem(c *fiber.Ctx, err string) error {
	return c.JSON(Response{Code: CodeErrorSystem, Msg: err})
}

func ErrorUnauthorized(c *fiber.Ctx, msg string, data interface{}) error {
	return c.JSON(Response{Code: CodeErrorUnauthorized, Msg: msg, Data: data})
}

func ErrorForbidden(c *fiber.Ctx, msg string, data interface{}) error {
	return c.JSON(Response{Code: CodeErrorForbidden, Msg: msg, Data: data})
}
