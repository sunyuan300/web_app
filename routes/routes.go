package routes

import (
	"web_app/controller"
	"web_app/logger"

	"github.com/gin-gonic/gin"
)

func SetUp() *gin.Engine {

	r := gin.New()
	r.Use(logger.GinLogger(), logger.GinRecovery(true))

	// 注册业务路由
	r.POST("/sign", controller.SignUpHandler)

	return r
}
