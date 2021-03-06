package httpserver

// header的值可能有多个，如Accept-Encoding: gzip, deflate, sdch, br
type Header map[string][]string

func (h Header) Add(key string, value string)  {
	h[key] = append(h[key], value)
}

func (h Header) Get(key string) string  {
	if v := h[key]; len(v) > 0 {
		return v[0]
	}

	return ""
}

func (h Header) GetAll(key string) []string {
	if v := h[key]; len(v) > 0 {
		return v
	}

	return []string{}
}

func (h Header) Set(key string, value string)  {
	h[key] = []string{value}
}

func (h Header) Del(key string)  {
	delete(h, key)
}