package models

type Question struct {
	Id      int64  `json:"id"`
	When    string `json:"when"`
	Where   string `json:"where"`
	Intro   string `json:"intro"`
	Tags    string `json:"tags"`
	Status  int    `json:"status" valid:"Range(0,1)"`
	Deleted int    `json:"deleted" valid:"Range(0,1)"`
	Creator int64  `json:"creator"`
	Created int64  `json:"created"`
	Updator int64  `json:"updator"`
	Updated int64  `json:"updated"`
	Ip      string `json:"ip" valid:"MaxSize(23)"`
}
