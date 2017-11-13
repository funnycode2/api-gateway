package sql

import (
	t "text/template"
	"text/template"
	"bytes"
	"github.com/labstack/gommon/log"
)

func NewSql(tmp *t.Template, params interface{}) string {
	var buffer bytes.Buffer
	tmp.Execute(&buffer, params)
	return buffer.String()
}

func newTemplate(text string) *t.Template {
	tpl, err := template.New("").Parse(text)
	if err != nil {
		log.Error(err)
	}
	return tpl
}

//var SELECT0, _ = template.New("SELECT0").Parse(`
var SELECT0 = newTemplate(`
	select
		a.api_id as Apiid,
		a.Uri, a.Status,
		a.display_name as 'Desc',
		a.service_id as ServiceId,
		s.Name, s.Port
	from api a
	left join service s on a.service_id = s.service_id
	where 1 = 1
		{{if and .Uri .Desc}}
	    and (a.uri like '%{{.Uri}}%' or a.display_name like '%{{.Desc}}%')
		{{else}}
			{{if .Uri}}
			and a.uri like '%{{.Uri}}%'
			{{end}}
			{{if .Desc}}
			a.display_name like '%{{.Desc}}%'
			{{end}}
		{{end}}
	order by a.uri
`)

var UPDATE0 = newTemplate(`
	update api
		set status = {{.Status}},
		uri = '{{.Uri}}',
		display_name = '{{.Desc}}',
		service_id = {{.ServiceId}}
	where
		api_id = {{.Apiid}}
`)
