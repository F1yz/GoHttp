package http

import (
	"io"
	"fmt"
)

type Request struct {
	// Http方法
	Method string

	Proto string

	Header Header

	Body io.ReadCloser

	Host string

	RequestURI string
}

func (req *Request) String() string {
	return fmt.Sprintf("Method: %s, Proto: %s, Host: %s, requestURI: %s",
		req.Method, req.Proto, req.Host, req.RequestURI)
}