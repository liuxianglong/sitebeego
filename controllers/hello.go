package controllers

import (
	"github.com/astaxie/beego"
)

type MainController struct {
	beego.Controller
}

func (c *MainController) Hello() {
	c.Data["Website"] = "beego.me1"
	c.Data["Email"] = "astaxie@gmail.com"
	c.TplName = "index.tpl"
}
