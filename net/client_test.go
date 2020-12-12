package net

import (
	"go-trace/mdl"
	"log"
	"testing"
)

var req = &mdl.Request{
	Url:     "https://www.baidu.com/",
	Method:  "GET",
	Timeout: 0,
}

func TestSend(t *testing.T) {
	_, time, err := Send(req)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Req url: %v, timeCost: %v(ts)", req.Url, time/1e6)
}
