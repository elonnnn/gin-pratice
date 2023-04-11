package model

// 定义 Result 结构体
type Result struct {
	Code    int         `json:"code"`
	Message string      `json:"msg"`
	Data    interface{} `json:"data"`
}

// 设置错误码
func (res *Result) SetCode(code int) *Result {
	res.Code = code
	return res
}

// 设置错误信息
func (res *Result) SetMessage(msg string) *Result {
	res.Message = msg
	return res
}

// 设置返回数据
func (res *Result) SetData(data interface{}) *Result {
	res.Data = data
	return res
}
