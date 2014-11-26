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
func (this *Question) Index() {
	this.SetTplNames("index")
}

//
func (this *Question) Ask() {
	this.SetTplNames()
}

//
func (this *Question) Insert() {
	// 读取表单数据，填充数据对象
	_m_question := new(models.Question)
	_m_question.When = this.GetString("when")
	_m_question.Where = this.GetString("where")
	_m_question.Intro = this.GetString("intro")
	_m_question.Tags = this.GetString("tags")

	this.ExtendEx(_m_question)
	fmt.Println(_m_question)

	// 验证数据的有效性
	valid := validation.Validation{}
	if _ok, _err := valid.Valid(_m_question); _err != nil {
		this.renderJson(utils.JsonMessage(false, "", _err.Error()))
		return
	} else if !_ok {
		// 读取并返回表单数据验证错误
		_errs := make([]models.Error, 0)
		for _, _err := range valid.Errors {
			_errs = append(_errs, models.Error{Key: _err.Key, Message: _err.Message})
		}

		this.renderJson(utils.JsonData(false, "", _errs))
		return
	}

	if _, _err := _m_question.Insert(); _err == nil {
		this.renderJson(utils.JsonData(true, "", _m_question))
	} else {
		this.renderJson(utils.JsonMessage(false, "", _err.Error()))
	}
}
