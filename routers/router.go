package routers

import (
	"zouzhe/controllers"

	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/", &controllers.Home{})
	beego.Router("/connect/qq_error/:msg", &controllers.Connect{})
	beego.Router("/profile", &controllers.Profile{})
	beego.AutoRouter(&controllers.Connect{})
}
