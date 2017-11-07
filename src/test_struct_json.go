package main

import (
	"gateway/src/goaway_example/web"
	"encoding/json"
)

func main() {
	ms := web.Mservice{
		Apiid:  1,
		Uri:    "a",
		Desc:   "d",
		Status: 2,
		Filters: []web.Mfilter{
			web.Mfilter{
				Filterid: 2,
				Name:     "mf",
				Status:   0,
			},
		},
	}
	bytes, _ := json.Marshal(ms)
	println(string(bytes))
}
