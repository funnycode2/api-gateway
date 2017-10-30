package business

import (
	"code.aliyun.com/wyunshare/wyun-zookeeper/go-client/src/conf_center"
	"github.com/ironcity/goconf"
	"log"
)

func LoadYmlDefault() conf_center.AppProperties {

	// 加载文件
	c, err := goconf.ReadConfigFile("conf.yml")

	if err != nil {
		log.Fatal(err)
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
