package setting

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
)

type Config struct {
	Database struct {
		User        string `yaml:"user"`
		Password    string `yaml:"password"`
		Host        string `yaml:"host"`
		Name        string `yaml:"name"`
		Port        int    `yaml:"port"`
		TablePrefix string `yaml:"table_prefix"`
	}
}

var AppConfig = &Config{}

func Setup() {

	configData, err := ioutil.ReadFile("config/app.yaml")
	if err != nil {
		log.Fatalln("读取配置文件错误：" + err.Error())
	}
	err = yaml.Unmarshal(configData, AppConfig)
	if err != nil {
		log.Fatalln("配置文件解析错误：" + err.Error())
	}
}
