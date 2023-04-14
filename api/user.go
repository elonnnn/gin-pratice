package api

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-server/common/response"
	"github.com/gin-server/global"
	"github.com/gin-server/model"
	"github.com/gin-server/utils"
	"github.com/gin-server/utils/jwt"
	uuid "github.com/satori/go.uuid"
)

type UserApi struct {
}

type LoginResponse struct {
	Token     string `json:"token"`
	ExpiresAt int64  `json:"expiresAt"`
}

func (u *UserApi) Login(c *gin.Context) {
	var l model.Login
	err := c.ShouldBindJSON(&l)

	// key := c.ClientIP()
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = utils.Verify(l, utils.LoginVerify)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	uuid := uuid.NewV4()

	u.TokenNext(c, jwt.BaseClaims{
		Username:    l.Username,
		NickName:    l.Username,
		AuthorityId: 1,
		UUID:        uuid,
	})
}

func (u *UserApi) TokenNext(c *gin.Context, user jwt.BaseClaims) {
	j := &jwt.JWT{SigningKey: []byte(global.GVA_VP.GetString("jwt.signing-key"))} // 唯一签名
	claims := j.CreateClaims(jwt.BaseClaims{
		NickName:    user.NickName,
		Username:    user.Username,
		AuthorityId: user.AuthorityId,
		UUID:        user.UUID,
	})
	token, err := j.CreateToken(claims)
	if err != nil {
		response.FailWithMessage("获取token失败", c)
		return
	}
	response.OkWithDetailed(LoginResponse{
		Token:     token,
		ExpiresAt: claims.StandardClaims.ExpiresAt * 1000,
	}, "登录成功", c)
}
