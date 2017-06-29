package httpserver

import (
	"fmt"
	"strings"
)

type Response struct {
	StatusCode int    // 200
	
	Proto string

	Header Header

	Body []byte

	ContentLength int64
}

func (resp *Response) SetHeaders(header Header)  {
	resp.Header = header
}

func (resp *Response) SetBody(body []byte)  {
	resp.Body = body
}


func (resp *Response) WriteHeaders() string {
	size := len(resp.Header)
	headerStr := make([]string, size)

	i := 0
	for key, values := range resp.Header {

		valStr := strings.Join(values, ",")

		headerStr[i] = key + ":" + valStr
		i++
	}

	return strings.Join(headerStr, "\r\n")
}


func (resp *Response) WriteStatusLine() string  {
	return fmt.Sprintf("%s %d %s", resp.Proto, resp.StatusCode, StatusText(resp.StatusCode))
}
