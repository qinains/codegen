package web

import (
	"github.com/gofiber/fiber/v2"
	"github.com/mojocn/base64Captcha"
	"{{$.moduleName}}/core"
	"{{$.moduleName}}/core/middleware"
	"{{$.moduleName}}/dto"
	"{{$.moduleName}}/service"
)

// Register 用户注册
// @Summary 用户注册
// @Tags 认证
// @Accept json
// @Produce json
// @Param json body dto.Register true "注册信息"
// @Success 200 {object} core.Response{data=dto.Token} "注册成功"
// @Failure 400 {object} core.Response{msg=string,data=[]core.ErrorResponse} "如果是业务异常，一般是msg显示异常值，data为null。如果是表单异常，一般msg为空字符串，data包含异常项"
// @Router /register [post]
func Register(c *fiber.Ctx) error {
	register := &dto.Register{}
	if err := c.BodyParser(register); err != nil {
		return core.ErrorForm(c, err.Error())
	}

	validate := core.NewValidate()

	// validate.RegisterStructValidation(func(sl validator.StructLevel) {
	// 	register := sl.Current().Interface().(dto.Register)
	// 	auth, _ := service.AuthService.GetByPhoneOrLoginName(register.TenantID, register.Phone, register.LoginName)

	// 	if auth.LoginName != "" {
	// 		sl.ReportError(auth.LoginName, "LoginName", "LoginName", "dbUnique", auth.LoginName)
	// 	}
	// 	if auth.Phone != "" {
	// 		sl.ReportError(auth.Phone, "Phone", "Phone", "dbUnique", auth.Phone)
	// 	}
	// }, register)

	if err := validate.Struct(register); err != nil {
		return core.ErrorFormValidationErrors(c, validate, err)
	}

	token, err := service.AuthService.Register(register)
	if err != nil {
		return core.ErrorBusiness(c, err.Error())
	}
	return core.OK(c, token)
}

// Login 用户登录
// @Summary 用户登录
// @Tags 认证
// @Accept json
// @Produce json
// @Param json body dto.Login true "登录信息"
// @Success 200 {object} core.Response{data=dto.Token} "token信息"
// @Failure 400 {object} core.Response{msg=string,data=[]core.ErrorResponse} "如果是业务异常，一般是msg显示异常值，data为null。如果是表单异常，一般msg为空字符串，data包含异常项"
// @Router /login [post]
func Login(c *fiber.Ctx) error {
	login := &dto.Login{}
	if err := c.BodyParser(login); err != nil {
		return core.ErrorForm(c, err.Error())
	}

	validate := core.NewValidate()

	if err := validate.Struct(login); err != nil {
		return core.ErrorFormValidationErrors(c, validate, err)
	}

	token, err := service.AuthService.Login(login)
	if err != nil {
		return core.ErrorBusiness(c, err.Error())
	}

	return core.OK(c, token)
}

// GetAuth 获取用户详情
// @Summary 获取用户详情
// @Tags 认证
// @Accept json
// @Produce json
// @Param Authorization header string true "Token信息" default(Bearer {AccessToken})
// @Success 200 {object} core.Response{data=dto.AuthResp} "用户详情"
// @Failure 400 {object} core.Response{msg=string,data=[]core.ErrorResponse} "如果是业务异常，一般是msg显示异常值，data为null。如果是表单异常，一般msg为空字符串，data包含异常项"
// @Router /auth/info [get]
func GetAuth(c *fiber.Ctx) error {
	_, UID, _, err := middleware.GetByContextKey(c)
	if err != nil {
		return core.ErrorBusiness(c, err.Error())
	}

	auth, err := service.AuthService.Get(UID)
	if err != nil {
		return core.ErrorBusiness(c, err.Error())
	}
	return core.OK(c, &auth)
}

// RefreshAccessToken 续期AccessToken
// @Summary 续期AccessToken
// @Tags 认证
// @Accept json
// @Produce json
// @Param json body dto.Token true "Token信息，必须包含RefreshToken"
// @Success 200 {object} core.Response{data=dto.Token} "Token信息"
// @Failure 400 {object} core.Response{msg=string,data=[]core.ErrorResponse} "如果是业务异常，一般是msg显示异常值，data为null。如果是表单异常，一般msg为空字符串，data包含异常项"
// @Router /refresh-access-token [post]
func RefreshAccessToken(c *fiber.Ctx) error {
	var dtoToken dto.Token
	if err := c.BodyParser(&dtoToken); err != nil {
		core.ErrorForm(c, err.Error())
		return nil
	}

	token, err := service.AuthService.RefreshAccessToken(&dtoToken)
	if err != nil {
		return core.ErrorBusiness(c, err.Error())
	}

	return core.OK(c, token)
}

// Captcha 图形验证码
// @Summary 图形验证码
// @Tags 认证
// @Accept json
// @Produce json
// @Success 200 {object} core.Response{data=dto.Captcha} "base64验证码"
// @Failure 400 {object} core.Response{msg=string,data=[]core.ErrorResponse} "如果是业务异常，一般是msg显示异常值，data为null。如果是表单异常，一般msg为空字符串，data包含异常项"
// @Router /captcha [post]
func Captcha(c *fiber.Ctx) error {
	driver := &base64Captcha.DriverString{
		Length:          4,
		Height:          60,
		Width:           240,
		ShowLineOptions: 2,
		NoiseCount:      0,
		Source:          "1234567890abcdefghijkmnprtuvw",
	}
	captcha := base64Captcha.NewCaptcha(driver, core.CaptchaStore)
	id, b64s, err := captcha.Generate()
	if err != nil {
		return core.ErrorBusiness(c, err.Error())
	}
	var dtoCaptcha dto.Captcha
	dtoCaptcha.CaptchaID = id
	dtoCaptcha.Data = b64s
	return core.OK(c, dtoCaptcha)
}

// CheckCaptcha 检验图形验证码
// @Summary 检验图形验证码
// @Tags 认证
// @Accept json
// @Produce json
// @Param json body dto.Captcha true "验证码"
// @Success 200 {object} core.Response{data=bool} "验证结果"
// @Failure 400 {object} core.Response{msg=string,data=[]core.ErrorResponse} "如果是业务异常，一般是msg显示异常值，data为null。如果是表单异常，一般msg为空字符串，data包含异常项"
// @Router /check-captcha [post]
func CheckCaptcha(c *fiber.Ctx) error {
	var dtoCaptcha dto.Captcha
	if err := c.BodyParser(&dtoCaptcha); err != nil {
		return core.ErrorForm(c, err.Error())
	}

	// todo: 验证条件判断（IP、客户端ID），获取N次失效，过期时间判断，全局性验证码

	return core.OK(c, core.CaptchaStore.Verify(dtoCaptcha.CaptchaID, dtoCaptcha.Code, false))
}
