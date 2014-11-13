package restweb

import (
	"container/list"
)

const (
	Before = iota
	After
)

type Filter func(ctx *Context) bool
type Filters struct {
	Filter  Filter
	Pattern string
	When    int
}

var FilterList = &list.List{}

func RegisterFilters(pattern string, when int, filter Filter) {
	FilterList.PushBack(&Filters{Filter: filter, Pattern: pattern, When: when})
}
