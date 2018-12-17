package controllers/front

import (
	"fmt"
)

type frontBaseController struct {
	beego.baseController
}

func (this *baseController) Prepare() {
	fmt.Println("asdasd")
}
