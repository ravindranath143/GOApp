package controllers

import (
	// "fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"GOApp/models"
	"GOApp/utils"
	// "sms-package/sms"
	// "encoding/json"
	// "net/http"
	// "io/ioutil"
 //    "os"

)
var sessionName = beego.AppConfig.String("SessionName")
var myslice []models.User
type LoginController struct {
	beego.Controller
}

type post struct {
	userId               int   
	id              int 
	title         string 
	body string 
}
type jsonobject struct {
    Object post
}
func (c *LoginController) LoginView() {
	// response := sms.Sendsms("https://jsonplaceholder.typicode.com/posts/1")
	// fmt.Println(response)
	// var data map[string]interface{}
	// _ = json.Unmarshal([]byte(response), &data)
	// c.Data["post"] = data
	v := c.GetSession(sessionName)
    if v != nil {
        c.Redirect("/", 302)
    }else{
    	c.TplName = "users/login.html"
    }
}
func (c *LoginController) LoginUser() {
	email := c.GetString("email")
	password := c.GetString("password")
	salt := utils.GetRandomString(15)
	encodedPwd := utils.EncodePassword(password, salt)
	o := orm.NewOrm()
	var user models.User
	err := o.Raw("select * from user where email=? AND password=?",email,encodedPwd).QueryRow(&user)
	if err == nil {
		v := c.GetSession(sessionName)
	    if v == nil {
	        c.SetSession(sessionName, user)
	        c.Redirect("/", 302)
	    }else{
	    	c.Redirect("/login", 302)
	    }	
	}

}
func (c *LoginController) LogoutUser() {
	c.DelSession(sessionName)
	c.Redirect("/login", 302)
}
