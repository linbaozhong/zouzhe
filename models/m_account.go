package models

import (
	// "errors"
	// "fmt"
	"github.com/astaxie/beego/validation"
	//"notes/utils"
	//"notes/logs"
)

type Accounts struct {
	Id          int64  `json:"accoundId"`
	LoginName   string `json:"loginName" valid:"MaxSize(100)"`
	Password    string `json:"password"`
	RealName    string `json:"realName" valid:"MaxSize(50)"`
	OpenId      string `json:"openId" valid:"MaxSize(32)"`
	OpenFrom    string `json:"openFrom" valid:"MaxSize(10)"`
	NickName    string `json:"nickName" valid:"MaxSize(50)"`
	Avatar      string `json:"avatar" valid:"MaxSize(250)"`
	AccessToken string `json:"accessToken" valid:"MaxSize(32)"`
	Status      int    `json:"status" valid:"Range(0,1)"`
	Deleted     int    `json:"deleted" valid:"Range(0,1)"`
	Updator     int64  `json:"updator"`
	Updated     int64  `json:"updated"`
	Ip          string `json:"ip" valid:"MaxSize(23)"`
}

//
func (this *Accounts) Valid(v *validation.Validation) {
	//登录名必须是email
	if this.LoginName != "" {
		v.Email(this.LoginName, "loginName")
	}
}

//
func (this *Accounts) New() (int64, error, []Error) {

	if this.RealName == "" {
		this.RealName = this.NickName
	}
	//数据有效性检验
	if d, err := dataCheck(this); err != nil {
		return 0, err, d
	}
	//
	id, err := db.Insert(this)
	return id, err, nil
}

//是否存在
func (this *Accounts) Exists() (bool, error) {
	return db.Get(this)
}
