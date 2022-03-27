package main

import "github.com/gin-gonic/gin"

func main() {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "怎么部署成功了草",
		})
	})
	r.Run() // 监听并在 0.0.0.0:8080 上启动服务
}
