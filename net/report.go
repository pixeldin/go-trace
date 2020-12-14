package net

import (
	"context"
	"fmt"
	"go-trace/mdl"
	"sync"
	"time"
)

var StatTicking = 1 * time.Second

func ReportStat(concurrent uint64, ch <-chan *mdl.RequestResult, wg *sync.WaitGroup) {
	defer wg.Done()

	// for cancel notify
	ctx, cancel := context.WithCancel(context.Background())

	var stm = new(mdl.StaticMetric)
	stm.Concurrency = concurrent
	//var (
	//	byNow       uint64
	//	reqCostTime uint64
	//	sucCount    uint64
	//	failCount   uint64
	//	qps float64
	//)

	startTime := uint64(time.Now().Nanosecond())

	// print head
	PrintHeader()

	// calculate with ticking
	go func(ctx context.Context) {
		for {
			time.Sleep(StatTicking)
			select {
			case <-ctx.Done():
				// report ctx canceled
				return
			default:
				stm.ByNow = uint64(time.Now().Nanosecond()) - startTime
				// asynchronous print
				go aggregate(stm)
			}
		}
	}(ctx)

	// consume request result and summary
	for ret := range ch {
		stm.ReqCostTime += ret.Time
		if ret.Succeed {
			stm.SucCount++
		} else {
			stm.FailCount++
		}
	}

	cancel()
}

//func aggregate(duration, totalTime, concurrency, sucCount, failCount uint64, qps float64) {
func aggregate(stat *mdl.StaticMetric) {
	// print table
	//result := fmt.Sprintf("costTime│concurrency│successCount│failCount│QPS")
	qps := float64(stat.SucCount*1e9) / float64(stat.ByNow)
	result := fmt.Sprintf("%7.0fs│%11d│%12d│%9d|%9.2f\n",
		float64(stat.ByNow) / 1e9, stat.Concurrency, stat.SucCount, stat.FailCount, qps)
	fmt.Print(result)
}

func PrintHeader() {
	// TODO: switch output way
	fmt.Println("\n────────┬───────────┬────────────┬─────────┬─────────")
	head := fmt.Sprintf("costTime│concurrency│successCount│failCount│   QPS   ")
	fmt.Println(head)
	fmt.Println("────────┼───────────┼────────────┼─────────┼─────────")
}
