package net

import (
	"go-trace/mdl"
	"sync"
	"time"
)

func PressTrace(req *mdl.Request, con, total uint64) {

	// new routine for calculate
	var wgReport sync.WaitGroup
	wgReport.Add(1)
	resultChannel := make(chan *mdl.RequestResult, 1000)
	go ReportStat(con, resultChannel, &wgReport)

	var wg sync.WaitGroup
	for i := uint64(0); i < con; i++ {
		wg.Add(1)
		go ProposeHTTP(resultChannel, req, total, &wg)
	}

	// blocking
	wg.Wait()
	time.Sleep(1 * time.Millisecond)
	close(resultChannel)
	wgReport.Wait()
	return
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
