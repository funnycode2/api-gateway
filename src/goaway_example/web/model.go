package web

type MResult struct {
	MPage
	Mservicelist []Mservice
}

type Mservice struct {
	Apiid   int
	Uri     string ``
	Status  int
	Desc    string
	Filters []Mfilter
}

type Mfilter struct {
	Filterid int
	Name     string
	Status   int
}

type MPage struct {
	TotalCount  int
	CurrentPage int
}
