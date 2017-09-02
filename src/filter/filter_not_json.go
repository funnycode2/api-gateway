package filter

import (
    "encoding/json"
    "log"
)

type NotJsonFilter struct{
    BaseFilter
}

func newNotJson() Filter {
    return &NotJsonFilter{}
}

func (f NotJsonFilter) Name() string {
    return FilterNotJson
}

// set string response
func (f NotJsonFilter) Post(c Context) (statusCode int, err error) {
    s := string(c.GetProxyResponse().Body())

    params := make(map[string]interface{})

    if json.Unmarshal([]byte(s), &params); err != nil {
        log.Println(err)
    }

    c.GetProxyResponse().SetBody([]byte(params["data"].(string)))
    return
}