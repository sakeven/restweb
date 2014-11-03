package main

import (
	"restweb"
	// "net/http"
)

type Index struct {
	restweb.Controller
}

// func (i Index) Get(w http.ResponseWriter, r *http.Request) {
// 	w.WriteHeader(200)
// 	w.Write([]byte("Hello world"))
// }

func main() {
	restweb.AddRouter("/", Index{})
	restweb.Run()
}
