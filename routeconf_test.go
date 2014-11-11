package restweb

import (
	"testing"
)

func Test_TrimArgs(t *testing.T) {
	args := []string{"  123  ", " 456 "}
	TrimArgs(args)
	if args[0] != "123" || args[1] != "456" {
		t.Error("test function TrimArgs failed")
	}
}
