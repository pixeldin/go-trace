package net

import (
	"go-trace/mdl"
	"log"
	"os"
	"testing"
)

var req = &mdl.Request{
	Url:     "https://www.baidu.com/",
	Method:  "GET",
	Timeout: 0,
}

var testLog = log.New(os.Stdout, "ClientTest# ", 0)
var testErrLog = log.New(os.Stderr, "ClientTest Err# ", 0)

func TestSend(t *testing.T) {
	suc, code, time := Send(req)
	testLog.Printf("Send succeed status: %v, code: %d", suc, code)
	log.Printf("Req url: %v, timeCost: %v(ts)", req.Url, time/1e6)
	//err := errors.New("Something bad")
	//testErrLog.Printf("err : %v", err)
}
