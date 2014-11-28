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

/*
方法：读取一条数据
参数：id
*/
func (this *Question) Get() {
	id, err := this.GetInt64("id")

	if err != nil {
		this.renderJson(utils.JsonMessage(false, "", "参数错误"))
		return
	}

	//
	q := new(models.Question)
	q.Id = id

	if has, err := q.Get(); err == nil {
		if has {
			this.renderJson(utils.JsonData(true, "", q))
		} else {
			this.renderJson(utils.JsonMessage(false, "", "数据不存在"))
		}
	} else {
		this.renderJson(utils.JsonMessage(false, "", err.Error()))
	}
}

//
func (this *Question) Ask() {
	this.SetTplNames()
}

/*
发布一条线路求助
*/
func (this *Question) Save() {
	q := new(models.Question)

	if id, err := this.GetInt64("id"); err == nil {
		q.Id = id
	} else {
		q.Id = 0
	}
	q.When = this.GetString("when")
	q.Where = this.GetString("where")
	q.Intro = this.GetString("intro")
	q.Tags = this.GetString("tags")

	// Id>0是Update，Id=0是insert
	if q.Id == 0 {
		this.ExtendEx(q)
	} else {
		this.Extend(q)
	}

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
	if _, err := q.Save(); err == nil {
		this.renderJson(utils.JsonData(true, "", q))
	} else {
		this.renderJson(utils.JsonMessage(false, "", err.Error()))
	}
}

// @Title List
// @Description 分页拉取问题列表
// @Param   size  form  int  false        "每页的记录条数"
// @Success 200 {object} utils.Response
// @Failure 200 {object} utils.Response
// @router /List [post]
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
	// 构造查询字符串
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
