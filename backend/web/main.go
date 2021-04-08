package main

import (
	"github.com/kataras/iris/v12"
)

func main()  {

	//1.创建实例
	app := iris.New()

	//2设置错误级别
	app.Logger().SetLevel("debug")

	//注册模板
	tmplate := iris.HTML("./backend/web/views",".html").Layout("shared/layout.html").Reload(true)
	app.RegisterView(tmplate)

	//静态文件的路径
	app.HandleDir("/assets","./backend/web/assets")

	//跳转错误页面 TODO 错误处理
	app.OnAnyErrorCode(func(ctx iris.Context) {
		ctx.ViewData("message",ctx.Values().GetStringDefault("message","访问错误"))
		ctx.ViewLayout("")
		ctx.View("shared/error.html")
	})


	//注册控制器

	//启动服务
	app.Run(
			iris.Addr("localhost:6666"),
			iris.WithoutServerError(iris.ErrServerClosed),
			iris.WithOptimizations,
			)


}
