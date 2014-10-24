package controllers

import (
	"github.com/astaxie/beego"
)

type Base struct {
	beego.Controller
}

func init() {}

//输出字符串
func (this *Base) toString(arg string) {
	this.Ctx.Output.Body([]byte(arg))
}
