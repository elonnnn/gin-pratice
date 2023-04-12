package api

type ApiGroup struct {
	UserApi UserApi
}

var ApiGroupApp = new(ApiGroup)
