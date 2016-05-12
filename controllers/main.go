package controllers

import (
	"github.com/astaxie/beego"
    "fmt"
    // "os"
    )

type MainController struct {
	beego.Controller
}

func (c *MainController) Get() {
    // l := logs.GetLogger()
    filename := c.Ctx.Input.Param(":urlname")
    fmt.Println(filename)
	c.Data["Website"] = "beego.me"
	c.Data["Email"] = "cxwshawn@gmail.com"
    // if _, err := os.Stat(beego.BConfig.WebConfig.ViewsPath + "/" + filename); err == nil {
    //     c.Layout = beego.BConfig.WebConfig.ViewsPath + "/" + filename
    // }
    if filename == "" {
        c.TplName = "index.html"
    } else {
        c.TplName = filename    
    }
}
