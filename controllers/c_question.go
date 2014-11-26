package controllers

import (
	"github.com/astaxie/beego/validation"
	"zouzhe/models"
	"zouzhe/utils"
)

type Question struct {
	Auth
}

//
func (this *Question) Get() {
	this.SetTplNames("index")
}

//
func (this *Question) Ask() {
	this.SetTplNames()
}

//
func (this *Question) Insert() {
	q := new(models.Question)
	q.When = this.GetString("when")
	q.Where = this.GetString("where")
	q.Intro = this.GetString("intro")
	q.Tags = this.GetString("tags")
	this.ExtendEx(q)

	// 检验数据的有效性
	valid := validation.Validation{}
	if ok, err := valid.Valid(q); err != nil {
		this.renderJson(utils.JsonMessage(false, "", err.Error()))
		return
	} else if !ok {
		errs := make([]models.Error, 0)
		for _, err := range valid.Errors {
			errs = append(errs, models.Error{Key: err.Key, Message: err.Message})
		}
		this.renderJson(utils.JsonData(false, "", errs))
		return
	}
	// 写入数据库
	if _, err := q.Insert(); err == nil {
		this.renderJson(utils.JsonData(true, "", q))
	} else {
		this.renderJson(utils.JsonMessage(false, "", err.Error()))
	}
}
