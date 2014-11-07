package controllers

type Auth struct {
	Base
}

func (this *Auth) Prepare() {
	this.Layout = "_frontLayout.html"
	this.LayoutSections = make(map[string]string)
	this.LayoutSections["Head"] = "_head.html"
	this.LayoutSections["Header"] = "_header.html"
	this.LayoutSections["Login"] = "_login.html"
	this.LayoutSections["Footer"] = "_footer.html"
	this.LayoutSections["Scripts"] = "_scripts.html"
}

func (this *Auth) Finish() {

}