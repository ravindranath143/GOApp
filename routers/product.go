package routers

import (
	"GOApp/controllers"
	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/products", &controllers.ProductController{},"get:Get")
	beego.Router("/add-products", &controllers.ProductController{},"get:Addproductsview;post:Uploadproducts")

}
