package main

import (
	"log"
	"os"
)

func usage() {
	log.Println("[Usage]:restweb [cmd] [app]")
}

var appName string

func main() {
	if len(os.Args) < 3 {
		usage()
	}
	cmd := os.Args[1]
	appName = os.Args[2]
	switch cmd {
	case "new":
		newApp()
	case "build":
		buildApp()
	case "phase":
		phaseApp()
	case "run":
		runApp()
	}
}

func phaseApp() {

}
