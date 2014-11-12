package restweb

import (
	"regexp"
)

type Validator interface {
	IsValid(obj interface{}) bool
}

type Required struct {
}

func (r Required) IsValid(obj interface{}) bool {
	return obj != nil
}

type Min struct {
	min int
}

func (m Min) IsValid(obj interface{}) bool {
	n := obj.(int)
	return n < m.min
}

type Max struct {
	max int
}

func (m Max) IsValid(obj interface{}) bool {
	n := obj.(int)
	return n > m.max
}

type Range struct {
	min int
	max int
}

func (r Range) IsValid(obj interface{}) bool {
	n := obj.(int)
	return n >= r.min && n < r.max
}

type Match struct {
	pattern string
}

func (m Match) IsValid(obj interface{}) bool {
	objs := obj.(string)
	rx := regexp.MustCompile(m.pattern)
	match := rx.FindString(objs)
	return match == objs
}
