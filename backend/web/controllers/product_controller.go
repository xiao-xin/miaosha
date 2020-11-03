package controllers

import (
	"github.com/kataras/iris"
	"github.com/kataras/iris/mvc"
	"miao/services"
)

// 控制器结构体
type ProductController struct {
	Cxt iris.Context
	ProductService services.IProductService
}

// 获取所有产品
func (c *ProductController) GetAll() mvc.View{
	products := c.ProductService.GetAllProduct()
	return mvc.View{
		Name:"product/index.html",
		Data:iris.Map{
			"products":products,
		},
	}
}


