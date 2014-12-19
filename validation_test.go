package restweb

import (
	"testing"
)

var va = &Validation{}

//some error occur
func Test_Required(t *testing.T) {
	va.Clear()
	va.Required("", "string")
	va.Required(0, "int")
	va.Required(nil, "ptr")
	verr := va.RenderErrMap()
	for key, err := range verr {
		if err != "required is missed" {
			t.Error("requried error of key ", key)
		}
	}
	va.Clear()
	va.Required("233", "string")
	va.Required(233, "int")
	va.Required(va, "ptr")
	verr = va.RenderErrMap()
	for key, err := range verr {
		if err == "required is missed" {
			t.Error("not requried error of key ", key)
		}
	}
}

func Test_Min(t *testing.T) {
	va.Clear()
	va.Min(10, 5, "not min")
	if !va.HasError {
		t.Error("not min error")
	}
	va.Clear()
	va.Min(5, 10, "min")
	if va.HasError {
		t.Error("min error")
	}
	va.Clear()
}

func Test_Max(t *testing.T) {
	va.Clear()
	va.Max(1, 5, "min")
	if !va.HasError {
		t.Error("not max error")
	}
	va.Clear()
	va.Max(5, 1, "not min")
	if va.HasError {
		t.Error("max error")
	}
	va.Clear()
}

func Test_Range(t *testing.T) {
	va.Clear()
	va.Range(5, 1, 10, "range_ok")
	if va.HasError {
		t.Error("range_ok error")
	}
	va.Clear()
	va.Range(5, 6, 10, "range_min")
	if !va.HasError {
		t.Error("range_min error")
	}
	va.Clear()
	va.Range(11, 1, 10, "range_max")
	if !va.HasError {
		t.Error("range_max error")
	}
	va.Clear()
}

func Test_Mail(t *testing.T) {
	va.Clear()
	va.Mail("123.esdf@123example.edu.cn", "mail_ok")
	if va.HasError {
		t.Error("mail_ok error")
	}
	va.Clear()
	va.Mail("123@", "mail_error")
	if !va.HasError {
		t.Error("mail_error erro")
	}
	va.Clear()
}

func Test_Match(t *testing.T) {
	va.Clear()
	va.Match("wqe213", `[A-Za-z]*[0-9]*`, "match_ok")
	if va.HasError {
		t.Error("match_ok error")
	}
	va.Clear()
	va.Match("qwe", `[0-9]`, "match_error")
	if !va.HasError {
		t.Error("match_error error")
	}
}
func Test_Equal(t *testing.T) {
	va.Clear()
	va.Equal("123", "123", "pwd")
	if va.HasError {
		t.Error("Equal error")
	}

	va.Clear()
	va.Equal("123", "321", "pwd")
	if !va.HasError {
		t.Error("Equal error")
	}
}
