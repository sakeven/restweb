package restweb

import (
	"net/http"
	"reflect"
	"strings"
)

type Server struct {
}

//路由，先处理静态文件，后处理控件，按照最大匹配原则匹配路由
func (s *Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	path := r.URL.Path
	path = strings.TrimRight(path, "/") + "/"
	filemaxlenth := 0
	var realFileHandler http.Handler
	for pattern, fileHandler := range FileMap {
		if len(pattern) > filemaxlenth && strings.HasPrefix(path, pattern) {
			filemaxlenth = len(pattern)
			realFileHandler = fileHandler
		}
	}

	macth := false
	var realRouter *Control
	for e := routerList.Front(); e != nil; e = e.Next() {
		c := e.Value.(*Control)
		if c.Rx.MatchString(path) {
			macth = true
			realRouter = c
			break
		}
	}

	if filemaxlenth > 0 {

		realFileHandler.ServeHTTP(w, r)

	} else if macth {
		ctx := &Context{Response: w, Requset: r}

		do_filter := func(when int) bool {
			for e := filterList.Front(); e != nil; e = e.Next() {
				filter := e.Value.(*Filters)

				if filter.When != when {
					continue
				}

				if filter.Rx.MatchString(path) {
					if filter.Filter(ctx) {
						return true
					}
				}
			}
			return false
		}

		if do_filter(Before) {
			return
		}

		action := realRouter.Action
		if r.Method != realRouter.Method {
			action = strings.Title(strings.ToLower(r.Method))
		}

		value := reflect.New(realRouter.Type)
		rv := GetReflectValue(ctx, action, realRouter.Type.Name())
		rm := value.MethodByName("Set")
		rm.Call(rv)
		rm = value.MethodByName(action)
		rm.Call(nil)

		do_filter(After)
	} else {
		http.Error(w, "no such page", 404)
	}
}

// 运行服务器
func Run() error {
	if err := LoadRouter(); err != nil { //import routers
		return err
	}
	if cfg.SessOn {
		SessionManager = NewManager()
		Logger.Info("Start New Session manager")
		go SessionManager.GC()
	}
	initFuncMap()
	return http.ListenAndServe(cfg.Port, &Server{})
}
