package net

import (
	"context"
	"go-trace/mdl"
	"sync"
	"time"
)

var StatTicking = 1 * time.Second

func ReportStat(ch <-chan *mdl.RequestResult, wg *sync.WaitGroup) {
	defer wg.Done()

	// for cancel notify
	ctx, cancel := context.WithCancel(context.Background())

	var (
		byNow       uint64
		reqCostTime uint64
		sucCount    uint64
		failCount   uint64
	)

	startTime := uint64(time.Now().Nanosecond())

	// print head

	// calculate with ticking
	go func(ctx context.Context) {
		for {
			time.Sleep(StatTicking)
			select {
			case <-ctx.Done():
				// report ctx canceled
				return
			default:
				byNow = uint64(time.Now().Nanosecond()) - startTime
				// asynchronous print
				go aggregate(reqCostTime, byNow, sucCount, failCount)
			}
		}
	}(ctx)

	// consume request result and summary
	for ret := range ch {
		reqCostTime += ret.Time
		if ret.Succeed {
			sucCount++
		} else {
			failCount++
		}
	}

	cancel()
}

func aggregate(duration, totalTime, sucCount, failCount uint64) {
	// print table
}
