package controllers

import (
	"github.com/astaxie/beego"
    "github.com/astaxie/beego/orm"
    "GOApp/models"
    "strconv"
    // "fmt"
)

type ApiResponse struct {
    Status       int 
    Message      string
    Cartprice    string
    Qty          string
}

type OrdersController struct {
	beego.Controller
}

func (c *OrdersController) Get() {
	v := c.GetSession(sessionName)
	if v != nil {
        c.Data["orders"] = true
        c.Data["userdata"] = v
        c.Data["heading"] = "Your Orders"
		c.TplName = "orders/main.html"
    }else{
    	c.Redirect("/login", 302)
    }	
}

func (c *OrdersController) Cart() {
    v := c.GetSession(sessionName)
    if v != nil {
        o := orm.NewOrm()
        o.Using("default")
        orm.Debug = true
        qty := c.GetString("qty")
        product_id := c.GetString("product_id")
        price := c.GetString("price")
        cart := new(models.Cart)
        cart.User_id = v.(models.User).Id
        cart.Price = price
        i, err := strconv.Atoi(qty)
        if(err == nil){
            cart.Qty = i
        }
        j, err := strconv.Atoi(product_id)
        if(err == nil){
            cart.Product_id = j
        }
        
        id, err := o.Insert(cart)
        if err == nil && id > 0{
            response := new(ApiResponse)
            response.Status = 200
            response.Message = "success"
            response.Cartprice = price
            response.Qty = qty
            c.Data["json"] = response
            c.ServeJSON()
        } 
    }
    
} 

func (c *OrdersController) RemovefromCart() {
    v := c.GetSession(sessionName)
    o := orm.NewOrm()
    if v != nil {
        id, err := strconv.Atoi(c.GetString("id"))
        if err == nil{}
        if num, err := o.Delete(&models.Cart{Id: id}); err == nil {
            if(num > 0){
                response := new(ApiResponse)
                response.Status = 200
                response.Message = "success"
                response.Cartprice = ""
                response.Qty = ""
                c.Data["json"] = response
                c.ServeJSON()
            }
        }
        
    } 
}
func (c *OrdersController) Getcart() {
    v := c.GetSession(sessionName)
    o := orm.NewOrm()
    if v != nil {
        var maps []orm.Params
        num, err := o.Raw("select c.id as cart_id,c.qty,c.price as cost,p.*,DATE_FORMAT(c.created_at, '%d-%M-%Y') as date_format from cart as c inner join products as p on (c.product_id = p.id) where user_id  = ? order by created_at desc", v.(models.User).Id).Values(&maps)

        if err == nil && num > 0 {
            c.Data["cart"] = maps
        }else{
            c.Data["cart"] = nil
        }
        c.Data["orders"] = true
        c.Data["userdata"] = v
        c.Data["heading"] = "Your Cart"
        c.TplName = "orders/cart.html"
    }else{
        c.Redirect("/login", 302)
    }   
}
