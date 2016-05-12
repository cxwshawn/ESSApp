package main

import (
	_ "ESSApp/routers"
	"github.com/astaxie/beego"
    _ "github.com/astaxie/beego/logs"
    "ESSApp/controllers"
    "ESSApp/controllers/essapp"
)

func main() {
    beego.Router("/:urlname(.*)", &controllers.MainController{}, "*:Get") //别名访问
    beego.Router("/login", &controllers.AccountController{})
    //TODO: implement pay and register later.
    beego.Router("/register/index", &essapp.RegisterController{}, "*:Index")
	beego.Run()
}

