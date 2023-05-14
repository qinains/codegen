package router

import (
	"github.com/gofiber/fiber/v2"
	"{[{$.moduleName}]}/core/middleware"
	"{[{$.moduleName}]}/web"
)

func OnInit(app *fiber.App) {
	initAuth(app)
	initStatic(app)
	initSwagger(app)
	initSystem(app)
	initRouter(app)
}

func initRouter(app *fiber.App) {
{[{range $k0,$table := .tables}]}
	app.Post("/{[{$table.tableName | Dash}]}/delete", middleware.JWT(), middleware.Auth(), web.Delete{[{$table.tableName | Pascal}]}) //删除{[{$table.tableComment | Breaker}]}
	app.Post("/{[{$table.tableName | Dash}]}/update", middleware.JWT(), middleware.Auth(), web.Update{[{$table.tableName | Pascal}]}) //更新{[{$table.tableComment | Breaker}]}
	app.Post("/{[{$table.tableName | Dash}]}", middleware.JWT(), middleware.Auth(), web.Find{[{$table.tableName | Pascal}]})          //获取{[{$table.tableComment | Breaker}]}列表
	app.Get("/{[{$table.tableName | Dash}]}/:ID", middleware.JWT(), middleware.Auth(), web.Get{[{$table.tableName | Pascal}]})        //获取{[{$table.tableComment | Breaker}]}详情
	app.Post("/{[{$table.tableName | Dash}]}/create", middleware.JWT(), middleware.Auth(), web.Create{[{$table.tableName | Pascal}]}) //添加{[{$table.tableComment | Breaker}]}
{[{end}]}
}
