package controllers

type Front struct {
	Base
}

func (this *Front) Prepare() {
	this.Layout = "_frontLayout.html"
	this.LayoutSections = make(map[string]string)
	this.LayoutSections["Head"] = "_head.html"
	this.LayoutSections["Header"] = "_header.html"
	this.LayoutSections["Footer"] = "_footer.html"
	this.LayoutSections["Scripts"] = "_scripts.html"
}

func (this *Front) Finish() {

}
