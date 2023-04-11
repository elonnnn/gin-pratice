package main

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-server/core"
	"github.com/gin-server/global"
	"github.com/gin-server/middleware"
	"github.com/gin-server/router"
)

func main() {
	global.GVA_VP = core.Viper() // 初始化Viper
	gin.SetMode(gin.ReleaseMode)
	engine := gin.Default()
	engine.Use(middleware.LoggerToFile())
	router.InitRouter(engine)
	engine.Run(global.GVA_VP.GetString("app.port"))
}
