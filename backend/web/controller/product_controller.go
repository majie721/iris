package controller

import (
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/mvc"
	"iris_stu/service"
)

type ProductController struct {
	Ctx iris.Context
	ProductService service.ProductService
}

func (p *ProductController) GetAll() mvc.View {
	productArray,_ := p.ProductService.GetAllProduct()
	return mvc.View{
		Name: "product/view.html",
		Data: iris.Map{
			"productArray":productArray,
		},
	}
}