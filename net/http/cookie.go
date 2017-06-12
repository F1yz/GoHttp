package http

import (
	"time"
)

type Cookie struct {
	Name  string
	Value string

	Path       string    // optional
	Domain     string    // optional
	Expires    time.Time // optional

	HttpOnly bool
}

// 4 debug?
func (c *Cookie) String() string {
	
}