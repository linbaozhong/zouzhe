package controllers

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
func (this *Question) Post() {

}
