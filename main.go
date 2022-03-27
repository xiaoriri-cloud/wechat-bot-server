package main

import (
	"wechat-bot-server/models"
	"wechat-bot-server/pkg/setting"
	"wechat-bot-server/routers"
)

func init() {
	setting.Setup()
	models.Setup()
}

func main() {
	r := routers.InitRouter()
	r.Run() // 监听并在 0.0.0.0:8080 上启动服务
}
