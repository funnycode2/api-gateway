package web

type MResult struct {
	MPage
	Mservicelist   []Mservice
	AllFilterNames *[]string //所有的过滤器名称
}

type Mservice struct {
	Apiid   int
	Uri     string
	Desc    string
	Status  int
	New     bool
	Filters []Mfilter
}

type Mfilter struct {
	Filterid int
	Name     string
	Status   int
	New      bool //是否是新过滤器, 新过滤器没有id, 需要新增
}

type MPage struct {
	TotalCount  int
	CurrentPage int
}
