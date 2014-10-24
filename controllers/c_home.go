package controllers

type Home struct {
	Front
}

func (this *Home) Get() {
	this.Data["Title"] = "走着……"
	this.TplNames = "index.html"
}
