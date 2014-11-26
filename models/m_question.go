package models

import (
	"github.com/astaxie/beego/validation"
)

type Question struct {
	Id      int64  `json:"id"`
	When    string `json:"when" valid:"Required;MaxSize(100)"`
	Where   string `json:"where" valid:"Required;MaxSize(100)"`
	Intro   string `json:"intro" valid:"MaxSize(250)"`
	Tags    string `json:"tags" valid:"MaxSize(250)"`
	Status  int    `json:"status" valid:"Range(0,1)"`
	Deleted int    `json:"deleted" valid:"Range(0,1)"`
	Creator int64  `json:"creator"`
	Created int64  `json:"created"`
	Updator int64  `json:"updator"`
	Updated int64  `json:"updated"`
	Ip      string `json:"ip" valid:"MaxSize(23)"`
}

//
func (this *Question) Valid(v *validation.Validation) {

}

func (this *Question) Insert() (int64, error) {
	return db.Insert(this)
}
