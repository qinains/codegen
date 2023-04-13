package router

import (
	"os"

	swagger "github.com/arsmn/fiber-swagger/v2"
	"github.com/gofiber/fiber/v2"
	"github.com/spf13/viper"
)

func initSwagger(app *fiber.App) {
	if viper.GetBool("isDeveloping") {
		if _, err := os.Stat("./docs/swagger.json"); os.IsNotExist(err) {
			app.Static("/swagger/doc.json", "../../docs/swagger.json")
		} else {
			app.Static("/swagger/doc.json", "/docs/swagger.json")
		}
		app.Get("/swagger/*", swagger.New(swagger.Config{ // custom
			URL:          "/swagger/doc.json",
			DeepLinking:  false,
			DocExpansion: "none",
			OAuth: &swagger.OAuthConfig{
				AppName: "{{$.projectName}} App",
			},
			OAuth2RedirectUrl: "/swagger/oauth2-redirect.html",
		}))
	}
}
