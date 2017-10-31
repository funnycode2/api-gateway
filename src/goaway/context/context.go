package context

import (
	"gateway/src/goaway/filter"
	"github.com/labstack/gommon/log"
	"reflect"
)

var (
	Context = &context{
		Filters: []filter.Filter{&filter.CoreFilter{}},
	}
)

type context struct {
	Filters []filter.Filter
}

func (c *context) AddFilter(filter filter.Filter) {
	if filter != nil {
		log.Info("Adding filter: ", reflect.ValueOf(filter).Type().String())
		c.Filters = append(c.Filters, filter)
	}
}

func init() {
	log.Info("go-away context starting ...")

	Context.AddFilter(&filter.GapFilter{})

	log.Info("go-away context finish loading .")
}
