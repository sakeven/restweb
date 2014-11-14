package restweb

import (
	"container/list"
	"encoding/json"
	"io"
	"net/http"
	"reflect"
	"regexp"
	"strings"
)

type Control struct {
	Type   reflect.Type
	Method string
	Rx     *regexp.Regexp
	Action string
}

var controllerList = &list.List{}

func RegisterController(controller Router) {
	ct := reflect.TypeOf(controller)
	controllerList.PushBack(ct)
}

type Controller struct {
	Data map[string]interface{}
	*Context
	Action string //method of controller being callled
	Name   string
}

func (ct *Controller) Init() {
}

func (ct *Controller) Set(ctx *Context, action, name string) {
	ct.Context = ctx
	ct.Action = action
	ct.Name = name
	ct.Data = make(map[string]interface{})
	SessionManager.StartSession(ct.Response, ct.Requset)
}
func (ct Controller) Post() {
	http.Error(ct.Response, "No such page", http.StatusNotFound)
}

func (ct Controller) Get() {
	http.Error(ct.Response, "No such page", http.StatusNotFound)
}

func (ct Controller) Put() {
	http.Error(ct.Response, "No such page", http.StatusNotFound)
}

func (ct Controller) Delete() {
	http.Error(ct.Response, "No such page", http.StatusNotFound)
}

func (ct Controller) Patch() {
	http.Error(ct.Response, "No such page", http.StatusNotFound)
}

func (ct Controller) Head() {
	http.Error(ct.Response, "No such page", http.StatusNotFound)
}

func (ct Controller) Options() {
	http.Error(ct.Response, "No such page", http.StatusNotFound)
}

func (c *Controller) RenderTemplate(tplfiles ...string) {
	t, err := ParseFiles(tplfiles...)
	if err == nil {
		err = t.Execute(c.Response, c.Data)
	}
	if err != nil {
		http.Error(c.Response, "No such page", http.StatusNotFound)
		// Logger.Debug(err)
	}
}

func (c Controller) Render() { //auto render-> views/ControllerName/ActionName.tpl
	tplpath := "views/" + strings.ToLower(c.Name) + "/" + strings.ToLower(c.Action) + ".tpl"
	c.RenderTemplate("views/layout.tpl", tplpath)
}

func (ct Controller) GetAction(path string, pos int) string {
	path = strings.Trim(path, "/")
	pathsplit := strings.Split(path, "/")
	if pos >= 0 && pos < len(pathsplit) {
		return pathsplit[pos]
	}
	return ""
}

func (ct Controller) PostReader(i interface{}) (r io.Reader, err error) {
	b, err := json.Marshal(i)
	if err != nil {
		return
	}
	r = strings.NewReader(string(b))
	return
}
