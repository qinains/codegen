package core

import (
	"reflect"
	"strings"

	"github.com/go-playground/validator/v10"
)

func NewValidate() *validator.Validate {
	validate := validator.New()

	//解析struct字段中的label标签，放入Error中
	validate.RegisterTagNameFunc(func(field reflect.StructField) string {
		name := strings.SplitN(field.Tag.Get("label"), ",", 2)[0]
		if name == "-" {
			return ""
		}

		return name
	})

	validate.RegisterValidation("captcha", func(fl validator.FieldLevel) bool {
		kind := fl.Top().Kind()
		var captcha reflect.Value
		if kind == reflect.Pointer {
			captcha = fl.Top().Elem()
		} else if kind == reflect.Struct {
			captcha = fl.Top()
		}
		return CaptchaStore.Verify(captcha.FieldByName("CaptchaID").String(), captcha.FieldByName(fl.StructFieldName()).String(), true)
	})

	validate.RegisterValidation("dbUnique", func(fl validator.FieldLevel) bool {
		sql := DB

		kind := fl.Top().Kind()
		var entity reflect.Value
		if kind == reflect.Pointer {
			entity = fl.Top().Elem()
		} else if kind == reflect.Struct {
			entity = fl.Top()
		}

		str := fl.Param() //比如：user:tenant_id->TenantID&login_name->LoginName
		split1 := strings.Split(str, ":")
		table := split1[0]
		sql = sql.Table(table)

		split2 := strings.Split(split1[1], "&")
		for _, v := range split2 {
			split3 := strings.Split(v, "->")
			filed := split3[0]
			structField := fl.StructFieldName()
			if len(split3) > 1 {
				structField = split3[1]
			}
			sql = sql.Where(filed+"=?", entity.FieldByName(structField).Interface())
		}

		var count int64
		sql.Count(&count)
		return count == 0
	})

	return validate
}
