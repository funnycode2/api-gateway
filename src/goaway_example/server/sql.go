package server

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

//搜索服务信息
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

//查询服务前缀和主机(含端口)的对应关系
var SELECT1 = newTemplate(`
	select a.Uri, concat(b.name, ':', b.port) as Host
	from api a
		left join service b on a.service_id = b.service_id
	where a.status = 1
`)

//查询服务前缀和过滤器名称的对应关系
var SELECT2 = newTemplate(`
	select a.Uri, c.name as FilterName
	from api a
	  left join service b on a.service_id = b.service_id
	  left join filter c on c.api_id = a.api_id
	where
	  a.status = 1 and c.name is not null
`)

var SELECT5 = newTemplate(`
	select distinct name from filter
`)

var SELECT7 = newTemplate(`
	select count(0) from api where uri = '{{.Uri}}'
`)

var SELECT9 = newTemplate(`
	select api_id from api where uri = '{{.Uri}}'
`)

var SELECT10 = newTemplate(`
	select s.service_id as ServiceId, s.name, s.port
	from service s ORDER by s.name, s.port
`)

var SELECT11 = newTemplate(`
	select
		a.filter_id as filterid, a.name, a.status
	from filter a
	where a.api_id = {{.Apiid}}
`)

//更新服务信息
var UPDATE0 = newTemplate(`
	update api
		set status = {{.Status}},
		uri = '{{.Uri}}',
		display_name = '{{.Desc}}',
		service_id = {{.ServiceId}}
	where
		api_id = {{.Apiid}}
`)

var UPDATE4 = newTemplate(`
	update filter set status = {{.Status}} where filter_id = {{.Filterid}}
`)

var INSERT6 = newTemplate(`
	insert into filter (api_id, name, status) values ({{.Apiid}}, '{{.Name}}', 1)
`)

var INSERT8 = newTemplate(`
	insert into api
		(display_name, uri, service_id, status)
	values
		('{{.Desc}}', '{{.Uri}}', {{.ServiceId}}, 1)
`)
