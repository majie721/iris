package main

import (
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/mvc"
	"iris_stu/common"
	"iris_stu/repositories"
	"iris_stu/service"
	"log"
	"context"
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


	db,err := common.NewMysqlConn()
	if err !=nil{
		log.Print("db connect error")
	}
	ctx,cancel := context.WithCancel(context.Background())
	defer cancel()


	//注册控制器
	productRepositories := repositories.NewProductManager("iris_product",db)
	productService := service.NewProductService(productRepositories)
	productParty := app.Party("/product")
	product :=mvc.New(productParty)
	product.Register(ctx,productService)

	//启动服务
	app.Run(
			iris.Addr("localhost:6666"),
			iris.WithoutServerError(iris.ErrServerClosed),
			iris.WithOptimizations,
			)


}
