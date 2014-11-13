package controller

import (
	"restweb"
)

type Index struct {
	restweb.Controller
}

func (i Index) Home() {
	i.Data["Say"] = "Say"
	i.Render()
}
func (i Index) Get() {
	i.Response.Write([]byte("test"))
}
