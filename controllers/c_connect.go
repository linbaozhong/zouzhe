package controllers

import (
	"fmt"
	"strings"
	"zouzhe/utils"

	"github.com/astaxie/beego/httplib"
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
		return
	}

	//---验证state防止csrf攻击
	if this.GetString("state") != this.GetSession("state").(string) {
		//---跳转至错误页
		this.Redirect(this.UrlFor("Connect.QQ_Error", ":msg", ""), 302)
		return
	}

	//---读取token的url
	//url := appconf("qq::token") + "?grant_type=authorization_code&client_id=" + appid + "&redirect_uri=" + callback + "&client_secret=" + appkey + "&code=" + this.GetString("code") + "&state=" + this.GetString("state")

	//fmt.Println(url)

	var access_token, refresh_token, openid string

	//---创建读取token的请求
	req := httplib.Get(appconf("qq::token"))
	req.Param("grant_type", "authorization_code")
	req.Param("client_id", appid)
	req.Param("client_secret", appkey)
	req.Param("redirect_uri", callback)
	req.Param("code", this.GetString("code"))
	req.Param("state", this.GetString("state"))

	//---读取返回的内容
	rep, err := req.String()

	if err == nil {

		//---解析返回的内容
		param := strings.Split(rep, "&")
		jmap := make(map[string]string)

		for _, item := range param {
			pos := strings.Index(item, "=")
			jmap[item[:pos]] = item[pos+1:]
		}
		//--- access_token
		access_token = jmap["access_token"]
		//--- refresh_token
		refresh_token = jmap["refresh_token"]

	} else {
		//---跳转至错误页
		this.Redirect(this.UrlFor("Connect.QQ_Error", ":msg", err.Error()), 302)
		return
	}

	//---创建读取openid的请求
	req = httplib.Get(appconf("qq::openid"))
	req.Param("access_token", access_token)

	//---读取返回的内容
	rep, err = req.String()

	if err == nil {

		fmt.Println(rep)

		//---解析返回的内容,检查如果包含callback,读取openid
		jmap := getCallback(rep)

		if len(jmap) > 0 {
			//--- openid
			openid = utils.Interface2str(jmap["openid"])
		}
	} else {
		//---跳转至错误页
		this.Redirect(this.UrlFor("Connect.QQ_Error", ":msg", err.Error()), 302)
		return
	}

	//---创建读取userinfo的请求
	req = httplib.Get(appconf("qq::userinfo"))
	req.Param("access_token", access_token)
	req.Param("openid", openid)
	req.Param("oauth_consumer_key", appid)

	//---读取返回的内容
	rep, err = req.String()

	if err == nil {

		fmt.Println(rep)
		//---解析返回的内容,检查如果包含callback,读取openid
		jmap := utils.JsonString2map(rep)

		if len(jmap) > 0 {
			//--- openid
			this.Data["nickname"] = utils.Interface2str(jmap["nickname"])
		}

	} else {
		//---跳转至错误页
		this.Redirect(this.UrlFor("Connect.QQ_Error", ":msg", err.Error()), 302)
		return
	}

	//
	this.Data["openid"] = openid
	this.Data["access_token"] = access_token
	this.Data["refresh_token"] = refresh_token
	this.SetTplNames()

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

//
func getCallback(str string) (jmap map[string]interface{}) {
	if strings.Contains(str, "callback") {
		jmap = utils.JsonString2map(str[strings.Index(str, "(")+1 : strings.LastIndex(str, ")")])

	}
	return
}
