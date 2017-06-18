package http


import (
	"strings"
	"bufio"
	"net/textproto"
	"io"
	"fmt"
)

type HttpParser struct {
}

func (httpParser *HttpParser) ParseRequest(requestBytes []byte) map[string]string {
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
			break
		}

		requestRow = strings.Split(val, ":")
		parseData[requestRow[0]] = requestRow[1]
	}

	return parseData
}

// struct version ParseRequest
func ReadRequest(br *bufio.Reader) (req *Request)  {
	tpReader := textproto.NewReader(br)
	req = new(Request)

	var s string
	var err error

	if s, err = tpReader.ReadLine(); err != nil {
		return req
	}

	var ok bool

	req.Method, req.RequestURI, req.Proto, ok = parseMethodAndRequestURI(s)

	//
	if !ok {
		return nil
	}

	req.Header = parseHeaders(tpReader)

	for key, value := range req.Header  {
		fmt.Printf("Key: %q, value: %q\n", key, value)
	}

	return req
}

func parseMethodAndRequestURI(line string) (method, requestURI, proto string, ok bool) {
	result := strings.SplitN(line, " ", 3)

	fmt.Printf("????%q\n", result)
	if 3 < len(result) {
		return
	}

	return result[0], result[1], result[2], true
}

func parseHeaders(tpR *textproto.Reader) Header {
	var headerString string
	var err error

	var header = Header{}

	for {
		if headerString, err = tpR.ReadLine(); err != nil {
			if err == io.EOF {
				break
			}

			fmt.Printf("Damn, we got an error: %q\n", err)
			return nil
		}

		fmt.Printf("what we got: %q\n", headerString)

		if headerString == "" {
			return header
			//fmt.Println("We got all headers")
		} else {
			keyValue := strings.Split(headerString, ":")
			header.Add(keyValue[0], keyValue[1])
		}

	}
	return nil


}
