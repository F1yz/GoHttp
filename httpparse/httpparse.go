package httpparse

import (
	"strings"
	"fmt"
)

type HttpParse struct {
}

func (httpParse *HttpParse) ParseRequest(requestBytes []byte) map[string]string {
	requestStr := string(requestBytes)
	requestArr := strings.Split(requestStr, "\r\n")
	requestRow := make([]string, 5)
	parseData := map[string]string{}

	for key, val := range requestArr {
		if key == 0 {
			requestRow = strings.Split(val, " ")
			parseData["method"] = string(requestRow[0])
			continue
		}

		if val == "" {
			parseData["body"] = strings.Join(requestArr[key + 1:], "\r\n")
			break
		}

		requestRow = strings.Split(val, ":")
		parseData[requestRow[0]] = requestRow[1]
	}

	fmt.Println(parseData)
	return parseData
}
