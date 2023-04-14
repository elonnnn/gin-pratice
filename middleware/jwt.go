package middleware

import (
	"errors"
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/gin-server/common/response"
	"github.com/gin-server/utils/jwt"
)

func JWTAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.Request.Header.Get("x-token")
		if token == "" {
			response.FailWithDetailed(gin.H{"reload": true}, "未登录或者非法访问", c)
			c.Abort()
			return
		}
		j := jwt.NewJWT()
		// parseToken 解析token包含的信息
		claims, err := j.ParseToken(token)
		if err != nil {
			if errors.Is(err, jwt.TokenExpired) {
				response.FailWithDetailed(gin.H{"reload": true}, "授权已过期", c)
				c.Abort()
				return
			}
			response.FailWithDetailed(gin.H{"reload": true}, err.Error(), c)
			c.Abort()
			return
		}
		fmt.Println(claims)
		c.Next()
	}
}
