package setting

import (
	"fmt"
	"github.com/nacos-group/nacos-sdk-go/clients"
	"github.com/nacos-group/nacos-sdk-go/common/constant"
	"github.com/nacos-group/nacos-sdk-go/vo"
	"gopkg.in/yaml.v2"
	"log"
	"os"
	"wechat-bot-server/models"
)

func GetConfig(dataId string, group string) string {
	var endpoint = "acm.aliyun.com"
	var namespaceId = os.Getenv("aliyun-namespaceId")
	var accessKey = os.Getenv("aliyun-accessKey")
	var secretKey = os.Getenv("aliyun-secretKey")

	clientConfig := constant.ClientConfig{
		Endpoint:       endpoint + ":8080",
		NamespaceId:    namespaceId,
		AccessKey:      accessKey,
		SecretKey:      secretKey,
		TimeoutMs:      5 * 1000,
		ListenInterval: 30 * 1000,
	}

	// Initialize client.
	configClient, _ := clients.CreateConfigClient(map[string]interface{}{
		"clientConfig": clientConfig,
	})

	// 监听配置
	configClient.ListenConfig(vo.ConfigParam{
		DataId: dataId,
		Group:  group,
		OnChange: func(namespace, group, dataId, data string) {
			err := yaml.Unmarshal([]byte(data), AppConfig)
			if err != nil {
				log.Fatalln("配置文件解析错误：" + err.Error())
			}
			models.Setup()
			fmt.Println("ListenConfig group:" + group + ", dataId:" + dataId + ", data:" + data)
		},
	})

	// Get plain content from ACM.
	content, _ := configClient.GetConfig(vo.ConfigParam{
		DataId: dataId,
		Group:  group})

	return content
}
