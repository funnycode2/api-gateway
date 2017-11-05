package core

import (
	"github.com/labstack/gommon/log"
	"reflect"
)

type GaContext struct {
	filters []Filter
}

var gaCoreFilter = &coreFilter{}

func NewContext() *GaContext {
	return &GaContext{
		filters: []Filter{gaCoreFilter},
	}
}

//加载过滤器到上下文中, 需要确保
//1. 核心过滤器总是存在, 并且总是为第一个过滤器
//2. 对于已经存在于上下文中的过滤器, 忽略添加
func (c *GaContext) LoadFilter(filter Filter) {
	filters := c.filters
	//确保核心过滤器总是存在并且在第一个
	if filters == nil || filters[0] != gaCoreFilter {
		c.filters = append([]Filter{gaCoreFilter}, filters...)
	}
	if filter == nil {
		log.Error("Ignoring nil filter")
		return
	}
	//添加非重复的过滤器
	for _, f := range filters {
		if f == filter {
			log.Infof("Ignoring filter for duplication %s(%s)",
				reflect.ValueOf(filter).Type().String(),
				filter)
			return
		}
	}
	log.Infof("Adding filter %s(%s)",
		reflect.ValueOf(filter).Type().String(),
		filter)
	c.filters = append(c.filters, filter)
}

func (c *GaContext) onDestroy() {
	for _, f := range c.filters {
		log.Infof("Destroying filter %s",
			reflect.ValueOf(f).Type().String(),
			f)
		f.OnDestroy()
	}
}

//为保证不修改内部filter数组, 牺牲效率和内存,返回拷贝
func (c *GaContext) Filters() []Filter {
	filters := c.filters
	copyFilters := make([]Filter, len(filters))
	copy(copyFilters, filters)
	return copyFilters
}
