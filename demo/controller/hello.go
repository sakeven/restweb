package controller

import (
	"net/http"
	"restweb"
)

type Index struct {
	restweb.Controller
}

func (i Index) Home(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("hello world"))
}
func (i Index) Get(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("haha world"))
}
