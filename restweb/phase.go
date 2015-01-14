package main

import (
	"errors"
	"strings"
)

func isSpace(char rune) bool {
	if char == ' ' || char == '\t' {
		return true
	}
	return false
}

func isCR(char rune) bool {
	if char == '\n' || char == '\r' {
		return true
	}
	return false
}

func isAt(char rune) bool {
	if char == '@' {
		return true
	}
	return false
}

func isColon(char rune) bool {
	if char == ':' {
		return true
	}
	return false
}

type dec struct {
	URL    string
	Method string
	Test   bool
}

func (d *dec) Clear() {
	d.Method = ""
	d.URL = ""
	d.Test = false
}

var adec = dec{}

func phaseApp(decorator string) error {
	isKey, isValue := false, false
	key, value := "", ""
	pair := make(map[string]string)
	decorator += "@"

	for _, char := range decorator {
		switch {
		case isSpace(char) || isCR(char):
			continue
		case isAt(char):
			isValue = false
			isKey = true
			if key != "" {
				pair[key] = value
			}
			key = ""
		case isColon(char):
			isKey = false
			isValue = true
			value = ""
		default:
			if isKey {
				key += string(char)
			} else if isValue {
				value += string(char)
			}
		}

	}

	for k, v := range pair {
		switch k {
		case "URL":
			adec.URL = v
		case "method":
			adec.Method = v
		case "test":
			adec.Test = true
		}
	}

	if adec.Method == "" || adec.URL == "" {
		return errors.New("Decorators value miss")
	}

	adec.URL = strings.TrimRight(adec.URL, "/") + "/"
	return nil
}
