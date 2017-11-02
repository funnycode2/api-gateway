package ext

import (
	"github.com/valyala/fasthttp"
	"gateway/src/goaway/core"
	"gateway/src/util"
	"net/http"
	"encoding/json"
	"strconv"
	"gateway/src/goaway/constants"
	"github.com/labstack/gommon/log"
)

type OauthFilter struct{}

var (
	oauthAddr = util.GetConfigCenterInstance().ConfProperties["oauth_center"]["oauth_addr"] +
		"/user/getUser?access_token="
)

func (f *OauthFilter) Matches(url string) bool {
	return true
}

func (f *OauthFilter) DoFilter(
	req *fasthttp.Request,
	res *fasthttp.Response,
	ctx *fasthttp.RequestCtx,
	chain *core.FilterChain) {

	accessToken := req.URI().QueryArgs().Peek("access_token")
	if nil == accessToken {
		res.SetStatusCode(http.StatusUnauthorized)
		return
	}

	authReq := fasthttp.AcquireRequest()
	authReq.Reset()
	defer fasthttp.ReleaseRequest(authReq)
	authRes := fasthttp.AcquireResponse()
	authRes.Reset()
	defer fasthttp.ReleaseResponse(authRes)

	//请求后端认证服务器
	authReq.SetRequestURI(oauthAddr + string(accessToken))
	authReq.Header.SetMethodBytes(constants.GET)

	fasthttp.Do(authReq, authRes)

	if authRes.StatusCode() == 200 {
		var oauthResult map[string]interface{}
		json.Unmarshal(authRes.Body(), &oauthResult)

		if clientId, ok := oauthResult["client_id"].(string); ok {
			req.PostArgs().Add("client_id", clientId)
		} else {
			log.Error("认证错误! 请求:\n" + req.String())
			log.Error("该请求的认证服务器响应为:\n" + authRes.String())
			res.SetStatusCode(fasthttp.StatusUnauthorized)
			return
		}

		if userId, ok := oauthResult["user_id"].(string); ok {
			req.PostArgs().Add("user_id", userId)
		} else if userId, ok := oauthResult["user_id"].(float64); ok {
			req.PostArgs().Add("user_id", strconv.Itoa(int(userId)))
		} else {
			log.Error("认证错误! 请求:\n" + req.String())
			log.Error("该请求的认证服务器响应为:\n" + authRes.String())
			res.SetStatusCode(fasthttp.StatusUnauthorized)
			return
		}

		//认证通过, 放通该请求
		chain.DoFilter(req, res, ctx)

	} else {
		res.SetStatusCode(authRes.StatusCode())
		return
	}
}
