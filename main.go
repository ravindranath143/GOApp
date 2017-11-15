package main

import (
	_ "GOApp/routers"
	_ "github.com/astaxie/beego/session/mysql"
	_ "github.com/astaxie/beego/session/redis"
	"github.com/astaxie/beego/session" 
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)
var globalSessions *session.Manager
func init() {
	orm.RegisterDriver("mysql", orm.DRMySQL)
    orm.RegisterDataBase("default", "mysql", "root:@/ecomapp?charset=utf8")
}

func main() {
	beego.Run()
	beego.BConfig.WebConfig.Session.SessionProvider = "redis"
    beego.BConfig.WebConfig.Session.SessionProviderConfig = "127.0.0.1:6379"
}

