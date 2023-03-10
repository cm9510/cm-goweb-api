package config

import (
	"io/ioutil"
	"log"
	"reflect"
	"sync"

	"gopkg.in/yaml.v3"
)

var (
	config   Items
	readOnce sync.Once
)

type Config Items

func GetConfig() *Items {
	if reflect.ValueOf(config).IsZero() {
		loadYamlConfig()
	}
	return &config
}

func loadYamlConfig() Items {
	defer func() {
		if recover := recover(); recover != nil {
			log.Fatal("pannnnnnic: ", recover)
		}
	}()

	readOnce.Do(func() {
		//读取文件
		data, err := ioutil.ReadFile("./config/app.yaml")
		if err != nil {
			panic("读取yaml错误: " + err.Error())
		}
		// 解析YAML
		if err := yaml.Unmarshal(data, &config); err != nil {
			panic("biz yaml content err: " + err.Error())
		}
	})
	return config
}
