package telegram

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"wechat-bot-server/models"
	"wechat-bot-server/pkg/response"
)

func GetToken(c *gin.Context) {
	configName := c.Query("name")
	var result = response.Result{
		Code: 200,
		Msg:  "ok",
		Data: "",
	}
	if configName != "" {
		config, _ := models.GetConfig(configName)
		if config.Name == "" {
			result.Data = nil
		} else {
			result.Data = config
		}
	} else {
		result.Code = 10001
		result.Msg = "参数校验错误，name不能为空"
		c.JSON(http.StatusOK, result)
		return
	}
	c.JSON(200, result)
}

type setTokenRequest struct {
	Name   string `form:"name" binding:"required"`
	Config string `form:"config" binding:"required"`
}

func SetToken(c *gin.Context) {
	var params setTokenRequest
	err := c.Bind(&params)
	var result = response.Result{
		Code: 200,
		Msg:  "ok",
		Data: params,
	}
	if err != nil {
		result.Code = 10001
		result.Msg = "参数校验错误"
		result.Data = err.Error()
		c.JSON(http.StatusOK, result)
		return
	}
	var config models.Config
	config.Name = params.Name
	config.Config = params.Config
	err = models.SetConfig(config)

	if err != nil {
		result.Code = 10001
		result.Msg = "保存失败"
		result.Data = err.Error()
	}
	c.JSON(http.StatusOK, result)
}
