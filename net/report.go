package net

import (
	"go-trace/mdl"
	"sync"
)

func ReportStat(ch <-chan *mdl.RequestResult, wg *sync.WaitGroup) {
	defer wg.Done()

	// consume req result
	for ret := range ch {

	}
}
