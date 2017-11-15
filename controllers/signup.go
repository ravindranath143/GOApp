package controllers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"GOApp/models"
	"GOApp/utils"
)

type SignupController struct {
	beego.Controller
}

func (c *SignupController) SignupView() {
	c.TplName = "users/signup.html"
}
func (c *SignupController) SignupUser() {
	o := orm.NewOrm()
	o.Using("default")
	orm.Debug = true
	name := c.GetString("name")
	email := c.GetString("email")
	phone := c.GetString("phone")
	address := c.GetString("address")
	password := c.GetString("password")
	salt := utils.GetRandomString(15)
	encodedPwd := utils.EncodePassword(password, salt)
	user := new(models.User)
    user.Name = name
    user.Email = email
    user.Phone = phone
    user.Address = address
    user.Password = encodedPwd
    o.Insert(user)

    c.Redirect("/login", 302)

}
