package main

import (
	"context"
	"fmt"
	"github.com/kataras/iris"
	"github.com/kataras/iris/mvc"
	"miao/backend/common"
	"miao/backend/web/controllers"
	"miao/repositories"
	"miao/services"
)

func main()  {
	// 1. 创建Application 实例
	app :=iris.New()
	// 2. 开启debug模式
	app.Logger().SetLevel("debug")
	// 3.注册模板 创建html视图引擎并设置基类模板
	template :=iris.HTML("./backend/web/views",".html").Layout("shared/" +
		"layout.html").Reload(true)
	app.RegisterView(template)
	// 4.设置模板静态资源路径
	app.StaticWeb("/assets","./backend/web/assets")
	//出现异常跳转到指定页面
	app.OnAnyErrorCode(func(ctx iris.Context) {
		ctx.ViewData("message",ctx.Values().GetStringDefault("message","访问的页面出错！"))
		ctx.ViewLayout("")
		ctx.View("shared/error.html")
	})
	// 6. 注册控制器
	// 6.1 连接数据库
	db ,err :=common.NewMysqlConn()
	if err != nil {
		fmt.Println(err)
		return
	}
	ctx,cancel :=context.WithCancel(context.Background())
	defer cancel()
	// 6.2 注册控制器
	productRepository := repositories.NewProductManager("product",db)
	productService :=services.NewProductService(productRepository)
	productParty := app.Party("/product")
	product := mvc.New(productParty)
	product.Register(ctx,productService)
	product.Handle(new(controllers.ProductController))

	// 7.启动服务
	if err =app.Run(
		iris.Addr("localhost:8080"),
		//iris.WithoutVersionChecker,
		iris.WithoutServerError(iris.ErrServerClosed),
		iris.WithOptimizations,
		); err != nil {
			fmt.Printf("app run failed,err:%v",err)
		}
}
