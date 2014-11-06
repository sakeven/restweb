package controller

import (
	"net/http"
	"restweb"
)

type Index struct {
	restweb.Controller
}

func (i Index) Get(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(200)
	w.Write([]byte("Hello world"))
}
