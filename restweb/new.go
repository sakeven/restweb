package main

import (
	// "log"
	"os"
)

const AppConf = `
{
	"port":":8080",
	"sesson":true,
	"log":"Dev"
}
`

func newApp() {
	os.Mkdir(appName, 0777)
	os.Mkdir(appName+"/controller", 0777)
	os.Mkdir(appName+"/views", 0777)
	os.Mkdir(appName+"/config", 0777)
	appConf, err := os.Create(appName + "/config/app.conf")
	if err != nil {
		appConf.Write([]byte(AppConf))
	}
	os.Mkdir(appName+"/static", 0777)
}
