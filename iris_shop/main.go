package main

import (
	"iris_shop/web/controllers"

	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/mvc"
)

func main() {
	app := iris.New()
	app.Logger().SetLevel("debug")
	app.RegisterView(iris.HTML("./web/views", ".html").Reload(true))
	mvc.New(app.Party("/")).Handle(new(controllers.MovieController))
	app.Listen(":8080")
}
