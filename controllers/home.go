package controllers

type HomeController struct {
	baseController
}

func (this *HomeController) Index() {
	this.Data["user"] = session
	this.SetDefaultTplName()
}
