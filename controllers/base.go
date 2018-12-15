package controllers

import (
	"fmt"
	"strings"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)

type baseController struct {
	beego.Controller
	o              orm.Ormer
	controllerName string
	actionName     string
}

func (this *baseController) Prepare() {
	// 设置控制器和行动名字
	controllerName, actionName := this.GetControllerAndAction()
	this.controllerName = strings.ToLower(controllerName[0 : len(controllerName)-10])
	this.actionName = strings.ToLower(actionName)
	// 设置orm
	this.o = orm.NewOrm()
	//登陆判断
	isAdminController := strings.ToLower(this.controllerName) == "admin"
	isLoginAction := strings.ToLower(this.actionName) != "login"
	// 检查
	if isAdminController && isLoginAction {
		if this.GetSession("user") == nil {
			this.JumpGoto("未登陆", "/admin/login")
		}
	}
}

func (this *baseController) JumpGoto(msg string, url string) {
	if url == "" {
		this.Ctx.WriteString("<script>alert('" + msg + "');window.history.go(-1);</script>")
		this.StopRun()
	} else {
		this.Redirect(url, 302)
	}
}

func (this *baseController) Jump(url string) {
	this.Redirect(url, 302)
}

func (this *baseController) SetDefaultTplName() {
	this.TplName = this.controllerName + "/" + this.actionName + ".html"
	fmt.Println(this.TplName)
}
