package core

import (
	"reflect"
	"strings"

	"github.com/go-playground/validator/v10"
)

var Validate *validator.Validate

func OnInitValidate() error {
	Validate = validator.New()
	Validate.RegisterTagNameFunc(func(field reflect.StructField) string {
		name := strings.SplitN(field.Tag.Get("label"), ",", 2)[0]
		if name == "-" {
			return ""
		}

		return name
	})
	err := Validate.RegisterValidation("sample_validation", CheckSampleValidation)
	return err
}

func CheckSampleValidation(fl validator.FieldLevel) bool {
	// todo 样例，需要修改
	return false
}
