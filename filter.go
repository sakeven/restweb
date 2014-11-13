package restweb

import (
	"container/list"
	"regexp"
)

const (
	Before = iota
	After
)

type Filter func(ctx *Context) bool
type Filters struct {
	Filter Filter
	Rx     *regexp.Regexp
	When   int
}

var filterList = &list.List{}

func RegisterFilters(pattern string, when int, filter Filter) {
	rx := regexp.MustCompile(pattern)
	filterList.PushBack(&Filters{Filter: filter, Rx: rx, When: when})
}
