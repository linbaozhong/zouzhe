package utils

type Response struct {
	Ok   bool        `json:"ok"`
	Key  string      `json:"key"`
	Data interface{} `json:"data"`
}

//返回JSON格式消息
func JsonMessage(ok bool, k string, d interface{}) *Response {
	var r = new(Response)
	r.Ok = ok
	r.Key = k
	r.Data = d
	return r
}

//返回JSON格式对象
func JsonData(ok bool, k string, d interface{}) *Response {
	var r = new(Response)
	r.Ok = ok
	r.Key = k
	r.Data = d
	return r
}
