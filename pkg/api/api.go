package api

import (
	"fmt"
	iris "github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/middleware/logger"
	"github.com/kataras/iris/v12/middleware/recover"
)

func Start(port string) {
	address := fmt.Sprintf(":%s", port)
	app := iris.New()
	app.Logger().SetLevel("debug")

	app.Use(recover.New())
	app.Use(logger.New())

	app.Handle("GET", "/", func(ctx iris.Context) {
		ctx.HTML("<h1>Welcome to Vizix</h1>")
	})

	app.Get("/ping", func(ctx iris.Context) {
		ctx.WriteString("pong")
	})

	app.Get("/health-check", func(ctx iris.Context) {
		ctx.JSON(iris.Map{"success": true})
	})

	app.Run(iris.Addr(address), iris.WithoutServerError(iris.ErrServerClosed))
}