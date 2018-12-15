package controllers

import (
	"blog/models"
	"blog/utils"
	"fmt"
	"strings"

	"github.com/astaxie/beego"
)

type AdminController struct {
	baseController
}

func (this *AdminController) Login() {
	if this.Ctx.Request.Method == "POST" {
		username := this.GetString("username")
		password := this.GetString("password")
		user := models.User{Username: username}
		// 查询用户
		err := this.o.Read(&user, "username")
		if err != nil {
			fmt.Printf("err :%v", err)
			beego.Critical("asdasd")
		}
		if user.Id == 0 {
			this.JumpGoto("账号不存在", "")
		}
		// 输入密码
		inputPasswordMd5 := strings.ToLower(utils.Md5(password))
		// 保存密码
		savePasswordMd5 := strings.ToLower(strings.Trim(user.Password, " "))
		// 判断密码
		if inputPasswordMd5 != savePasswordMd5 {
			this.JumpGoto("密码错误", "")
		}
		this.SetSession("user", user)
		this.Jump("/home/index")
	}
	this.SetDefaultTplName()
}

func (this *AdminController) Main() {
	this.SetDefaultTplName()
}
