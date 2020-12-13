package net

import (
	"go-trace/mdl"
	"go-trace/utils"
	"log"
	"net/http"
	"time"
)

func Send(req *mdl.Request) (succeed bool, errCode int, reqTime uint64) {
	client := &http.Client{
		Timeout: req.Timeout,
	}

	// built basic request
	request, err := http.NewRequest(req.Method, req.Url, req.GetBody())
	if err != nil {
		return
	}

	// default as utf-8
	if _, ok := req.Headers["Content-Type"]; !ok {
		if req.Headers == nil {
			req.Headers = make(map[string]string)
		}
		req.Headers["Content-Type"] = "application/x-www-form-urlencoded; charset=utf-8"
	}

	// fill header
	for k, v := range req.Headers {
		request.Header.Set(k, v)
	}

	// timecost
	begin := time.Now()
	resp, err := client.Do(request)
	reqTime = uint64(utils.DiffNano(begin))
	if err != nil {
		// TODO: add target of output
		log.Fatalf("Request err: %v", err)
		errCode = mdl.RequestErr
		return
	}

	// TODO: verify resp
	errCode, succeed = mdl.VerifyHttp(req, resp)
	return
}
