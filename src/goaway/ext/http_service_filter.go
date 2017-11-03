package ext

import (
	"strings"
	"github.com/valyala/fasthttp"
	"gateway/src/goaway/core"
	"errors"
	"github.com/labstack/gommon/log"
)

type commonServiceFilter struct{
	prefix string //服务名前缀
}

func NewCommonServiceFilter(prefix string) *commonServiceFilter {
	normalizeUrl, e := normalizeUrl(prefix)
	if e != nil {
		log.Panic(e)
	}
	return &commonServiceFilter{
		prefix: normalizeUrl,
	}
}

func (f *commonServiceFilter) Matches(url string) bool {
	return strings.HasPrefix(url, f.prefix)
}

func (f *commonServiceFilter) DoFilter(
	req *fasthttp.Request,
	res *fasthttp.Response,
	ctx *fasthttp.RequestCtx,
	chain *core.FilterChain) {
	req.SetHost("localhost:8080")
	chain.DoFilter(req, res, ctx)
}

var errorEmptyUrl = errors.New("empty url not allowed")

//将url正则化如: url := "\\aaa\\\\bb/\\" 转化成  /aaa/bb
func normalizeUrl(url string) (string, error) {
	url = strings.Replace(url, "\\", "/", -1)
	if len(url) == 0 {
		return string(nil), errorEmptyUrl
	}
	splits := strings.Split(url, "/")
	var normalized string
	for _, split := range splits {
		if trimmed := strings.TrimSpace(split); len(trimmed) > 0 {
			normalized = normalized + "/" + trimmed
		}
	}
	if len(normalized) == 0 {
		return string(nil), errorEmptyUrl
	}
	return normalized, nil
}