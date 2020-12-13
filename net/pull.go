package net

import (
	"go-trace/mdl"
	"sync"
)

func PressTrace(req *mdl.Request, con, total uint64) {
	var wg sync.WaitGroup

	for i := uint64(0); i < con; i++ {
		wg.Add(1)
		go ProposeHTTP(req, total, &wg)
	}
}

func ProposeHTTP(req *mdl.Request, total uint64, wg *sync.WaitGroup) {

	defer wg.Done()

	for i := uint64(0); i < total; i++ {
		//
	}
}
