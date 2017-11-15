package models

import (
	"github.com/astaxie/beego/orm"
)

type User struct {
	Id           int       `orm:"column(id);auto"`
	Name         string    `orm:"column(name);size(255)"`
	Email        string    `orm:"column(email);size(255)"`
	Address      string    `orm:"column(address);size(255)"`
	Password     string    `orm:"column(password);size(255)"`
	Phone        string    `orm:"column(phone);size(255)"`
}

func init() {
	orm.RegisterModel(new(User))
}