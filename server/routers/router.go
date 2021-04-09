package routers

import (
	"github.com/kataras/iris"
	"github.com/lwaly/tars_gateway_web/server/controllers"
)

func Route() {
	app := iris.New()
	app.Post("/", controllers.ServerListGet)
	app.Listen(":8080")
}
