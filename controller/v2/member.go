package v2

import (
	"github.com/gin-gonic/gin"
)

func AddMember(c *gin.Context) {
	// 获取 Get 参数
	name := c.Query("name")

	c.JSON(200, gin.H{
		"v2":   "AddMember",
		"name": name,
	})
}
