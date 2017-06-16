package http

import (
	"io"
)

type Request struct {
	// Http方法
	Method string

	Header Header

	Body io.ReadCloser

	Host string

	RequestURI string
}