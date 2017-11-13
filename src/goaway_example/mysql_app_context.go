package goaway_example

import (
	"gateway/src/goaway/core"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"gateway/src/goaway_example/web"
	s "gateway/src/goaway_example/web/sql"
	"strconv"
	"fmt"
	"errors"
)

type mysqlAppContext struct {
	db *sql.DB
}

func NewMqlAppContext() *mysqlAppContext {
	context := mysqlAppContext{}
	context.init()
	return &context
}

const connUrl = "root:Tc123456@tcp(rm-wz9s84w75709ryaw7o.mysql.rds.aliyuncs.com:3306)/gateway"

func (a *mysqlAppContext) init() {
	// 获取数据库
	db, _ := sql.Open("mysql", connUrl)
	db.SetMaxOpenConns(100) //最大连接数
	db.SetMaxIdleConns(50)  //最大闲置数
	db.Ping()
	a.db = db
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
	SQL4 = `
		  update filter set status = %d where filter_id = %d`
	SQL5 = `
		  select distinct name from filter`
	SQL6 = `
		  insert into filter (api_id, name, status) values (%d, '%s', 1)`
	SQL7 = `
		  select count(0) from api where uri = '%s'`
	SQL8 = `
		  insert into api (display_name, uri, service_id, status) values ('%s', '%s', %d, 1)`
	SQL9 = `
		  select api_id from api where uri = '%s'`
	SQL10 = `
		select s.service_id as ServiceId, s.name, s.port
		from service s ORDER by s.name, s.port`
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
	rows, _ := a.db.Query(SQL1)
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
	rows, _ := a.db.Query(SQL2)
	defer rows.Close()
	var uriFilters []uriFilter
	for rows.Next() {
		uf := uriFilter{}
		rows.Scan(&uf.Uri, &uf.FilterName)
		uriFilters = append(uriFilters, uf)
	}
	return uriFilters
}

const PAGE_SIZE = 50

func (a *mysqlAppContext) QueryService(
	uri string,
	desc string,
	currentPage int) *web.MResult {

	sqltext := s.NewSql(s.SELECT0, &web.Mservice{
		Uri:  uri,
		Desc: desc,
	})

	//查询总条数
	countSql := fmt.Sprintf("select count(0) from (%s) t", sqltext)
	countRow, _ := a.db.Query(countSql)
	defer countRow.Close()
	mPage := web.MPage{}
	if countRow.Next() {
		countRow.Scan(&mPage.TotalCount)
	}

	//计算设置分页的参数
	totalPage := (mPage.TotalCount + PAGE_SIZE - 1) / PAGE_SIZE
	if totalPage < currentPage {
		currentPage = totalPage
	}
	if currentPage < 1 {
		currentPage = 1
	}
	mPage.CurrentPage = currentPage
	sqltext += " limit " + strconv.Itoa((currentPage-1)*PAGE_SIZE) + ", " + strconv.Itoa(PAGE_SIZE)

	//获取服务查询结果
	rows, _ := a.db.Query(sqltext)
	defer rows.Close()
	var services []web.Mservice
	for rows.Next() {
		ms := web.Mservice{}
		rows.Scan(&ms.Apiid, &ms.Uri, &ms.Status, &ms.Desc, &ms.ServiceId, &ms.Name, &ms.Port)
		services = append(services, ms)
	}

	//关联过滤器
	for index, _ := range services {
		rows, _ := a.db.Query("select a.filter_id as filterid, a.name, a.status from filter a where a.api_id = " + strconv.Itoa(services[index].Apiid))
		defer rows.Close()
		for rows.Next() {
			mf := web.Mfilter{}
			rows.Scan(&mf.Filterid, &mf.Name, &mf.Status)
			services[index].Filters = append(services[index].Filters, mf)
		}
	}

	//获取服务查询结果
	rows0, _ := a.db.Query(SQL5)
	defer rows0.Close()
	var allFilterNames []string
	for rows0.Next() {
		var name string
		rows0.Scan(&name)
		allFilterNames = append(allFilterNames, name)
	}

	rows1, _ := a.db.Query(SQL10)
	defer rows1.Close()
	var allHosts []web.Mhost
	for rows1.Next() {
		var host web.Mhost
		rows1.Scan(&host.ServiceId, &host.Name, &host.Port)
		allHosts = append(allHosts, host)
	}

	result := web.MResult{}
	result.MPage = mPage
	result.Mservicelist = services
	result.AllFilterNames = &allFilterNames
	result.AllHosts = &allHosts

	return &result
}

func (a *mysqlAppContext) UpdateService(mservice *web.Mservice) error {
	if mservice.New {
		//如果是新的, 则需要插入, 并在插入前需要校验是否重复
		var uriCount int
		rows, _ := a.db.Query(fmt.Sprintf(SQL7, mservice.Uri))
		defer rows.Close()
		for rows.Next() {
			rows.Scan(&uriCount)
		}
		if uriCount > 0 {
			return errors.New("uri already exists")
		}
		a.db.Exec(fmt.Sprintf(SQL8, mservice.Desc, mservice.Uri, mservice.ServiceId))
		rows0, _ := a.db.Query(fmt.Sprintf(SQL9, mservice.Uri))
		defer rows0.Close()
		if rows0.Next() {
			rows0.Scan(&mservice.Apiid)
		} else {
			return errors.New("no uri = " + mservice.Uri + " found")
		}
	} else {
		//执行修改
		update := s.NewSql(s.UPDATE0, mservice)
		a.db.Exec(update)
	}

	mfilters := mservice.Filters
	if len(mfilters) > 0 {
		for _, fitler := range mfilters {
			if fitler.New {
				//前台新添加的过滤器需要插入
				a.db.Exec(fmt.Sprintf(SQL6, mservice.Apiid, fitler.Name))
			} else {
				//老的过滤器需要更新
				a.db.Exec(fmt.Sprintf(SQL4, fitler.Status, fitler.Filterid))
			}

		}
	}
	return nil
}
