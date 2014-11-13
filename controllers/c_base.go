package controllers

import (
	"fmt"
	"net/url"
	"reflect"
	"regexp"
	"strings"
	"time"
	"zouzhe/models"
	"zouzhe/utils"

	"github.com/astaxie/beego"
	"github.com/beego/i18n"
)

type langType struct {
	Lang, Name string
}
type Base struct {
	beego.Controller
	i18n.Locale
	page        models.Page
	currentUser *models.Current
}

//全部单词字符包括中文
const sub = "\\w\u4e00-\u9fa5"

var (
	//log *logs.BeeLogger
	langTypes []*langType
)

func init() {
	//log = logs.NewLogger(10000)

	// 引用beego官网代码
	langs := strings.Split(appconf("lang::types"), "|")
	names := strings.Split(appconf("lang::names"), "|")
	langTypes = make([]*langType, 0, len(langs))
	for i, v := range langs {
		langTypes = append(langTypes, &langType{
			Lang: v,
			Name: names[i],
		})
	}

	for _, lang := range langs {
		beego.Trace("Loading language: " + lang)
		if err := i18n.SetMessage(lang, "conf/"+"locale_"+lang+".ini"); err != nil {
			beego.Error("Fail to set message file: " + err.Error())
			return
		}
	}
}

//读取配置
func appconf(key string) string {
	return beego.AppConfig.String(key)
}

func (this *Base) Prepare() {
	this.initPage()
}

//
func (this *Base) initPage() {

	this.Data["PageStartTime"] = time.Now()

	// Redirect to make URL clean.
	if this.setLangVer() {
		i := strings.Index(this.Ctx.Request.RequestURI, "?")
		this.Redirect(this.Ctx.Request.RequestURI[:i], 302)
		return
	}
	//this.page.Author = this.lang("appauthor")
	//this.page.Company = this.lang("appcompany")
	//this.page.Copyright = this.lang("appcopyright")
	//this.page.Description = this.lang("appdescription")
	//this.page.Domain = this.lang("appdomain")
	//this.page.Keywords = this.lang("appkeywords")
	//this.page.SiteName = this.lang("appsitename")
	//this.page.Title = this.lang("apptitle")
	//this.page.Product = this.lang("appproduct")
	//this.page.Version = this.lang("appversion")

	//this.Data["page"] = &this.page
}

/*
* setLangVer 设置网址语言版本.引用beego官网
 */
func (this *Base) setLangVer() bool {
	isNeedRedir := false
	hasCookie := false

	// 1. Check URL arguments.
	lang := this.Input().Get("lang")

	// 2. Get language information from cookies.
	if len(lang) == 0 {
		lang = this.Ctx.GetCookie("lang")
		hasCookie = true
	} else {
		isNeedRedir = true
	}

	// Check again in case someone modify by purpose.
	if !i18n.IsExist(lang) {
		lang = ""
		isNeedRedir = false
		hasCookie = false
	}

	// 3. Get language information from 'Accept-Language'.
	if len(lang) == 0 {
		al := this.Ctx.Request.Header.Get("Accept-Language")
		if len(al) > 4 {
			al = al[:5] // Only compare first 5 letters.
			if i18n.IsExist(al) {
				lang = al
			}
		}
	}

	// 4. Default language is English.
	if len(lang) == 0 {
		lang = "en-US"
		isNeedRedir = false
	}

	curLang := langType{
		Lang: lang,
	}

	// Save language information in cookies.
	if !hasCookie {
		this.Ctx.SetCookie("lang", curLang.Lang, 1<<31-1, "/")
	}

	restLangs := make([]*langType, 0, len(langTypes)-1)
	for _, v := range langTypes {
		if lang != v.Lang {
			restLangs = append(restLangs, v)
		} else {
			curLang.Name = v.Name
		}
	}

	// Set language properties.
	this.Lang = lang
	this.Data["Lang"] = curLang.Lang
	this.Data["CurLang"] = curLang.Name
	this.Data["RestLangs"] = restLangs

	return isNeedRedir
}

