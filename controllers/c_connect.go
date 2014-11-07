package controllers

import (
	"fmt"
)

type Connect struct {
	Base
}

var (
	appid    = appconf("qq::appid")
	callback = appconf("qq::callback")
	appkey   = appconf("qq::appkey")
	//---只读取用户信息
	scope = ""
)

/*
* QQ登录
 */
func (this *Connect) QQ_Login() {
	//---生成唯一随机串防止csrf攻击
	state := this.XsrfToken()
	//---将随机串存入Session
	this.SetSession("state", state)
	//---登录的url
	url := appconf("qq::auth") + "?response_type=code&client_id=" + appid + "&redirect_uri=" + callback + "&state=" + state + "&scope=" + scope

	fmt.Println(url)
	//---
	this.Redirect(url, 302)

	fmt.Println("after redirect ...")
}

/*
* QQ登录回调
 */
func (this *Connect) QQ_Callback() {

	//---检查是否有接口调用错误
	if this.GetString("msg") != "" {
		//---跳转至错误页
		this.Redirect(this.UrlFor("Connect.QQ_Error", ":msg", this.GetString("msg")), 302)
	}

	//---验证state防止csrf攻击
	if this.GetString("state") != this.GetSession("state").(string) {
		//---跳转至错误页
	}

	//---读取token的url
	url := appconf("qq::token") + "?grant_type=authorization_code&client_id=" + appid + "&redirect_uri=" + callback + "&client_secret=" + appkey + "&code=" + this.GetString("code")

	fmt.Println(url)
	//---
	this.Redirect(url, 302)

	fmt.Println("after redirect ...")

	//---在这里解析返回包

}

/*
* QQ登录回调地址
 */
func (this *Connect) QQ_Openid() {
	//---检查是否有接口调用错误
	if this.GetString("msg") != "" {
		//---跳转至错误页
		this.Redirect(this.UrlFor("Connect.QQ_Error", ":msg", this.GetString("msg")), 302)
	}
}

/*
* QQ登录错误地址
 */
func (this *Connect) QQ_Error() {
	fmt.Println(this.GetString("msg"))
	this.Data["message"] = this.GetString("msg")
}
