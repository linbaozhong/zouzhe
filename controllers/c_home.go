package controllers

type Home struct {
	Front
}

func (this *Home) Get() {
	this.SetTplNames("index")
}
