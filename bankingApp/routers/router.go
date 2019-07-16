package routers

import (
	"bankingApp/controllers"

	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/", &controllers.MainController{})
	beego.Router("/home", &controllers.ManageController{}, "get,post:Home")
	beego.Router("/manage/add", &controllers.ManageController{}, "get,post:Add")
	beego.Router("/manage/view", &controllers.ManageController{}, "get,post:View")
}
