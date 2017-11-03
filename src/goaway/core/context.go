package core

import (
	"github.com/labstack/gommon/log"
	"reflect"
)

type context struct {
	filters []Filter
}

var gaCoreFilter = &coreFilter{}

func NewContext() *context {
	return &context{
		filters: []Filter{gaCoreFilter},
	}
}

//加载过滤器到上下文中, 需要确保
//1. 核心过滤器总是存在, 并且总是为第一个过滤器
//2. 对于已经存在于上下文中的过滤器, 忽略添加
func (c *context) LoadFilter(filter Filter) {
	filters := c.filters
	//确保核心过滤器总是存在并且在第一个
	if filters == nil || filters[0] != gaCoreFilter {
		log.Panic("CoreFilter does not exist!")
	}
	if filter != nil {
		//添加非重复的过滤器
		for _, f := range filters {
			if f == filter {
				log.Infof("Ignore filter for duplication: %s",
					reflect.ValueOf(filter).Type().String())
				return
			}
		}
		log.Infof("Adding filter: %s",
			reflect.ValueOf(filter).Type().String())
		c.filters = append(c.filters, filter)
	}
}

func (c *context) onDestroy() {
	for _, f := range c.filters {
		log.Infof("Destorying filter : %s",
			reflect.ValueOf(f).Type().String())
		f.OnDestroy()
	}
}

//为保证不修改内部filter数组, 牺牲效率和内存,返回拷贝
func (c *context) Filters() []Filter {
	return c.filters
}
