package routers

import (
	"zouzhe/controllers"

	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/", &controllers.Home{})
}
