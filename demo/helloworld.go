package main

import (
	"log"
	"restweb"
	"restweb/demo/controller"
)

func main() {
	restweb.RegisterController(controller.Index{})
	restweb.RegisterFilters("/haha", restweb.Before, filter)
	log.Fatal(restweb.Run())
}

func filter(ctx *restweb.Context) bool {
	ctx.Redirct("/")
	return true
}
