package routers

import (
	"github.com/gin-gonic/gin"
	v1 "wechat-bot-server/routers/api/v1"
)

func InitRouter() *gin.Engine {
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())
	apiv1 := r.Group("/api/v1")

	apiv1.GET("/userinfo", v1.GetUser)
	apiv1.POST("/userinfo", v1.AddUser)
	return r
}
