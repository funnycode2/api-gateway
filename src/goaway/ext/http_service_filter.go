package ext

import (
	"strings"
	"github.com/valyala/fasthttp"
	"gateway/src/goaway/core"
	"github.com/labstack/gommon/log"
	"gateway/src/goaway/util"
)

type basicServiceFilter struct {
	core.BaseFilter
	prefix             string //服务名前缀
	targetPrefix       string //目标前缀
	targetHostWithPort string //目标主机
}

//基本服务过滤器, 提供简单的前缀匹配和目标前缀重写,主机端口重写
func NewBasicServiceFilter(
	prefix string,
	targetPrefix string,
	targetHostWithPort string) *basicServiceFilter {
	normalizePrefix, e := util.NormalizeUri(prefix)
	if e != nil {
		log.Panic(e)
	}
	normalizeTargetPrefix, e := util.NormalizeUri(targetPrefix)
	if e != nil {
		log.Panic(e)
	}
	if !util.MatchHost(targetHostWithPort) {
		log.Panicf("Invalid host with port: %s, only 'service', 'service:8080', '12.12.12.12', 'user.service:1080' allowed", targetHostWithPort)
	}
	return &basicServiceFilter{
		prefix:             normalizePrefix,
		targetPrefix:       normalizeTargetPrefix,
		targetHostWithPort: targetHostWithPort,
	}
}

func (f *basicServiceFilter) Matches(url string) bool {
	return strings.HasPrefix(url, f.prefix)
}

func (f *basicServiceFilter) DoFilter(
	req *fasthttp.Request,
	res *fasthttp.Response,
	ctx *fasthttp.RequestCtx,
	chain *core.FilterChain) {
	req.SetHost(f.targetHostWithPort)
	reqURI := string(req.Header.RequestURI())
	targetURI := strings.Replace(reqURI, f.prefix, f.targetPrefix, -1)
	req.Header.SetRequestURI(targetURI)
	chain.DoFilter(req, res, ctx)
}

func (f *basicServiceFilter) OnDestroy() {
	log.Info("destroying basic service filter")
}
