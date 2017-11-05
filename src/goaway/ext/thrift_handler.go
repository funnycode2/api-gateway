package ext

import (
	"github.com/valyala/fasthttp"
	"code.aliyun.com/wyunshare/thrift-server/gen-go/server"
	"encoding/json"
	"github.com/labstack/gommon/log"
)

type thriftHandler struct {
	serviceName string
}

func NewThriftHandler() *thriftHandler {
	return &thriftHandler{}
}

func (h *thriftHandler) Do(req *fasthttp.Request, res *fasthttp.Response) {
	thriftReq := server.NewRequest()
	thriftReq.ServiceName = h.serviceName

	// 解析参数，转化成json格式
	params := make(map[string]interface{})

	if nil != req.Body() {
		if err := json.Unmarshal(req.Body(), &params); nil != err {
			log.Errorf("thrift handler request body json parse error: %s", string(req.Body()))
			return
		}
	}

	var f = func(k []byte, v []byte) {
		params[string(k)] = string(v)
	}
	req.URI().QueryArgs().VisitAll(f)
	req.PostArgs().VisitAll(f)

	// 转化成json
	delete(params, "access_token")
	thriftReq.ParamJSON, _ = json.Marshal(params)

	//TODO 完成实现
}
