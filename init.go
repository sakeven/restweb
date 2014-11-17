package restweb

import (
	"golog"
	"os"
	"restweb/config"
)

var SessionManager *Manager
var Logger *golog.Log
var cfg *config.Config

func init() {
	cfg = new(config.Config)
	cfg.ReadConfig("config/app.conf")
	Logger = golog.NewLog(os.Stdout, golog.Ldebug|golog.Linfo)
	initFuncMap()
}
