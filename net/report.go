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

	startTime := uint64(time.Now().UnixNano())

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
				stm.ByNow = uint64(time.Now().UnixNano()) - startTime
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
	stm.ByNow = uint64(time.Now().UnixNano()) - startTime
	aggregate(stm)
}

func aggregate(stat *mdl.StaticMetric) {
	// print table
	//result := fmt.Sprintf("costTime│concurrency│successCount│failCount│QPS")
	qps := float64(stat.SucCount*1e9) / float64(stat.ByNow)
	result := fmt.Sprintf("%8.2fms│%11d│%12d│%9d|%9.2f",
		float64(stat.ByNow)/1e6, stat.Concurrency, stat.SucCount, stat.FailCount, qps)
	fmt.Println(result)
	return
}

func PrintHeader() {
	// TODO: switch output way
	format := `──────────┬───────────┬────────────┬─────────┬─────────
 costTime │concurrency│successCount│failCount│   QPS   
──────────┼───────────┼────────────┼─────────┼─────────`
	fmt.Println(format)
	return

	//fmt.Println("\n──────────┬───────────┬────────────┬─────────┬─────────")
	//head := fmt.Sprintf(" costTime │concurrency│successCount│failCount│   QPS   \n" +
	//	"──────────┼───────────┼────────────┼─────────┼─────────")
	//fmt.Println(head)
	//fmt.Println("──────────┼───────────┼────────────┼─────────┼─────────")
}
