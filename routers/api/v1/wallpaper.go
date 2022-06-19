package v1

import (
	"github.com/gin-gonic/gin"
	"wechat-bot-server/models"
)

func PageWallpaper(c *gin.Context) {

	data, _ := models.PageWallpaper(1, 2)
	var result []models.Wallpaper
	for _, item := range data {
		item.ImgUrl = "https://cdn.xiaoriri.com/" + item.ImgUrl
		result = append(result, item)
	}

	c.JSON(200, gin.H{
		"message": "",
		"total":   7533,
		"code":    200,
		"data":    result,
	})
}
