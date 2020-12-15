package net

import (
	"fmt"
	"go-trace/mdl"
	"testing"
)

func TestPrintHeader(t *testing.T) {
	PrintHeader()

	stat := new(mdl.StaticMetric)
	stat.ByNow = 12333
	stat.Concurrency = 1
	stat.SucCount = 1
	qps := float64(stat.SucCount*1e9) / float64(stat.ByNow)
	result := fmt.Sprintf("%8.2fms│%11d│%12d│%9d|%9.2f\n",
		float64(stat.ByNow)/1e6, stat.Concurrency, stat.SucCount, stat.FailCount, qps)
	fmt.Println(result)
}
