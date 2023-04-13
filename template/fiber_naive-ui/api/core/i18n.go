package core

import (
	"strings"

	"github.com/go-playground/locales/en"
	"github.com/go-playground/locales/zh"
	"github.com/go-playground/locales/zh_Hans_HK"
	"github.com/go-playground/locales/zh_Hant_TW"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	enTrans "github.com/go-playground/validator/v10/translations/en"
	zhTrans "github.com/go-playground/validator/v10/translations/zh"
	zhTWTrans "github.com/go-playground/validator/v10/translations/zh_tw"
)

var Uni *ut.UniversalTranslator

func OnInitI18N() (err error) {
	zhTranslator := zh.New()
	zhHansHKTranslator := zh_Hans_HK.New()
	zhHantTWTranslator := zh_Hant_TW.New()
	enTranslator := en.New()
	Uni = ut.New(zhTranslator, zhTranslator, zhHansHKTranslator, zhHantTWTranslator, enTranslator)
	return
}

func GetTranslator(validate *validator.Validate, language string) ut.Translator {
	if strings.HasPrefix(language, "zh-TW") || strings.HasPrefix(language, "zh-HK") {
		language = "zh_tw"
	} else if strings.HasPrefix(language, "en") {
		language = "en"
	} else {
		language = "zh"
	}
	trans, _ := Uni.GetTranslator(language)

	switch language {
	case "zh":
		_ = zhTrans.RegisterDefaultTranslations(validate, trans)

		validate.RegisterTranslation("captcha", trans, func(ut ut.Translator) error {
			return ut.Add("captcha", "{0}错误", true)
		}, func(ut ut.Translator, fe validator.FieldError) string {
			t, _ := ut.T("captcha", fe.Field())
			return t
		})
		validate.RegisterTranslation("dbUnique", trans, func(ut ut.Translator) error {
			return ut.Add("dbUnique", "{0}已经存在", true)
		}, func(ut ut.Translator, fe validator.FieldError) string {
			t, _ := ut.T("dbUnique", fe.Field())
			return t
		})
	case "zh_tw":
		_ = zhTWTrans.RegisterDefaultTranslations(validate, trans)

		validate.RegisterTranslation("captcha", trans, func(ut ut.Translator) error {
			return ut.Add("captcha", "驗證碼錯誤", true)
		}, func(ut ut.Translator, fe validator.FieldError) string {
			t, _ := ut.T("captcha", fe.StructField())
			return t
		})
		validate.RegisterTranslation("dbUnique", trans, func(ut ut.Translator) error {
			return ut.Add("dbUnique", "{0}已經存在", true)
		}, func(ut ut.Translator, fe validator.FieldError) string {
			t, _ := ut.T("dbUnique", fe.StructField())
			return t
		})
	case "en":
		_ = enTrans.RegisterDefaultTranslations(validate, trans)

		validate.RegisterTranslation("captcha", trans, func(ut ut.Translator) error {
			return ut.Add("captcha", "captcha is not match", true)
		}, func(ut ut.Translator, fe validator.FieldError) string {
			t, _ := ut.T("captcha", fe.StructField())
			return t
		})
		validate.RegisterTranslation("dbUnique", trans, func(ut ut.Translator) error {
			return ut.Add("dbUnique", "{0} is exist", true)
		}, func(ut ut.Translator, fe validator.FieldError) string {
			t, _ := ut.T("dbUnique", fe.StructField())
			return t
		})
	}
	return trans
}
