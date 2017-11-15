package models

import (
	"github.com/astaxie/beego/orm"
)

type Cart struct {
	Id           int       `orm:"column(id);auto"`
	User_id      int       `orm:"column(user_id);size(11)"`
	Product_id   int       `orm:"column(product_id);size(11)"`
	Qty          int       `orm:"column(qty);size(11)"`
	Price        string    `orm:"column(price);size(10)"`
	created_at   string    `orm:"column(price);size(10)"`
	updated_at   string    `orm:"column(price);size(10)"`
}

func init() {
	orm.RegisterModel(new(Cart))
}