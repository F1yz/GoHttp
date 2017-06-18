package httpserver

import (
	"fmt"
)

type Request struct {
	// Http方法
	Method string

	Proto string

	Header Header

	Body string // string or []bytes ?

	Host string

	RequestURI string
}

func (req *Request) String() string {
	return fmt.Sprintf("Method: %s, Proto: %s, Host: %s, requestURI: %s, headers: %q",
		req.Method, req.Proto, req.Host, req.RequestURI, req.Header)
}