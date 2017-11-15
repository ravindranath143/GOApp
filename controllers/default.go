package controllers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context"
	"github.com/astaxie/beego/orm"
	// "GOApp/models"
)

type MainController struct {
	beego.Controller
}

func (c *MainController) Get() {
	o := orm.NewOrm()
	var maps []orm.Params
  	num, err := o.Raw("select * from user").Values(&maps)
  	if err == nil && num > 0 {
	    c.Data["users"] = maps
	}else{
		c.Data["users"] = nil
	}
	v := c.GetSession(sessionName)
	if v != nil {
        c.Data["home"] = true
        c.Data["userdata"] = v
		c.TplName = "home/index.html"
    }else{
    	c.Redirect("/login", 302)
    }
}
// customize filters for fine grain authorization
var FilterUser = func(ctx *context.Context) {
	sess, _ := beego.GlobalSessions.SessionStart(ctx.ResponseWriter, ctx.Request)
    // defer sess.SessionRelease(ctx.ResponseWriter)
    // read the session from the request
    ses := sess.Get(sessionName)
	// _, ok := ctx.Input.Session(sessionName).(int)
	if ses == nil && ctx.Input.URL() != "/login" && ctx.Input.URL() != "/signup" {
		ctx.Redirect(302, "/login")
	}
}