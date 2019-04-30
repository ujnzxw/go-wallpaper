package routers

import (
	"go-wallpaper/controllers"

	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/", &controllers.MainController{})
	beego.Router("/switchChange", &controllers.MainController{}, "post:DoSwitchChange")
}
