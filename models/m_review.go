package models

import (
	"github.com/astaxie/beego/validation"
)

type Review struct {
	Id      int64  `json:"id"`
	Type    int    `json:"type" valid:"Required"`
	TypeId  int64  `json:"typeId" valid:"Required"`
	Content string `json:"intro"`
	Status  int    `json:"status" valid:"Range(0,1)"`
	Deleted int    `json:"deleted" valid:"Range(0,1)"`
	Creator int64  `json:"creator"`
	Created int64  `json:"created"`
	Updator int64  `json:"updator"`
	Updated int64  `json:"updated"`
	Ip      string `json:"ip" valid:"MaxSize(23)"`
}

//
func (this *Review) Valid(v *validation.Validation) {

}

// 新问题
func (this *Review) Insert() (int64, error) {
	return db.Insert(this)
}

// 读取一个问题
func (this *Review) Get() (bool, error) {
	return db.Get(this)
}
