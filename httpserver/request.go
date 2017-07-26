package httpserver

import (
	"fmt"
	"strings"
	"strconv"
	"errors"
)

type Request struct {
	// Http方法
	Method string

	Proto string

	Header Header

	Body []byte // string or []bytes ?

	Host string

	RequestURI string
}

func NewRequest(requestData map[string]interface{}) (request *Request) {
	request = &Request{}
	request.Method = requestData["Method"].(string)
	request.Proto = requestData["Proto"].(string)
	request.RequestURI = requestData["RequestURI"].(string)

	body, err := requestData["Body"].([]byte)
	if err {
		request.Body = body
	}

	host, err := requestData["Host"].(string)
	if err {
		request.Host = host
	}

	delete(requestData, "Method")
	delete(requestData, "Proto")
	delete(requestData, "Body")
	delete(requestData, "Host")
	delete(requestData, "RequestURI")

	request.Header = Header{}
	for key, val := range requestData {
		request.Header.Add(key, val.(string))
	}

	return
}

func (req *Request) IsRangeRequest() bool {
	if rangeHeader := req.Header.Get("Range"); rangeHeader != "" {
		return true
	}

	return false
}

func (req *Request) GetStartEndRange() ([]int, error)  {
	rangeHeader := req.Header.Get("Range")

	fmt.Println(fmt.Sprintf("------------------%v", rangeHeader))
	if !strings.HasPrefix(rangeHeader, "bytes=") {
		return nil, errors.New("Invalid range")
	}

	directives := strings.Split(strings.Trim(rangeHeader, "bytes="), " ")

	startEndRange := directives[1]
	rangeInfo := strings.TrimRight(startEndRange, ",")
	rangeArray := strings.Split(rangeInfo, "-")

	if len(rangeArray) == 1 {
		return []int{ 999, 0,}, errors.New("Invalid range")
	}

	length := 2
	if rangeArray[1] == "" {
		length = 1
	}

	rangeVal := make([]int, length)

	for i := 0; i < len(rangeVal); i++ {
		rangeIntVal, _ := strconv.Atoi(rangeArray[i])
		rangeVal[i] = rangeIntVal
	}

	return rangeVal, nil
}

func (req *Request) String() string {
	return fmt.Sprintf("Method: %s, Proto: %s, Host: %s, requestURI: %s, headers: %q",
		req.Method, req.Proto, req.Host, req.RequestURI, req.Header)
}