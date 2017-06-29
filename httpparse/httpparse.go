package httpparse

import (
	"strings"
)

type HttpParse struct {
}

func (httpParse *HttpParse) ParseRequest(requestBytes []byte) map[string]interface{} {
	requestStr := string(requestBytes)
	requestArr := strings.Split(requestStr, "\r\n")
	requestRow := make([]string, 5)

	requestLine := requestArr[0]
	requestData := make(map[string]interface{})
	requestData["Method"], requestData["RequestURI"], requestData["Proto"] = parseRequestLine(requestLine)

	for key := 1; key < len(requestArr); key ++ {
		val := requestArr[key]
		if val == "" {
			requestData["Body"] = strings.Join(requestArr[key + 1:], "\r\n")
			break
		}

		// parse headers, header value not trimmed yet.
		requestRow = strings.Split(val, ":")
		requestData[requestRow[0]] = requestRow[1]
	}

	return requestData
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
