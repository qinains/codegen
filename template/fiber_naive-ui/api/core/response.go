package core

import (
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

const (
	CodeOK                = 20000
	CodeErrorBusiness     = 30000
	CodeErrorForm         = 40000
	CodeErrorCaptcha      = 40001
	CodeErrorUnauthorized = 40401
	CodeErrorForbidden    = 40403
	CodeErrorSystem       = 50000
	CodeErrorToken        = 50008
	CodeErrorTokenExpired = 50014
)

type Response struct {
	Code int         // 错误码：20000正常,30000业务异常,40000表单异常,40001需要验证码,40403无权限,50000系统异常,50008Token异常,50014Token过期
	Msg  string      // 错误信息
	Data interface{} // 正常返回内容
}
type ErrorResponse struct {
	Field string //错误字段
	Tag   string //错误标记
	Value string //错误值
	Error string //错误内容
}
type List struct {
	Total int64       // 总页码
	Items interface{} // 项目列表
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

func ErrorBusiness(c *fiber.Ctx, msg string) error {
	return c.JSON(Response{Code: CodeErrorBusiness, Msg: msg})
}

func ErrorForm(c *fiber.Ctx, data interface{}) error {
	return c.JSON(Response{Code: CodeErrorForm, Data: data})
}

func ErrorFormValidationErrors(c *fiber.Ctx, validate *validator.Validate, err error) error {
	trans := GetTranslator(validate, c.Get("accept-language"))
	var errors []*ErrorResponse
	for _, err := range err.(validator.ValidationErrors) {
		var element ErrorResponse
		element.Field = err.StructNamespace()
		element.Tag = err.Tag()
		element.Value = err.Param()
		element.Error = err.Translate(trans)
		errors = append(errors, &element)
	}
	return c.JSON(Response{Code: CodeErrorForm, Data: errors})
}

func ErrorSystem(c *fiber.Ctx, err string) error {
	return c.JSON(Response{Code: CodeErrorSystem, Msg: err})
}

func Error(c *fiber.Ctx, msg string, data interface{}) error {
	return c.JSON(Response{Code: CodeErrorForm, Msg: msg, Data: data})
}
