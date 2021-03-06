package routers

import (
	"sitebeego/controllers"
	"github.com/astaxie/beego"
)

func init() {
    beego.Router("/", &controllers.MainController{})
    beego.Router("/hello", &controllers.HelloController{}, "get:Hello")
}
