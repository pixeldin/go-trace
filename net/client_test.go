package net

import (
	"context"
	"fmt"
	"go-trace/mdl"
	"go-trace/utils"
	"log"
	"os"
	"sync"
	"testing"
	"time"
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

func TestSome(t *testing.T) {
	var wg sync.WaitGroup
	wg.Add(1)
	fmt.Println("Just wait")
	wg.Wait()
	// never run here
	return
}

func TestTicking(t *testing.T) {
	fmt.Printf("%s Mock start.\n", utils.GetPresentFormat())
	ctx, cancelFunc := context.WithCancel(context.Background())
	go func(ctx context.Context) {
		for {
			time.Sleep(3 * time.Second)
			select {
			case <-ctx.Done():
				fmt.Printf("%s Ctx done!\n", utils.GetPresentFormat())
				return
			default:
				go SomethingCostTime()
			}
		}
	}(ctx)

	// wait
	time.Sleep(20 * time.Second)
	cancelFunc()
	fmt.Printf("%s Mock end.\n", utils.GetPresentFormat())
	// wait for ctx done.
	time.Sleep(5 * time.Second)
}

func SomethingCostTime() {
	fmt.Printf("%s Cost start.\n", utils.GetPresentFormat())
	time.Sleep(5 * time.Second)
	fmt.Printf("%s Cost end.\n", utils.GetPresentFormat())
}
