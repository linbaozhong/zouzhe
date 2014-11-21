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
	beego.AutoRouter(&controllers.Connect{})

	qst := &controllers.Question{}
	beego.Router("/question", qst)
	beego.AutoRouter(qst)

	beego.Router("/profile", &controllers.Profile{})

	act := &controllers.Account{}
	beego.Router("/signin", act, "post:SignIn")
	beego.Router("/signout", act, "post:SignOut")
	beego.Router("/signup", act, "get:SignUp")
	beego.Router("/passwordreset", act, "get:PasswordReset")
}
