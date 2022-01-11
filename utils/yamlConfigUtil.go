package utils

import (
	"fmt"
	"io/ioutil"
	"os"

	"gopkg.in/yaml.v2"
)

// YamlConfig 配置文件对应结构体
type YamlConfig struct {
	Mysql struct {
		Username string `yaml:"username"`
		Password string `yaml:"password"`
		Url      string `yaml:"url"`
		Dbname   string `yaml:"dbname"`
	} `yaml:"mysql"`
}

// ReadConfig 去读取配置文件
func ReadConfig(fileName string) *YamlConfig {
	f, err := os.OpenFile(fileName, os.O_RDONLY, 0600)
	defer f.Close()
	if err != nil {
		fmt.Println("配置文件打开失败" + err.Error())
	} else {
		contentByte, err := ioutil.ReadAll(f)
		if err != nil {
			fmt.Println("配置文件读取失败")
			fmt.Println(err)
			return nil
		}

		yamlConfig := YamlConfig{}

		err = yaml.Unmarshal([]byte(contentByte), &yamlConfig)
		if err != nil {
			fmt.Println("解析yaml配置文件失败" + err.Error())
		}

		return &yamlConfig

	}
	return nil
}

// InitConfig 初始化读取配置文件
func InitConfig() *YamlConfig {

	return ReadConfig("configs/config.yaml")
}
