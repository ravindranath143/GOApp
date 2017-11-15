package models

import (
	"github.com/astaxie/beego/orm"
)

type Products struct {
	Id           int       `orm:"column(id);auto"`
	Title        string    `orm:"column(title);size(255)"`
	Description  string    `orm:"column(description);size(255)"`
	Image        string    `orm:"column(image);size(255)"`
	Price        string    `orm:"column(price);size(10)"`
	Stock        int       `orm:"column(stock);size(10)"`
}

func init() {
	orm.RegisterModel(new(Products))
}