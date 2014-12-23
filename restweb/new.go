package main

import (
	// "log"
	"os"
)

func newApp() {
	os.Mkdir(appName, 0777)
	os.Mkdir(appName+"/controller", 0777)
	os.Mkdir(appName+"/views", 0777)
	os.Mkdir(appName+"/config", 0777)
	os.Mkdir(appName+"/static", 0777)
}
