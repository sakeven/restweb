package config

import (
	"bufio"
	"encoding/json"
	"log"
	"os"
	"reflect"
	"strings"
)

type Config struct {
	Port    string `json:"port"`
	SessOn  bool   `json:"sesson"`
	Log     string `json:"log"`
	PkgPath string `json:pkgpath`
}

func NewConfig() *Config {
	return &Config{}
}

func (cfg *Config) ReadConfig(configFile string) {
	file, err := os.Open(configFile)
	if err != nil {
		log.Fatal(err)
	}

	br := bufio.NewReader(file)
	err = json.NewDecoder(br).Decode(cfg)
	if err != nil {
		log.Fatal(err)
	}
}

func (cfg *Config) Get(Key string) interface{} {
	rv := reflect.ValueOf(*cfg)
	return rv.FieldByName(strings.Title(Key))
}
