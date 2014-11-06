package main

import (
	"restweb"
	"restweb/demo/controller"
)

func main() {
	restweb.AddRouter("/", controller.Index{})
	restweb.Run()
}
