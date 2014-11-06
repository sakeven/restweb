package config

import (
	"testing"
)

func Test_config(t *testing.T) {
	cfg := &Config{}
	cfg.ReadConfig("config.json")
	if cfg.SessOn != true || cfg.Port != ":8080" {
		t.Error("test failed")
	}
}
