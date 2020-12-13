package main

import (
	"flag"
	"fmt"
	"go-trace/mdl"
)

var param mdl.Param

func init() {
	flag.Uint64Var(&param.Concurrency, "c", 1, "Concurrency")
	flag.Uint64Var(&param.TotalQuest, "n", 1, "totalRequest")
	flag.StringVar(&param.Url, "u", "", "requestURL")
	flag.Var(&param.Header, "H", "'Content-Type: application/json'")
	flag.StringVar(&param.Body, "data", "", "data to POST")

	flag.Parse()
}

func main() {

	// param checking
	if param.Concurrency == 0 || param.TotalQuest == 0 || param.Url == "" {
		fmt.Printf("example: go run main.go -c 1 -n 1 -u https://www.google.com/ \n")
		fmt.Printf("recept param: -c %d -n %d -u %s \n", param.Concurrency, param.TotalQuest, param.Url)
		flag.Usage()
		return
	}

	// build request
	request, err := mdl.NewRequest(param, 0)
	if err != nil {
		fmt.Printf("Build request err: %v\n", err)
		return
	}

	// 处理请求

}
