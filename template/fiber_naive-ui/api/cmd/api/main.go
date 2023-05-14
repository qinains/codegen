package main

import (
	"fmt"
	"log"

	"github.com/fsnotify/fsnotify"
	"github.com/goccy/go-json"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/spf13/viper"
	"{[{$.moduleName}]}/core"
	"{[{$.moduleName}]}/core/middleware"
	"{[{$.moduleName}]}/router"
)

// @title {[{$.projectName | Pascal}]} App
// @version 1.0
// @description {[{$.projectDescription}]}
// @termsOfService https://lninl.com/terms/
// @contact.name qinains
// @contact.email qinains@qq.com
// @contact.url https://lninl.com/me/
// @license.name MIT
// @license.url ./LICENSE
// @BasePath /
//
//go:generate go install github.com/swaggo/swag/cmd/swag
//go:generate swag init --generalInfo cmd/api/main.go --dir ../../. --outputTypes json --output ../../docs --propertyStrategy camelcase
func main() {
	viper.SetConfigName("config")
	viper.SetConfigType("json")
	viper.AddConfigPath("./")
	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			panic(fmt.Errorf("config file not found(配置文件未找到): %w", err))
		} else {
			panic(fmt.Errorf("config file error(配置文件错误): %w", err))
		}
	}

	core.OnInitDB()
	viper.OnConfigChange(func(e fsnotify.Event) {
		fmt.Println("config file changed(配置文件更新):", e.Name)
		core.OnInitDB()
	})
	viper.WatchConfig()

	app := fiber.New(fiber.Config{
		JSONEncoder: json.Marshal,
		JSONDecoder: json.Unmarshal,
		ErrorHandler: func(c *fiber.Ctx, err error) error {
			code := fiber.StatusInternalServerError
			if e, ok := err.(*fiber.Error); ok {
				code = e.Code
			}
			return c.Status(code).JSON(core.Response{Code: code, Msg: err.Error()})
		},
	})
	app.Use(recover.New(recover.Config{EnableStackTrace: viper.GetBool("isDeveloping")})).Use(cors.New())

	router.OnInit(app)

	core.OnInitI18N()
	core.OnInitCaptchaStore()

	middleware.OnInitAuthz(core.DB)

	core.OnInitApp(app)
	log.Fatal(app.Listen(":" + viper.GetString("http.port")))
}
