package routers

import (
	"GOApp/controllers"
	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/login", &controllers.LoginController{},"get:LoginView;post:LoginUser")
	beego.Router("/logout", &controllers.LoginController{},"get:LogoutUser")
	beego.Router("/signup", &controllers.SignupController{},"get:SignupView;post:SignupUser")
    beego.InsertFilter("/*", beego.BeforeRouter, controllers.FilterUser)
}
