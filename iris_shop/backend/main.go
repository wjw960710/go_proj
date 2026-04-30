package main

import "github.com/kataras/iris/v12"

func main() {
	app := iris.New()
	app.Logger().SetLevel("debug")
	tmpl := iris.HTML("./backend/web/views", ".html").Layout("shared/layout.html").Reload(true)
	app.RegisterView(tmpl)
	app.HandleDir("/assets", iris.Dir("./backend/web/assets"))
	//出現異常跳轉的頁面
	app.OnAnyErrorCode(func(ctx iris.Context) {
		ctx.ViewData("message", ctx.Values().GetStringDefault("message", "訪問的頁面出錯!"))
		ctx.ViewLayout("")
		ctx.View("shared/error.html")
	})
	app.Run(iris.Addr(":8080"), iris.WithoutServerError(iris.ErrServerClosed), iris.WithOptimizations)
}
