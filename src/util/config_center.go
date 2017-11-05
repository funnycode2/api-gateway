package util

import (
	"sync"
	"code.aliyun.com/wyunshare/wyun-zookeeper/go-client/src/conf_center"
	"fmt"
	"github.com/ironcity/goconf"
	"github.com/labstack/gommon/log"
)

var m conf_center.AppProperties
var once sync.Once

func GetConfigCenterInstance() conf_center.AppProperties {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("ERROR!! ,配置中心连接失败 ,", err)
		}
	}()

	once.Do(func() {
		m = LoadYmlDefault()
		m.Init()
	})
	return m
}

func LoadYmlDefault() conf_center.AppProperties {

	// 加载文件
	c, err := goconf.ReadConfigFile("conf.yml")

	if err != nil {
		log.Panic(err)
	}

	m := conf_center.New("local")

	prop := make(map[string]map[string]string)
	prop["jdbc"] = c.GetSectionNode("jdbc")
	prop["oauth_center"] = c.GetSectionNode("oauth_center")
	prop["kafaka"] = c.GetSectionNode("kafaka")
	prop["zookeeper"] = c.GetSectionNode("zookeeper")

	m.ConfProperties = prop
	return m

}
