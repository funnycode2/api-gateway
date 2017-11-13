package main

import (
	"gateway/src/goaway_example/web"
	"bytes"
	"text/template"
	t "text/template"
)

func NewSql(tmp *t.Template, params interface{}) string {
	var buffer bytes.Buffer
	tmp.Execute(&buffer, params)
	return buffer.String()
}

var TEMPLATE = template.New("SQL")

func main() {
	var UPDATE0, _ = TEMPLATE.Parse(`
	update api set
		status = {{.Status}},
		uri = '{{.Uri}}',
		display_name = '{{.Desc}}',
		service_id = {{.ServiceId}}
	where
		api_id = {{.Apiid}}`)
	params := web.Mservice{
		//Uri: "/aaa",
		//Desc: "没有描述",
	}
	print(NewSql(UPDATE0, params))
}
