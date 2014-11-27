package controllers

import (
	"fmt"
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
		this.renderJson(utils.JsonData(false, "", getValidErrors(&valid)))
		return
	}
	// 写入数据库
	if _, err := q.Insert(); err == nil {
		this.renderJson(utils.JsonData(true, "", q))
	} else {
		this.renderJson(utils.JsonMessage(false, "", err.Error()))
	}
}

func (this *Question) List() {
	// 读取分页规则
	p := new(models.Pagination)

	if size, err := this.GetInt("size"); err != nil || size == 0 {
		p.Size = 20
	}
	p.Index, _ = this.GetInt("index")
	// 读取查询条件
	when := this.GetString("when")
	where := this.GetString("where")

	cond := "1=1"
	if when != "" {
		cond += fmt.Sprintf(" and when='%s'", when)
	}
	if where != "" {
		cond += fmt.Sprintf(" and where='%s'", where)
	}

	// 拉取
	q := new(models.Question)

	if qs, err := q.List(cond, p); err != nil {
		this.renderJson(utils.JsonMessage(false, "", err.Error()))
	} else {
		this.renderJson(utils.JsonData(true, "", qs))
	}
}
