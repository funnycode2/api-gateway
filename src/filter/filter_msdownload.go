package filter

type MsdownloadFilter struct{
    BaseFilter
}

func newMsdownloadFilter() Filter {
    return &MsdownloadFilter{}
}

func (f MsdownloadFilter) Name() string {
    return FilterMsdownload
}

// 有待测试的头
func (f MsdownloadFilter) Post(c Context) (statusCode int, err error) {
    c.GetOriginRequestCtx().Response.Header.Set("Content-Type", "application/x-msdownload")
    return f.BaseFilter.Post(c)
}