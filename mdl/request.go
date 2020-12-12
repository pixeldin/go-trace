package mdl

import (
	"go-trace/utils"
	"io"
	"strings"
	"time"
)

// HTTP Resp Code
const (
	HTTPOK = 200
)

type Request struct {
	Url     string // Url
	Form    string // http/webSocket/tcp
	Method  string // GET/POST/PUT
	Headers map[string]string
	Body    string
	Timeout time.Duration
}

func NewRequest(param Param, timeout time.Duration) (req *Request, err error) {

	// default index
	var (
		method  = "GET"
		headers = make(map[string]string)
		body    string
	)

	if param.Body != "" {
		method = "POST"
		body = param.Body
		headers["Content-Type"] = "application/x-www-form-urlencoded; charset=utf-8"
	}

	for _, hd := range param.Header {
		utils.FillInHeaderValue(hd, headers)
	}

	// verify

	// default as 30s
	if timeout == 0 {
		timeout = 30 * time.Second
	}
	req = &Request{
		Url:     param.Url,
		Method:  method,
		Headers: headers,
		Body:    body,
		Timeout: timeout,
	}

	return
}

func (r *Request) GetBody() (bReader io.Reader) {
	bReader = strings.NewReader(r.Body)
	return
}
