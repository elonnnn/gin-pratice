package model

import "github.com/gin-server/validators"

// 定义 Member 结构体
type Member struct {
	Name string `form:"name" json:"name" binding:"required,NameValid"`
	Age  int    `form:"age"  json:"age"  binding:"required,gt=10,lt=120"`
}

func (Member Member) GetMessage() validators.ValidatorMessages {
	return validators.ValidatorMessages{
		"Name.required": "请输入用户名",
		"Age.required":  "请输入年龄",
	}
}
