package routers

import (
	"GOApp/controllers"
	"github.com/astaxie/beego"
)

func init() {
    beego.Router("/orders", &controllers.OrdersController{})
    beego.Router("/add-to-cart", &controllers.OrdersController{},"post:Cart")
    beego.Router("/delete-from-cart", &controllers.OrdersController{},"post:RemovefromCart")
    beego.Router("/cart", &controllers.OrdersController{},"get:Getcart")

}
