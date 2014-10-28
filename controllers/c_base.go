package controllers

import (
	"github.com/astaxie/beego"
)

type Base struct {
	beego.Controller
}

var (
	appTitle       = beego.AppConfig.String("apptitle")
	appDescription = beego.AppConfig.String("appdescription")
	appAuthor      = beego.AppConfig.String("appauthor")
	appKeywords    = beego.AppConfig.String("appkeywords")
)

func init() {}

//输出字符串
func (this *Base) toString(arg string) {
	this.Ctx.Output.Body([]byte(arg))
}
