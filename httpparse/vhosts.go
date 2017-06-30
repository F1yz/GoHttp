package httpparse

type VHostItem struct {
	web_root string
	host string
	router interface{}
}