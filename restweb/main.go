package main

import (
	"fmt"
	"os"
)

func usage() {
	fmt.Println("Usage: restweb [cmd] [app]")
}

var appName string

func main() {
	if len(os.Args) < 3 {
		usage()
		return
	}
	cmd := os.Args[1]
	appName = os.Args[2]
	switch cmd {
	case "new":
		newApp()
	case "build":
		buildApp()
	case "run":
		runApp()
	case "clean":
		cleanApp()
	default:
		usage()
	}
}

func cleanApp() {
	os.Remove(appName + "/" + appName)
	os.Remove(appName + "/main.go")
	os.Remove(appName + "/config/router.conf")
}
