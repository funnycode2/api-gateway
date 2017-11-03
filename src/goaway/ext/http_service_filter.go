package ext

import (
	"strings"
	"github.com/valyala/fasthttp"
	"gateway/src/goaway/core"
	"errors"
	"github.com/labstack/gommon/log"
)

type basicServiceFilter struct {
	prefix             string //服务名前缀
	targetPrefix       string //目标前缀
	targetHostWithPort string //目标主机
}

//基本服务过滤器, 提供简单的前缀匹配和目标前缀重写,主机端口重写
func NewBasicServiceFilter(
	prefix string,
	targetPrefix string,
	targetHostWithPort string) *basicServiceFilter {
	normalizePrefix, e := normalizeUrl(prefix)
	if e != nil {
		log.Panic(e)
	}
	normalizeTargetPrefix, e := normalizeUrl(targetPrefix)
	if e != nil {
		log.Panic(e)
	}
	checkedTargetHost := targetHostWithPort
	return &basicServiceFilter{
		prefix:             normalizePrefix,
		targetPrefix:       normalizeTargetPrefix,
		targetHostWithPort: checkedTargetHost,
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

var emptyUrlError = errors.New("empty url not allowed")

//将url正则化如: url := "\\aaa\\\\bb/\\" 转化成  /aaa/bb
func normalizeUrl(url string) (string, error) {
	url = strings.Replace(url, "\\", "/", -1)
	if len(url) == 0 {
		return "", emptyUrlError
	}
	splits := strings.Split(url, "/")
	var normalized string
	for _, split := range splits {
		if trimmed := strings.TrimSpace(split); len(trimmed) > 0 {
			normalized = normalized + "/" + trimmed
		}
	}
	if len(normalized) == 0 {
		return "", emptyUrlError
	}
	return normalized, nil
}
