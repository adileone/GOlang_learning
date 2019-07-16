package controllers

import (
	"github.com/astaxie/beego"
)

type MainController struct {
	beego.Controller
}

func (main *MainController) HelloBank() {
	main.TplName = "default/hello-bank.tpl"
}
