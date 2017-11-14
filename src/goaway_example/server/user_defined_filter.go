package server

import (
	"github.com/valyala/fasthttp"
	"gateway/src/goaway/core"
	"gateway/src/goaway/constants"
	"encoding/json"
	"github.com/labstack/gommon/log"
	"strings"
)

type MsDownloadFilter struct {
	BaseUriServiceFilter
}

func (b *MsDownloadFilter) DoFilter(
	req *fasthttp.Request,
	res *fasthttp.Response,
	ctx *fasthttp.RequestCtx,
	chain *core.FilterChain) {

	chain.DoFilter(req, res, ctx)

	res.Header.SetBytesKV(
		constants.CONTENT_TYPE,
		constants.APPLICATION_X_MSDOWNLOAD)
}

func (b *MsDownloadFilter) String() string {
	return "ms-download filter (Uri: " + b.Uri + ")"
}

type TextFilter struct {
	BaseUriServiceFilter
}

func (b *TextFilter) DoFilter(
	req *fasthttp.Request,
	res *fasthttp.Response,
	ctx *fasthttp.RequestCtx,
	chain *core.FilterChain) {

	chain.DoFilter(req, res, ctx)

	res.Header.SetBytesKV(
		constants.CONTENT_TYPE,
		constants.TEXT_HTML)
}

func (b *TextFilter) String() string {
	return "text filter (Uri: " + b.Uri + ")"
}

type NoneJsonFilter struct {
	BaseUriServiceFilter
}

func (b *NoneJsonFilter) DoFilter(
	req *fasthttp.Request,
	res *fasthttp.Response,
	ctx *fasthttp.RequestCtx,
	chain *core.FilterChain) {

	chain.DoFilter(req, res, ctx)

	s := string(res.Body())
	params := make(map[string]interface{})
	err := json.Unmarshal([]byte(s), &params)
	if err != nil {
		log.Error(err)
		return
	}
	res.SetBody([]byte(params["data"].(string)))
}

func (b *NoneJsonFilter) String() string {
	return "none json filter (Uri: " + b.Uri + ")"
}

type UpdateFlightFilter struct {
	BaseUriServiceFilter
}

func (b *UpdateFlightFilter) DoFilter(
	req *fasthttp.Request,
	res *fasthttp.Response,
	ctx *fasthttp.RequestCtx,
	chain *core.FilterChain) {

	bodyStr := string(ctx.PostBody())
	params := make(map[string]string)
	params["Notify"] = strings.TrimSpace(bodyStr)
	p, _ := json.Marshal(&params)
	req.SetBody(p)

	chain.DoFilter(req, res, ctx)

}

func (b *UpdateFlightFilter) String() string {
	return "update flight filter (Uri: " + b.Uri + ")"
}

type CORSFilter struct {
	BaseUriServiceFilter
}

var allowHeaders = []byte("Content-Type,Access-Token")

func (b *CORSFilter) DoFilter(
	req *fasthttp.Request,
	res *fasthttp.Response,
	ctx *fasthttp.RequestCtx,
	chain *core.FilterChain) {

	chain.DoFilter(req, res, ctx)

	res.Header.SetBytesKV(
		constants.ACCESS_CONTROL_ALLOW_ORIGIN,
		constants.VALUE_ALL)
	res.Header.SetBytesKV(
		constants.ACCESS_CONTROL_ALOOW_HEADERS,
		[]byte(allowHeaders))

}

func (b *CORSFilter) String() string {
	return "CORS filter (Uri: " + b.Uri + ")"
}

type AirportRightsFilter struct {
	BaseUriServiceFilter
}

func (b *AirportRightsFilter) DoFilter(
	req *fasthttp.Request,
	res *fasthttp.Response,
	ctx *fasthttp.RequestCtx,
	chain *core.FilterChain) {

	chain.DoFilter(req, res, ctx)

	//这块代码有点多, 让熟悉业务的来写吧
}

func (b *AirportRightsFilter) String() string {
	return "airport rights filter (Uri: " + b.Uri + ")"
}
