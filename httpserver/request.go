package httpserver

type Request struct {
	header map[string]string
	reqParams map[string]string
	fileData []map[string]string
}