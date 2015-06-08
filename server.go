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
    defer func() {
        if err := recover(); err != nil {
            w.WriteHeader(http.StatusInternalServerError)
            Logger.Debugf("Panic: %v\n", err)
            w.WriteHeader(500)
        }
    }()

    ctx := NewContext(w, r)

    do_filter := func(when int) bool {
        for e := filterList.Front(); e != nil; e = e.Next() {
            filter := e.Value.(*Filters)
            path := r.URL.Path
            path = strings.TrimRight(path, "/") + "/"

            if filter.When == when &&
                (filter.Method == r.Method || filter.Method == ANY) &&
                filter.Rx.MatchString(path) {
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
    var matchs []string
    var realRouter *Control
    for e := routerList.Front(); e != nil; e = e.Next() {
        c := e.Value.(*Control)
        if c.Rx.MatchString(path) { //TODO not match method
            macth = true
            realRouter = c
            matchs = c.Rx.FindStringSubmatch(path)[1:]
            if c.Method == r.Method {
                break
            }
        }
    }
    if filemaxlenth > 0 {

        realFileHandler.ServeHTTP(w, r)

    } else if macth {

        action := realRouter.Action
        if r.Method != realRouter.Method {
            action = strings.Title(strings.ToLower(r.Method))
        }

        value := reflect.New(realRouter.Type)
        rv := GetReflectValue(ctx, action, realRouter.Type.Name())
        rm := value.MethodByName("Set")
        rm.Call(rv)
        rm = value.MethodByName("Init")
        rm.Call(nil)
        rm = value.MethodByName(action)

        if do_filter(Middle) {
            return
        }
        rv = make([]reflect.Value, 0)
        for _, j := range matchs {
            rw := reflect.ValueOf(j)
            rv = append(rv, rw)
        }
        rm.Call(rv)

    } else {
        http.Error(w, "no such page", 404)
    }
    do_filter(After)
}

// 运行服务器
func Run() error {
    LoadRouter("config/default_router.conf")
    if err := LoadRouter("config/router.conf"); err != nil { //import routers
        return err
    }
    Logger.Info("Start Server at " + cfg.Port + " Port. Please visit http://localhost" + cfg.Port)
    if cfg.SessOn {
        SessionManager = NewManager()
        Logger.Info("Start New Session manager.")
        go SessionManager.GC()
    }
    return http.ListenAndServe(cfg.Port, &Server{})
}
