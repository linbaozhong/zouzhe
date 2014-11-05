package controllers

import (
	"fmt"
	"zouzhe/models"
	"zouzhe/utils"

	"github.com/astaxie/beego"
)

type Account struct {
	Base
}

// var (
// 	m_account     = &models.Account{}
// 	m_openAccount = &models.OpenAccount{}
// )

/*
* 新账户
 */
func (this *Account) New() {
	// //过滤跨域
	// if ok, _ := this.CheckXsrf(); !ok {
	// 	this.renderLoseToken()
	// 	return
	// }

	oa := &models.Accounts{OpenId: this.GetString("openId"), OpenFrom: this.GetString("openFrom")}

	//是否已经存在
	if has, err := oa.Exists(); has {
		//fmt.Println(oa)
	} else {

		oa.AccessToken = this.GetString("accessToken")
		oa.Avatar = this.GetString("avatar")
		oa.NickName = this.GetString("nickName")

		this.Extend(oa)

		if _, err, _ = oa.New(); err == nil {
			this.Trace(oa)
		} else {
			this.Trace(err)
		}

	}
	//登录状态存入cookie，缺省时间是1年：365*24*60*60
	var cookieDuration interface{}
	always := false
	if always {
		cookieDuration, _ = beego.AppConfig.Int("CookieDuration")

	} else {
		cookieDuration = ""
	}

	this.Ctx.SetCookie(beego.AppConfig.String("CookieName"),
		utils.CookieEncode(fmt.Sprintf("%d|%s|%s", oa.Id, oa.NickName, oa.Avatar)),
		cookieDuration, "/")

	this.renderJson(utils.JsonData(true, "", oa))
}
