package constants

var (
	GET    = []byte("GET")
	POST   = []byte("POST")
	PUT    = []byte("PUT")
	DELETE = []byte("DELETE")
	HEAD   = []byte("HEAD")

	CONTENT_TYPE = []byte("Content-Type")

	APPLICATION_X_MSDOWNLOAD = []byte("application/x-msdownload")
	TEXT_HTML                = []byte("text/html;charset=utf-8")

	VALUE_ALL                    = []byte("*")
	ACCESS_CONTROL_ALLOW_ORIGIN  = []byte("Access-Control-Allow-Origin")
	ACCESS_CONTROL_ALOOW_HEADERS = []byte("Access-Control-Allow-Headers")
)
