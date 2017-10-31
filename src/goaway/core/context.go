package core

import (
	"github.com/labstack/gommon/log"
	"reflect"
)

type Context struct {
	Filters []Filter
}

func (c *Context) AddFilter(filter Filter) {
	if filter != nil {
		log.Info("Adding filter: ", reflect.ValueOf(filter).Type().String())
		c.Filters = append(c.Filters, filter)
	}
}