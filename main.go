package main

import (
	"blog/controllers"
	"blog/models"
	_ "blog/routers"

	"github.com/astaxie/beego"
)

func init() {
	models.Init()
	beego.BConfig.WebConfig.Session.SessionOn = true
}

func main() {
	beego.ErrorController(&controllers.ErrorController{})
	beego.Run()
}
