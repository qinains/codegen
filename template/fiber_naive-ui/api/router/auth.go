package router

import (
	"github.com/gofiber/fiber/v2"
	"{{$.moduleName}}/core/middleware"
	"{{$.moduleName}}/web"
)

func initAuth(app *fiber.App) {
	app.Post("/captcha", web.Captcha)                         // 获取验证码
	app.Post("/check-captcha", web.CheckCaptcha)              // 验证验证码
	app.Post("/register", web.Register)                       // 注册
	app.Post("/tenant", web.FindTenant)                       // 获取租户列表
	app.Post("/login", web.Login)                             // 登录
	app.Post("/refresh-access-token", web.RefreshAccessToken) // 刷新AccessToken
	app.Get("/auth/info", middleware.JWT(), web.GetAuth)      // 获取用户角色关联详情
}
