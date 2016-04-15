package controllers

import (
	"github.com/astaxie/beego"
)

type HelloController struct {
	beego.Controller
}

func (c *HelloController) Hello() {
	c.Data["Website"] = "beego.me1"
	c.Data["Email"] = "astaxie@gmail.com"
	c.TplName = "index.tpl"
}
