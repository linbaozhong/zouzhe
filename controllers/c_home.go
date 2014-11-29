package controllers

type Home struct {
	Front
}

func (this *Home) Get() {
	this.SetTplNames("index")
}

//
func (this *Home) Login() {
	this.Layout = "_frontLayout.html"
	this.SetTplNames()
}
