package restweb

import (
	"container/list"
	"regexp"
)

const (
	Before = iota
	Middle
	After
)

const (
	POST    = "POST"
	GET     = "GET"
	PUT     = "PUT"
	DELETE  = "DELETET"
	ANY     = "ANY"
	PATCH   = "PATCH"
	HEAD    = "HEAD"
	OPTIONS = "OPTIONS"
)

type Filter func(ctx *Context) bool
type Filters struct {
	Filter Filter
	Rx     *regexp.Regexp
	When   int
	Method string
}

var filterList = &list.List{}

func RegisterFilters(method string, pattern string, when int, filter Filter) {
	rx := regexp.MustCompile(pattern)
	filterList.PushBack(&Filters{Filter: filter, Rx: rx, When: when, Method: method})
}
