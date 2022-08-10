package routers

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"
	"wechat-bot-server/models"
	v1 "wechat-bot-server/routers/api/v1"
	"wechat-bot-server/telegram"
)

func InitRouter() *gin.Engine {
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())
	apiv1 := r.Group("/api/v1")

	apiv1.GET("/userinfo", v1.GetUser)
	appPath, _ := os.Getwd()
	r.StaticFS("/telegram", http.Dir(appPath+"/telegram/views"))

	telegramR := r.Group("/tg/")
	telegramR.GET("/config", telegram.GetConfig)
	telegramR.POST("/config", telegram.SetConfig)
	telegramR.GET("/getUpdates", func(c *gin.Context) {
		res := telegram.GetUpdates()
		c.JSON(200, res)
	})

	r.POST("/haier/create", func(c *gin.Context) {
		var haierinfoDto models.HaierInfo
		c.BindJSON(&haierinfoDto)
		err := models.AddHaierInfo(haierinfoDto)
		if err != nil {
			log.Printf("插入数据错误：%v", err)
			c.JSON(200, gin.H{
				"message": "服务器错误",
				"code":    401,
				"data":    nil,
			})
		}
		c.JSON(200, gin.H{
			"message": "ok",
			"code":    200,
			"data":    "创建成功",
		})
	})

		r.GET("/vote", func(context *gin.Context) {

		netTransport := &http.Transport{
			//Proxy:                 http.ProxyURL(proxy),
			MaxIdleConnsPerHost:   10,
			ResponseHeaderTimeout: time.Second * time.Duration(5),
		}

		client := &http.Client{
			Timeout:   time.Second * 10,
			Transport: netTransport,
		}
		req, _ := http.NewRequest("GET", "http://nbd332.wh.changqingmall.cn/Home/index.php?m=Index&a=vote&vid=610717&id=12231&tp=", nil)
		resp, _ := client.Do(req)
		if resp != nil {
			defer func(Body io.ReadCloser) {
				err := Body.Close()
				if err != nil {
					log.Println("defer 报错：" + err.Error())
				}
			}(resp.Body)
		}
		body, _ := ioutil.ReadAll(resp.Body)
		var isTrue = strings.Contains(string(body), "投票成功")

		if isTrue {
			context.JSON(200, gin.H{
				"message": "投票成功",
				"code":    200,
				"data":    "创建成功",
			})
		} else {
			context.JSON(200, gin.H{
				"message": "投票失败",
				"code":    400,
				"data":    string(body),
			})
		}
	})
	apiv1.POST("/userinfo", v1.AddUser)
	r.GET("api/getwalls", v1.PageWallpaper)
	return r
}

type ProxyResponse struct {
	Code int64 `json:"code"`
	Data []struct {
		IP   string `json:"ip"`
		Port int64  `json:"port"`
	} `json:"data"`
}

func getProxyAddr() string {
	resp, _ := http.Get("http://api.shenlongip.com/ip?key=lwyxsn7l&pattern=json&count=1&need=1000&protocol=1")
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	var proxy = &ProxyResponse{}
	json.Unmarshal(body, proxy)

	return "http://" + proxy.Data[0].IP + ":" + strconv.FormatInt(proxy.Data[0].Port, 10)
}
