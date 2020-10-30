package main

import (
	"github.com/kataras/iris"
)

func main()  {
	// 1. 创建Application 实例
	app :=iris.New()
	// 2. 开启debug模式
	app.Logger().SetLevel("debug")
	// 3.注册模板 创建html视图引擎并设置基类模板
	template :=iris.HTML("./backend/web/views",".html").Layout("shared/" +
		".html").Reload(true)
	app.RegisterView(template)
	// 4.设置模板静态资源路径
	app.StaticWeb("/assets","./backend/web/assets")
	// 5.指定异常页面
	app.OnAnyErrorCode(func(ctx iris.Context) {
		// 设置message的值
		ctx.ViewData("message",ctx.Values().GetStringDefault("message","访问页面出错"))
		// 设置布局文件，不填使用前面设置的
		ctx.ViewLayout("")
		// 设置错误页面
		ctx.View("shared/error.html")
	})
	// 6. 注册控制器

	// 7.启动服务
	app.Run(
		iris.Addr("127.0.0.1:8080"),
		iris.WithoutServerError(iris.ErrServerClosed),
		iris.WithOptimizations,
		)
}
