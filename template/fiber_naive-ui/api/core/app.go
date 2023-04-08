package core

import "github.com/gofiber/fiber/v2"

var App *fiber.App

func OnInitApp(app *fiber.App) {
	App = app
}
