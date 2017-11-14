package web

type MResult struct {
	MPage
	Mservicelist   []Mservice
	AllFilterNames *[]string //所有的过滤器名称
	AllHosts       *[]Mhost
}

type Mservice struct {
	Mhost
	Apiid   int
	Uri     string
	Desc    string
	Status  int
	New     bool
	Filters []Mfilter
}

type Mhost struct {
	ServiceId int
	Name      string
	Port      int
}

type Mfilter struct {
	Apiid    int
	Filterid int
	Name     string
	Status   int
	New      bool //是否是新过滤器, 新过滤器没有id, 需要新增
}

type MPage struct {
	TotalCount  int
	CurrentPage int
}
