package main

import (
	"testing"
)

func Test_phaseApp(t *testing.T) {
	err := phaseApp("@URL:/test @method:GET")
	if err != nil {
		t.Error(err)
	}
	if adec.Method != "GET" || adec.URL != "/test/" {
		t.Error("Decorators value error")
	}
}
