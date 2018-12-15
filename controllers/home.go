package controllers

type HomeController struct {
	baseController
}

func (this *HomeController) Index() {
	session := this.GetSession("user")
	this.Data["user"] = session
	this.SetDefaultTplName()
}
