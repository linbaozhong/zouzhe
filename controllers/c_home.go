package controllers

type Home struct {
	Front
}

func (this *Home) Get() {
	this.Data["Title"] = appTitle
	this.TplNames = "home/index.html"
}
