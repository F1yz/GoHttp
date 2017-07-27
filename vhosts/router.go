package vhosts

import "fmt"

type Router struct {
	Pattern string `yaml:"pattern"`
	Rules []Rule `yaml:"rules"`
}

func (router *Router) String() string {
	return fmt.Sprintf("Pattern: %s, Rules %v", router.Pattern, router.Rules)
}

type Rule struct {
	ETag string
	Expires int
}