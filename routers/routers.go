package routers

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"time"
	v1 "wechat-bot-server/routers/api/v1"
)

func InitRouter() *gin.Engine {
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())
	apiv1 := r.Group("/api/v1")

	apiv1.GET("/userinfo", v1.GetUser)
	r.GET("/vote", func(context *gin.Context) {

		proxyAddr := getProxyAddr()
		log.Println("神龙代理IP：" + proxyAddr)
		proxy, err := url.Parse(proxyAddr)
		if err != nil {
			log.Fatal(err)
		}

		netTransport := &http.Transport{
			Proxy:                 http.ProxyURL(proxy),
			MaxIdleConnsPerHost:   10,
			ResponseHeaderTimeout: time.Second * time.Duration(5),
		}

		client := &http.Client{
			Timeout:   time.Second * 10,
			Transport: netTransport,
		}
		req, _ := http.NewRequest("GET", "http://nbd332.wh.changqingmall.cn/Home/index.php?m=Index&a=vote&vid=610717&id=12231&tp=", nil)
		//req.Header.Add("X-Forwarded-For", genIpaddr())
		resp, _ := client.Do(req)
		if resp != nil {
			defer resp.Body.Close()
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
	return r
}
func genIpaddr() string {
	rand.Seed(time.Now().Unix())
	ip := fmt.Sprintf("%d.%d.%d.%d", rand.Intn(255), rand.Intn(255), rand.Intn(255), rand.Intn(255))
	return ip
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
