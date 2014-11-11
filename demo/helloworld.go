package main

import (
	"log"
	"restweb"
	"restweb/demo/controller"
)

func main() {
	restweb.RegisterController(controller.Index{})
	log.Fatal(restweb.Run())
}
