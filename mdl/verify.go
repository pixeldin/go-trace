package mdl

import "net/http"

func VerifyHttp(request *Request, response *http.Response) (code int, isSucceed bool) {
	defer response.Body.Close()
	code = response.StatusCode
	isSucceed = code == HTTPOK
	return
}
