package ext

import (
	"github.com/valyala/fasthttp"
	"gateway/src/goaway/core"
	"gateway/src/util"
)

type OauthFilter struct{}

var (
	oauthAddr = util.GetConfigCenterInstance().ConfProperties["oauth_center"]["oauth_addr"]
)

func (OauthFilter) Matches(url string) bool {
	return true
}

func (OauthFilter) DoFilter(
	req *fasthttp.Request,
	res *fasthttp.Response,
	ctx *fasthttp.RequestCtx,
	chain *core.FilterChain) {

	/*accessToken := req.URI().QueryArgs().Peek("access_token")
	if nil == accessToken {
		res.SetStatusCode(http.StatusUnauthorized)
		return
	}
	//res, err := h.fastHTTPClient.Do(outReq, config.TConfig.OauthHost
	fasthttp.Re
	authRes, err := http.Get(oauthAddr + "/user/getUser?access_token=" + string(accessToken))

	if err != nil {
		log.Println(err)
		return false, err
	}
	defer res.Body.Close()

	if res.StatusCode == 200 {
		var oauthResult map[string]interface{}
		body, _ := ioutil.ReadAll(res.Body)
		json.Unmarshal(body, &oauthResult)

		// 设置user_id
		//req.Header.Add("user_id", strconv.Itoa(int(oauthResult["user_id"].(float64))))
		//req.PostArgs().Add("user_id", oauthResult["user_id"].(string))
		//req.PostArgs().Add("client_id", oauthResult["client_id"].(string))
		if clientId, ok := oauthResult["client_id"].(string); ok {
			req.PostArgs().Add("client_id", clientId)
		}

		if userId, ok := oauthResult["user_id"].(string); ok {
			req.PostArgs().Add("user_id", userId)
		} else if userId, ok := oauthResult["user_id"].(float64); ok {
			req.PostArgs().Add("user_id", strconv.Itoa(int(userId)))
		}

		return true, nil
	} else {
		return false, err
	}*/
}
