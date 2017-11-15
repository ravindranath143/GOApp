package routers

import (
	"GOApp/controllers"
	"github.com/astaxie/beego"
)

func init() {
    beego.Router("/profile", &controllers.UserController{}, "get:Get;post:Update")
}
