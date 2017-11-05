package goaway_example

import (
	"gateway/src/model"
	"gateway/src/goaway/core"
)

type mysqlAppContext struct {
	db *model.MysqlStore
}

func NewMqlAppContext() *mysqlAppContext {
	context := mysqlAppContext{}
	context.init()
	return &context
}

func (a *mysqlAppContext) init() {
	// 获取数据库
	a.db = model.NewMysqlStore(
		"rm-wz9s84w75709ryaw7o.mysql.rds.aliyuncs.com:3306",
		"root",
		"Tc123456",
		"gateway")

}

type uriHost struct {
	Uri  string
	Host string
}

const (
	//查询服务前缀和主机(含端口)的对应关系
	SQL1 = `
		select a.Uri, concat(b.name, ':', b.port) as host
		from api a
			left join service b on a.service_id = b.service_id
		where a.status = 1`
	//查询服务前缀和过滤器名称的对应关系
	SQL2 = `
		select a.Uri, c.name as FilterName
		from api a
		  left join service b on a.service_id = b.service_id
		  left join filter c on c.api_id = a.api_id
		where
		  a.status = 1 and c.name is not null`
)

func (a *mysqlAppContext) VisitUriHosts(ctx *core.GaContext) {
	uriHosts := a.queryUriHosts()
	if len(uriHosts) > 0 {
		for _, uh := range uriHosts {
			filter := NewForwardFilter(uh.Uri, uh.Host)
			ctx.LoadFilter(filter)
		}
	}
}

func (a *mysqlAppContext) VisitUriFilters(ctx *core.GaContext) {
	filters := a.queryUriFilters()
	if len(filters) > 0 {
		for _, uf := range filters {
			filter := NewBaseServiceFilter(uf.Uri, uf.FilterName)
			ctx.LoadFilter(filter)
		}
	}
}

func (a *mysqlAppContext) queryUriHosts() []uriHost {
	rows, _ := a.db.DB.Query(SQL1)
	defer rows.Close()
	var uriHosts []uriHost
	for rows.Next() {
		uh := uriHost{}
		rows.Scan(&uh.Uri, &uh.Host)
		uriHosts = append(uriHosts, uh)
	}
	return uriHosts
}

type uriFilter struct {
	Uri        string
	FilterName string
}

func (a *mysqlAppContext) queryUriFilters() []uriFilter {
	rows, _ := a.db.DB.Query(SQL2)
	defer rows.Close()
	var uriFilters []uriFilter
	for rows.Next() {
		uf := uriFilter{}
		rows.Scan(&uf.Uri, &uf.FilterName)
		uriFilters = append(uriFilters, uf)
	}
	return uriFilters
}
