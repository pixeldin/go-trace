package net

import (
	"go-trace/mdl"
	"go-trace/utils"
	"log"
	"net/http"
	"os"
	"time"
)

var logger = log.New(os.Stderr, "client:", 0)

func Send(req *mdl.Request) (succeed bool, errCode int, reqTime uint64) {
	client := &http.Client{
		Timeout: req.Timeout,
	}

	// built basic request
	request, err := http.NewRequest(req.Method, req.Url, req.GetBody())
	if err != nil {
		return
	}

	//if _, ok := req.Headers["Content-Type"]; !ok {
	//	if req.Headers == nil {
	//		req.Headers = make(map[string]string)
	//	}
	//	req.Headers["Content-Type"] = "application/x-www-form-urlencoded; charset=utf-8"
	//}

	// fill header
	for k, v := range req.Headers {
		request.Header.Set(k, v)
	}

	// default as utf-8
	if request.Header.Get("Content-Type") == "" {
		request.Header.Set("Content-Type", "application/x-www-form-urlencoded; charset=utf-8")
		//	req.Headers["Content-Type"] = "application/x-www-form-urlencoded; charset=utf-8"
	}

	// timecost
	begin := time.Now()
	resp, err := client.Do(request)
	reqTime = uint64(utils.DiffNano(begin))
	if err != nil {
		// TODO: add target of output
		logger.Printf("Request err: %v", err)
		errCode = mdl.RequestErr
		return
	}

	// TODO: verify resp
	errCode, succeed = mdl.VerifyHttp(req, resp)
	return
}
