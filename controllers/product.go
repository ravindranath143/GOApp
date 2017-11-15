package controllers

import (
	// "fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"GOApp/models"
	"strconv"
)

type ProductController struct {
	beego.Controller
}

func (c *ProductController) Get() {
	v := c.GetSession(sessionName)
	o := orm.NewOrm()
	if v != nil {
		var maps []orm.Params
	  	num, err := o.Raw("select * from products").Values(&maps)

	  	if err == nil && num > 0 {
		    c.Data["products_data"] = maps
		}else{
			c.Data["products_data"] = nil
		}

        c.Data["products"] = true
        c.Data["userdata"] = v
        c.Data["heading"] = "Available Products"
		c.TplName = "products/main.html"
    }else{
    	c.Redirect("/login", 302)
    }
	
}
func (c *ProductController) Addproductsview() {
	v := c.GetSession(sessionName)
	if v != nil {
		c.TplName = "products/add-products.html"
    }else{
    	c.Redirect("/login", 302)
    }
	
}
func (c *ProductController) Uploadproducts() {
	v := c.GetSession(sessionName)
	if v != nil {
		o := orm.NewOrm()
		o.Using("default")
		orm.Debug = true
		title := c.GetString("title")
		description := c.GetString("description")
		image := c.GetString("image")
		price := c.GetString("price")
		stock := c.GetString("stock")
		product := new(models.Products)
	    product.Title = title
	    product.Description = description
	    product.Image = "/static/img/"+image
	    product.Price = price
	    i, err := strconv.Atoi(stock)
	    if(err == nil){
	    	product.Stock = i
	    }else{
	    	product.Stock = 0
	    }
	    
	    id, err := o.Insert(product)
	    if err == nil && id > 0{
	    	c.Redirect("/products", 302)
	    }

	    
    }else{
    	c.Redirect("/login", 302)
    }
	
}
