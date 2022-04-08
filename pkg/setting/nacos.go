package setting

import (
	"fmt"
	"github.com/nacos-group/nacos-sdk-go/clients"
	"github.com/nacos-group/nacos-sdk-go/common/constant"
	"github.com/nacos-group/nacos-sdk-go/vo"
	"os"
)

func GetConfig(dataId string, group string) string {
	var endpoint = "acm.aliyun.com"
	var namespaceId = os.Getenv("aliyun-namespaceId")
	var accessKey = os.Getenv("aliyun-accessKey")
	var secretKey = os.Getenv("aliyun-secretKey")

	fmt.Println(fmt.Sprintf("namespaceId：%s", namespaceId))
	fmt.Println(fmt.Sprintf("accessKey：%s", accessKey))
	fmt.Println(fmt.Sprintf("secretKey：%s", secretKey))

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

	// Get plain content from ACM.
	content, _ := configClient.GetConfig(vo.ConfigParam{
		DataId: dataId,
		Group:  group})

	return content
}
