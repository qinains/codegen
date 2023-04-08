package core

import "github.com/mojocn/base64Captcha"

var CaptchaStore base64Captcha.Store

func OnInitCaptchaStore() {
	CaptchaStore = base64Captcha.DefaultMemStore
}
