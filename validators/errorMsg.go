package validators

import (
	"github.com/go-playground/validator/v10"
)

// 验证器接口
type Validator interface {
	// 获取验证器自定义错误信息
	GetMessage() ValidatorMessages
}

// 验证器自定义错误信息字典
type ValidatorMessages map[string]string

// 获取自定义错误信息
func GetErrorMsg(request Validator, err error) string {
	for _, v := range err.(validator.ValidationErrors) {
		if message, exist := request.GetMessage()[v.Field()+"."+v.Tag()]; exist {
			return message
		}
		return v.Error()
	}
	return "Parameter error"
}