//公共字段
func (this *Base) Extend(dst interface{}) {
	d := reflect.Indirect(reflect.ValueOf(dst))
	if v := d.FieldByName("Updator"); v.IsValid() && v.Int() == 0 && this.currentUser != nil {
		v.SetInt(this.currentUser.Id)
	}
	if v := d.FieldByName("Updated"); v.IsValid() && v.Int() == 0 {
		v.SetInt(utils.Millisecond(time.Now()))
	}
	if v := d.FieldByName("Ip"); v.IsValid() && v.String() == "" {
		v.SetString(this.Ctx.Input.IP())
	}
}

//公共字段
func (this *Base) ExtendEx(dst interface{}) {
	d := reflect.Indirect(reflect.ValueOf(dst))
	if v := d.FieldByName("Updator"); v.IsValid() && v.Int() == 0 && this.currentUser != nil {
		v.SetInt(this.currentUser.Id)
	}
	if v := d.FieldByName("Updated"); v.IsValid() && v.Int() == 0 {
		v.SetInt(utils.Millisecond(time.Now()))
	}
	if v := d.FieldByName("Creator"); v.IsValid() && v.Int() == 0 && this.currentUser != nil {
		v.SetInt(this.currentUser.Id)
	}
	if v := d.FieldByName("Created"); v.IsValid() && v.Int() == 0 {
		v.SetInt(utils.Millisecond(time.Now()))
	}
	if v := d.FieldByName("Ip"); v.IsValid() && v.String() == "" {
		v.SetString(this.Ctx.Input.IP())
	}
}

/*
* 输出 Json 格式数据
 */
func (this *Base) outputJson(data interface{}, err error) {
	if err == nil {
		this.renderJson(utils.JsonData(true, "", data))
	} else {
		this.renderJson(utils.JsonMessage(false, "", err.Error()))
	}
}

//
func (this *Base) setJsonData(data interface{}) {
	//操作成功，清除token
	if resp := reflect.Indirect(reflect.ValueOf(data)); resp.FieldByName("Ok").Bool() {
		this.XsrfToken()
	}
	this.Data["json"] = data
}

//返回json响应格式
func (this *Base) renderJson(data interface{}) {
	this.setJsonData(data)
	this.ServeJson()
}

//返回jsonp响应
func (this *Base) renderJsonp(data interface{}) {
	this.setJsonData(data)
	this.ServeJsonp()
}

//返回html字符串格式响应
func (this *Base) ServeString(arg string) {
	this.Ctx.Output.Body([]byte(arg))
}

//响应签名丢失错误
func (this *Base) renderLoseToken() {
	data := utils.JsonMessage(false, "invalidFormToken", "invalidFormToken")
	this.renderJson(data)
}

//是否外链
func (this *Base) isOutLink() bool {
	host, err := url.Parse(this.Ctx.Request.Referer())
	if err != nil {
		return true
	}
	return this.Ctx.Request.Host != host.Host
}

////文件服务
//func (this *Base) serverFile(file, filename string) {
//	file = filepath.Join(".", file)

//	//友好文件名
//	if len(filename) == 0 {
//		filename = utils.UrlEncode(filepath.Base(file))
//	}

//	this.Ctx.ResponseWriter.Header().Set("Content-Description", "File Transfer")
//	this.Ctx.ResponseWriter.Header().Set("Content-Type", "application/octet-stream;charset=UTF-8")
//	this.Ctx.ResponseWriter.Header().Set("Content-Disposition", "attachment; filename="+filename)
//	this.Ctx.ResponseWriter.Header().Set("Content-Transfer-Encoding", "binary")
//	this.Ctx.ResponseWriter.Header().Set("Expires", "0")
//	this.Ctx.ResponseWriter.Header().Set("Cache-Control", "must-revalidate")
//	this.Ctx.ResponseWriter.Header().Set("Pragma", "public")

