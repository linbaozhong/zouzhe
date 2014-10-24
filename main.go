package main

import (
	_ "zouzhe/routers"

	"github.com/astaxie/beego"
)

func main() {
	beego.Run()
}

func init() {
	//模板后缀
	beego.AddTemplateExt(".html")
	//静态目录
	beego.SetStaticPath("/htm", "html")
}
