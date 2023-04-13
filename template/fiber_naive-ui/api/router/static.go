package router

import (
	"github.com/gofiber/fiber/v2"
	"github.com/spf13/viper"
)

func initStatic(app *fiber.App) {
	app.Static("favicon.ico", "./favicon.ico")
	app.Static("/", "./index.html")
	app.Static("/static", "./static")
	app.Static(viper.GetString("file.webRelativePath"), viper.GetString("file.webUploadRoot")) //文件访问
}
