package context

import (
	"gateway/src/goaway/filter"
	"gateway/src/goaway/mapping"
	"github.com/labstack/gommon/log"
	"reflect"
	"gateway/src/goaway/handler"
)

var (
	Context = &context{
		Filters:    []filter.Filter{},
		Mappings:   make([]mapping.Mapping, 0),
		Handlers:   make([]handler.Handler, 0),
	}
)

type context struct {
	Filters    []filter.Filter
	Mappings   []mapping.Mapping
	Handlers   []handler.Handler
}

func (c *context) AddFilter(filter filter.Filter) {
	if filter != nil {
		log.Info("Adding filter: ", reflect.ValueOf(filter).Type().String())
		c.Filters = append(c.Filters, filter)
	}
}

func (c *context) AddMapping(mapping mapping.Mapping) {
	if mapping != nil {
		log.Info("Adding mapping: ", reflect.ValueOf(mapping).Type().String())
		c.Mappings = append(c.Mappings, mapping)
	}
}

func (c *context) AddHandler(handler handler.Handler) {
	if handler != nil {
		log.Info("Adding handler: ", reflect.ValueOf(handler).Type().String())
		c.Handlers = append(c.Handlers, handler)
	}
}

func init() {
	Context.AddFilter(&filter.DefaultFilter{})
	Context.AddFilter(&filter.OauthFilter{})

	Context.AddMapping(&mapping.DefaultMapping{})

	Context.AddHandler(&handler.ProxyHandler2{})
}
