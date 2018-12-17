package controllers

import (
	"fmt"
	"strings"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)

type baseController struct {
	beego.Controller
	o       orm.Ormer
	ctlName string
	actName string
}

func (this *baseController) Prepare() {
	// 设置控制器和行动名字
	ctlName, actName := this.GetControllerAndAction()
	this.ctlName = strings.ToLower(ctlName[0 : len(ctlName)-10])
	this.actName = strings.ToLower(actName)
	// 设置orm
	this.o = orm.NewOrm()
}

func (this *baseController) Jump(url string) {
	this.Redirect(url, 302)
}

func (this *baseController) SetDefaultTplName() {
	this.TplName = this.controllerName + "/" + this.actionName + ".html"
	fmt.Println(this.TplName)
}
