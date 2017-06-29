package httpserver

import (
	"fmt"
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

func (req *Request) String() string {
	return fmt.Sprintf("Method: %s, Proto: %s, Host: %s, requestURI: %s, headers: %q",
		req.Method, req.Proto, req.Host, req.RequestURI, req.Header)
}