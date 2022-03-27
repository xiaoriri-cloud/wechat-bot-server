package v1

import (
	"github.com/gin-gonic/gin"
	"log"
	"wechat-bot-server/models"
)

func GetUser(c *gin.Context) {
	wid := c.Query("wid")

	isExist, err := models.ExistUserInfoByWid(wid)
	if err != nil {
		c.JSON(200, gin.H{
			"message": "服务器错误",
			"code":    400,
			"data":    nil,
		})
	}

	if isExist {
		userinfo, err := models.GetUserInfo(wid)
		if err != nil {
			log.Printf("查询错误：%v", err)
			c.JSON(200, gin.H{
				"message": "服务器错误",
				"code":    401,
				"data":    nil,
			})
		}
		c.JSON(200, gin.H{
			"message": "ok",
			"code":    200,
			"data":    userinfo,
		})
	} else {
		c.JSON(200, gin.H{
			"message": "数据不存在",
			"code":    200,
			"data":    nil,
		})
	}
}

func AddUser(c *gin.Context) {
	var userInfoDto models.UserInfo
	c.BindJSON(&userInfoDto)
	if userInfoDto.Wid != "" {

		isExist, err := models.ExistUserInfoByWid(userInfoDto.Wid)
		if !isExist {
			err = models.AddUserInfo(userInfoDto)
			if err != nil {
				log.Printf("插入数据错误：%v", err)
				c.JSON(200, gin.H{
					"message": "服务器错误",
					"code":    401,
					"data":    nil,
				})
			}
		}

		c.JSON(200, gin.H{
			"message": "ok",
			"code":    200,
			"data":    "创建成功",
		})

	} else {
		c.JSON(200, gin.H{
			"message": "wid不能为空",
			"code":    400,
			"data":    nil,
		})
	}

}
