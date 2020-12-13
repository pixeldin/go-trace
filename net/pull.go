package net

import (
	"go-trace/mdl"
	"sync"
)

func PressTrace(req *mdl.Request, con, total uint64) {
	resultChannel := make(chan *mdl.RequestResult, 1000)

	var wg sync.WaitGroup
	for i := uint64(0); i < con; i++ {
		wg.Add(1)
		go ProposeHTTP(resultChannel, req, total, &wg)
	}
}

func ProposeHTTP(resultChan chan<- *mdl.RequestResult, req *mdl.Request, total uint64, wg *sync.WaitGroup) {

	defer wg.Done()

	for i := uint64(0); i < total; i++ {
		succeed, code, time := Send(req)

		// build result
		reqResult := &mdl.RequestResult{
			Time:    time,
			Succeed: succeed,
			ErrCode: code,
		}

		resultChan <- reqResult
	}

	return
}
