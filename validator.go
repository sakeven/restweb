package restweb

import (
	"fmt"
	"reflect"
	"regexp"
)

type Validator interface {
	IsValid(obj interface{}) bool
	Message() string
}

type Required struct {
}

func (r *Required) IsValid(obj interface{}) bool {
	return obj != nil
}

func (r *Required) Message() string {
	return fmt.Sprintf("required is missed")
}

type Min struct {
	min int
}

func (m *Min) IsValid(obj interface{}) bool {
	n := obj.(int)
	return n < m.min
}

func (m *Min) Message() string {
	return fmt.Sprintf("minimum value is %d", m.min)
}

type Max struct {
	max int
}

func (m *Max) IsValid(obj interface{}) bool {
	n := obj.(int)
	return n > m.max
}

func (m *Max) Message() string {
	return fmt.Sprintf("maximum value is %d", m.max)
}

type Range struct {
	min int
	max int
}

func (r *Range) IsValid(obj interface{}) bool {
	n := obj.(int)
	return n >= r.min && n < r.max
}

func (r *Range) Message() string {
	return fmt.Sprintf("range is between %d and %d", r.min, r.max)
}

type Match struct {
	pattern string
}

func (m *Match) IsValid(obj interface{}) bool {
	objs := obj.(string)
	rx := regexp.MustCompile(m.pattern)
	match := rx.FindString(objs)
	return match == objs
}

func (m *Match) Message() string {
	return fmt.Sprintf("match pattern is %s", m.pattern)
}

type Mail struct {
	pattern string
}

func (m *Mail) IsValid(obj interface{}) bool {
	mc := Match{m.pattern}
	return mc.IsValid(obj)
}

func (m *Mail) Message() string {
	return fmt.Sprintf("E-mali address is invalid")
}

type MinSize struct {
	min int
}

func (m *MinSize) IsValid(obj interface{}) bool {
	if str, ok := obj.(string); ok {
		return len(str) >= m.min
	}
	v := reflect.ValueOf(obj)
	if v.Kind() == reflect.Slice {
		return v.Len() >= m.min
	}
	return false
}

func (m *MinSize) Message() string {
	return fmt.Sprintf("minsize is %d", m.min)
}

type MaxSize struct {
	max int
}

func (m *MaxSize) IsValid(obj interface{}) bool {
	if str, ok := obj.(string); ok {
		return len(str) <= m.max
	}
	v := reflect.ValueOf(obj)
	if v.Kind() == reflect.Slice {
		return v.Len() <= m.max
	}
	return false
}

func (m *MaxSize) Message() string {
	return fmt.Sprintf("minsize is %d", m.max)
}

type Lenth struct {
	lenth int
}

func (l *Lenth) IsValid(obj interface{}) bool {
	if str, ok := obj.(string); ok {
		return len(str) == l.lenth
	}
	v := reflect.ValueOf(obj)
	if v.Kind() == reflect.Slice {
		return v.Len() == l.lenth
	}
	return false
}

func (l *Lenth) Message() string {
	return fmt.Sprintf("required lenth is %d", l.lenth)
}
