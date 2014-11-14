/*
* 实现第三方账号登录的相关逻辑
 */
package controllers

import (
	"errors"
	"fmt"
	"strings"
	"zouzhe/utils"

	"github.com/astaxie/beego/httplib"
)

type Connect struct {
	Base
}

type OpenSign struct {
	Ret      string //返回码 qq=0表示成功
	Msg      string //错误时的消息提示
	Id       string //本站账户id
	From     string //第三方标识
	OpenId   string //账户id
	Token    string //access_token
	Refresh  string //refresh_token
	Nickname string //昵称
	Gender   string //性别
	Avatar_1 string //40*40的qq头像
	Avatar_2 string //100*100的qq头像
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

	//fmt.Println(url)
	//---
	this.Redirect(url, 302)

	//fmt.Println("after redirect ...")
}

/*
* QQ登录回调
 */
func (this *Connect) QQ_Callback() {

	//---检查是否有接口调用错误
	if this.GetString("msg") != "" {
		//---跳转至错误页
		this.Redirect(this.UrlFor("Connect.Connect_Error", ":msg", this.GetString("msg")), 302)
		return
	}

	//---验证state防止csrf攻击
	if this.GetString("state") != this.GetSession("state").(string) {
		//---跳转至错误页
		this.Redirect(this.UrlFor("Connect.Connect_Error", ":msg", ""), 302)
		return
	}

	//---opensign第三方账户信息
	_account := new(OpenSign)
	_account.From = "qq"

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

		jmap := make(map[string]string)
		//---解析返回的内容
		param := strings.Split(rep, "&")

		for _, item := range param {
			pos := strings.Index(item, "=")
			jmap[item[:pos]] = item[pos+1:]
		}
		//--- access_token
		_account.Token = jmap["access_token"]
		//--- refresh_token
		_account.Refresh = jmap["refresh_token"]

	} else {
		//---跳转至错误页
		this.Redirect(this.UrlFor("Connect.Connect_Error", ":msg", err.Error()), 302)
		return
	}

	//---创建读取openid的请求
	if this.qq_openid(_account) != nil {
		//---跳转至错误页
		this.Redirect(this.UrlFor("Connect.Connect_Error", ":msg", err.Error()), 302)
		return
	}

	//---创建读取userinfo的请求
	if this.qq_userinfo(_account) != nil {
		//---跳转至错误页
		this.Redirect(this.UrlFor("Connect.Connect_Error", ":msg", err.Error()), 302)
		return
	}

	// 写入cookie
	//this.cookie("openid", _account.OpenId)
	this.cookie("token", _account.Token)
	this.cookie("nickname", _account.Nickname)
	this.cookie("avatar", _account.Avatar_1)
	this.cookie("from", _account.From)

	this.Trace(_account)
	//
	this.Data["sign"] = _account

	this.SetTplNames("callback")

}

/*
* QQ登录获取openid
 */
func (this *Connect) qq_openid(act *OpenSign) (err error) {
	//---创建读取openid的请求
	req := httplib.Get(appconf("qq::openid"))
	req.Param("access_token", act.Token)

	//---读取返回的内容
	rep, err := req.String()

	if err == nil {

		jmap := make(map[string]interface{})
		//---解析返回的内容,检查如果包含callback,读取openid
		if strings.Contains(rep, "callback") {
			jmap = utils.JsonString2map(rep[strings.Index(rep, "(")+1 : strings.LastIndex(rep, ")")])
		}

		if len(jmap) > 0 {
			//--- openid
			act.OpenId = utils.Interface2str(jmap["openid"])
		} else {
			err = errors.New("return value is empty")
		}
	}
	return
}

/*
* QQ登录用户信息
 */
func (this *Connect) qq_userinfo(act *OpenSign) (err error) {
	//---创建读取userinfo的请求
	req := httplib.Get(appconf("qq::userinfo"))
	req.Param("access_token", act.Token)
	req.Param("openid", act.OpenId)
	req.Param("oauth_consumer_key", appid)

	//---读取返回的内容
	rep, err := req.String()

	if err == nil {

		//---解析返回的内容,检查如果包含callback,读取openid
		jmap := utils.JsonString2map(rep)

		if len(jmap) > 0 {

			act.Ret = utils.Interface2str(jmap["ret"])
			//--- 检查返回码是否正确
			if act.Ret == "0" {
				act.Nickname = utils.Interface2str(jmap["nickname"])
				act.Gender = utils.Interface2str(jmap["gender"])
				act.Avatar_1 = utils.Interface2str(jmap["figureurl_qq_1"])
				act.Avatar_2 = utils.Interface2str(jmap["figureurl_qq_2"])
			} else {
				act.Msg = utils.Interface2str(jmap["msg"])
			}
		} else {
			err = errors.New("return value is empty")
		}

	}
	return
}

/*
* QQ登录access_token续期
 */
func (this *Connect) qq_refresh(act *OpenSign) (err error) {
	//---创建读取token的请求
	req := httplib.Get(appconf("qq::token"))
	req.Param("grant_type", "refresh_token")
	req.Param("client_secret", appkey)
	req.Param("client_id", appid)
	req.Param("refresh_token", act.Refresh)

	//---读取返回的内容
	rep, err := req.String()

	if err == nil && strings.Contains(rep, "access_token") && strings.Contains(rep, "refresh_token") {

		jmap := make(map[string]string)
		//---解析返回的内容
		param := strings.Split(rep, "&")

		for _, item := range param {
			pos := strings.Index(item, "=")
			jmap[item[:pos]] = item[pos+1:]
		}
		//--- access_token
		act.Token = jmap["access_token"]
		//--- refresh_token
		act.Refresh = jmap["refresh_token"]
		// 写入cookie
		this.cookie("token", act.Token)
	}
	return
}

/*
* 登录错误地址
 */
func (this *Connect) Connect_Error() {
	this.Trace(this.GetString("msg"))
	this.Data["message"] = this.GetString("msg")
}

/*
* 记录登录历史
* 1.检查该账户是否存在，如果不存在，创建之，反之，记录登录日期和ip地址
* 2.检查是否需要进行授权续期，如需要，续期
 */
func (this *Connect) SignTrace() {

	_account := new(OpenSign)
	_account.Id = this.Ctx.GetCookie("_snow_id")
	_account.From = this.GetString("from")
	_account.Gender = this.GetString("gender")
	_account.Nickname = this.GetString("nickName")
	_account.OpenId = this.GetString("openId")
	_account.Token = this.GetString("token")
	_account.Refresh = this.GetString("refresh")
	_account.Avatar_1 = this.GetString("avatar_1")
	_account.Avatar_2 = this.GetString("avatar_2")

	//无效的请求
	if _account.Id == "" && _account.OpenId == "" {
		this.renderJson(utils.JsonMessage(false, "", ""))
		return
	}
	// 1.检查该账户是否存在，
	// 如果不存在，创建之，反之，记录登录日期和ip地址
	// 保存登录状态
	this.signin(this._sonw_key(_account.Id, _account.From))

	// 并返回是否需要续期
	_needRefresh := false

	// 2.续期
	if _needRefresh {
		this.qq_refresh(_account)
	}

	this.renderJson(utils.JsonMessage(true, "", ""))
}
