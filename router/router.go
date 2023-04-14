package router

import (
	"net/url"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/gin-server/api"
	"github.com/gin-server/common"
	v1 "github.com/gin-server/controller/v1"
	v2 "github.com/gin-server/controller/v2"
	"github.com/gin-server/middleware"
	"github.com/gin-server/validators"
	"github.com/go-playground/validator/v10"
)

func InitRouter(r *gin.Engine) {
	r.GET("/sn", SignDemo)
	// Login Api
	r.POST("user/login", api.ApiGroupApp.UserApi.Login)
	// v1
	GroupV1 := r.Group("/v1")
	{
		GroupV1.Any("/product/add", v1.AddProduct)
		GroupV1.Any("/member/add", v1.AddMember)
	}
	GroupV2 := r.Group("/v2")
	GroupV2.Use(middleware.JWTAuth())
	{
		GroupV2.Any("/product/add", v2.AddProduct)
		GroupV2.Any("/member/add", v2.AddMember)
	}

	// 绑定验证器
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		v.RegisterValidation("NameValid", validators.NameValid)
	}
}

func SignDemo(c *gin.Context) {
	ts := strconv.FormatInt(common.GetTimeUnix(), 10)
	res := map[string]interface{}{}
	params := url.Values{
		"name":  []string{"a"},
		"price": []string{"10"},
		"ts":    []string{ts},
	}
	res["sn"] = common.CreateSign(params)
	res["ts"] = ts
	common.RetJson("200", "success", res, c)
}