//	http.ServeFile(this.Ctx.ResponseWriter, this.Ctx.Request, file)
//	this.end()
//}

//获取URL参数
func (this *Base) getParamsInt64(key string) (int64, error) {
	i64, err := utils.Str2int64(this.getParamsString(key))
	return i64, err
}

func (this *Base) getParamsInt(key string) (int, error) {
	i64, err := this.getParamsInt64(key)
	return int(i64), err
}

func (this *Base) getParamsString(key string) string {
	return this.Ctx.Input.Param(key)
}

//验证合法用户
func (this *Base) validUser() (*models.Current, bool) {
	coo := this.Ctx.GetCookie(beego.AppConfig.String("CookieName"))

	if coo == "" {
		return nil, false
	}

	this.currentUser = this.GetCurrentUser(coo)

	if this.currentUser.Id == 0 {
		return nil, false
	}

	return this.currentUser, true
}

//允许新的请求，数据通用字段初始信息，附带验证用户是否合法(err)，
func (this *Base) allowRequest() (ok bool) {
	_, ok = this.validUser()

	return
}

//读取登录用户的Cookie信息
func (this *Base) GetCurrentUser(cookie string) (currentuser *models.Current) {
	currentuser = new(models.Current)

	cookie = utils.CookieDecode(cookie)

	//拆分cookie
	curr := strings.Split(cookie, "|")
	if len(curr) > 0 {
		currentuser.Id, _ = utils.Str2int64(curr[0]) //strconv.ParseInt(curr[0], 10, 0)
	}
	if len(curr) > 1 {
		currentuser.Name = curr[1]
	}
	if len(curr) > 2 {
		currentuser.Avatar = curr[2]
	}
	if len(curr) > 3 {
		currentuser.Role = curr[3]
	}
	return
}

/*
* 解析首尾#中的字符串
 */
func (this *Base) ParseSharp(str string) []string {
	//匹配字符串
	p := fmt.Sprintf("#([%s]+)#[^%s]*", sub, sub)
	return this.ParseString(str, p)
}

/*
* 解析首字符@尾字符是空格中的子串
 */
func (this *Base) ParseAite(str string) []string {
	//匹配字符串
	p := fmt.Sprintf("@([%s]+)", sub)
	return this.ParseString(str, p)
}

/*
* 解析指定首尾字符中的字符串
 */
func (this *Base) ParseString(str, p string) []string {
	//正则
	re := regexp.MustCompile(p)

	result := make([]string, 0)
	//查找子串
	for _, tags := range re.FindAllStringSubmatch(str, -1) {
		for _, tag := range tags[1:] {
			if tag != "" {
				result = append(result, tag)
			}
		}
	}
	if len(result) == 0 {
		return nil
	}
	return result
}

/*
* xsrf过滤
 */
func (this *Base) CheckXsrf() (bool, string) {
	if this.CheckXsrfCookie() {
		return true, this.XsrfToken()
	}
	return false, ""
}

////获取当前语言
//func (this *Base) lang(k string) string {
//	return utils.Lang(k, this.Ctx.Request.Header.Get("Accept-Language"))
//}

//终止服务
func (this *Base) end() {
	this.Layout = ""
	this.TplNames = ""

	this.StopRun()
}
func (this *Base) SetTplNames(name ...string) {
	c, a := this.Controller.GetControllerAndAction()

	if len(name) > 0 && name[0] != "" {
		a = name[0]
	}
	this.TplNames = strings.ToLower(fmt.Sprintf("%s/%s.html", c, a))
}

/*
* 跟踪
 */
func (this *Base) Trace(v ...interface{}) {
	c, a := this.Controller.GetControllerAndAction()
	beego.Trace(fmt.Sprintf("Controller:%s Action:%s ", c, a) + fmt.Sprintf("Info:%v", v...))
}
