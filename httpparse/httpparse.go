package httpparse

import (
	"strings"
	"fmt"
	"httpserver"
)

type HttpParse struct {
}

func (httpParse *HttpParse) ParseRequest(requestBytes []byte) *httpserver.Request {
	requestStr := string(requestBytes)
	requestArr := strings.Split(requestStr, "\r\n")
	requestRow := make([]string, 5)

	var req = new(httpserver.Request)

	requestLine := requestArr[0]


	req.Method, req.RequestURI, req.Proto = parseRequestLine(requestLine)
	req.Header = httpserver.Header{}

	headersAndBody := requestArr[1:]

	for key, val := range headersAndBody {
		if val == "" {
			req.Body = strings.Join(requestArr[key + 1:], "\r\n")
			break
		}

		// parse headers, header value not trimmed yet.
		requestRow = strings.Split(val, ":")
		req.Header.Add(requestRow[0], requestRow[1])
	}

	fmt.Println(req)
	return req
}

func parseRequestLine(line string) (method, requestURI, proto string)  {
	result := strings.Split(line, " ")

	if 3 != len(result) {
		return
	}

	return result[0], result[1], result[2]
}

func parseHeader(headerStr string) (key, val string)  {
	//headerInfo := strings.Split(headerStr, ":")
	return
}
