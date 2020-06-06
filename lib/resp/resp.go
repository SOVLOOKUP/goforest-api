package resp

type Resp struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

func Resp404(data ...interface{}) *Resp {
	return &Resp{
			Code: 404,
			Msg:  "ResourceNotFound",
			Data: data,
		}
}
func Resp500(data ...interface{}) *Resp {
	return &Resp{
		Code: 500,
		Msg:  "InternetServerError",
		Data: data,
	}
}

func Resp403(data ...interface{}) *Resp {
	return &Resp{
		Code: 403,
		Msg:  "Forbidden",
		Data: data,
	}
}