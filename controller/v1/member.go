package v1

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gin-server/constants/code"
	"github.com/gin-server/model"
	"github.com/gin-server/validators"
)

func AddMember(c *gin.Context) {
	res := model.Result{}
	mem := model.Member{}
	if err := c.ShouldBind(&mem); err != nil {
		res.SetCode(code.ERROR)
		res.SetMessage(validators.GetErrorMsg(mem, err))
		c.JSON(http.StatusForbidden, res)
		c.Abort()
		return
	}
	// 处理业务(下次再分享)

	data := map[string]interface{}{
		"name": mem.Name,
		"age":  mem.Age,
	}
	res.SetCode(code.SUCCESS)
	res.SetData(data)
	c.JSON(http.StatusOK, res)

}
