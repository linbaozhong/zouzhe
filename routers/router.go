package routers

import (
	"zouzhe/controllers"

	"github.com/astaxie/beego"
)

func init() {
	// 屏蔽路由大小写敏感
	beego.RouterCaseSensitive = false
	beego.Router("/", &controllers.Home{})
	beego.Router("/connect/qq_error/:msg", &controllers.Connect{})
	beego.Router("/profile", &controllers.Profile{})
	beego.AutoRouter(&controllers.Connect{})

	act := &controllers.Account{}
	beego.Router("/signin", act, "post:SignIn")
	beego.Router("/signout", act, "post:SignOut")
	beego.Router("/signup", act, "get:SignUp")
	beego.Router("/passwordreset", act, "get:PasswordReset")
}
