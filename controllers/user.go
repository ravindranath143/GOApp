package controllers

import (
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"GOApp/models"
)
type UserController struct {
	beego.Controller
}

func (c *UserController) Get() {
	o := orm.NewOrm()
	v := c.GetSession(sessionName)
	if v != nil {
		userget := models.User{Id: v.(models.User).Id}
		err := o.Read(&userget)
		if err == orm.ErrNoRows {
		    fmt.Println("No result found.")
		} else if err == orm.ErrMissPK {
		    fmt.Println("No primary key found.")
		} else {
		    c.Data["userdata"] = userget
		}
        c.Data["profile"] = true
        c.Data["heading"] = "User Profile"
		c.TplName = "users/main.html"
    }else{
    	c.Redirect("/login", 302)
    }	
	
}

func (c *UserController) Update() {
	o := orm.NewOrm()
	v := c.GetSession(sessionName)
	if v != nil {
		user := models.User{Id: v.(models.User).Id}
		if o.Read(&user) == nil {
		    user.Name = c.GetString("name")
		    user.Phone = c.GetString("phone")
		    user.Address = c.GetString("address")
		    if num, err := o.Update(&user); err == nil {
		    	c.SetSession(sessionName, user)
		    	fmt.Println(num)
		        c.Redirect("/", 302)
		    }
		}
    }else{
    	c.Redirect("/login", 302)
    }	
	
	
}