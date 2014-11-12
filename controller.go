package restweb

import (
	"container/list"
	"encoding/json"
	"io"
	"net/http"
	"reflect"
	"strings"
)

type Control struct {
	Type    reflect.Type
	Method  string
	Pattern string
	Action  string
}

var controllerList = &list.List{}

func RegisterController(controller Router) {
	ct := reflect.TypeOf(controller)
	controllerList.PushBack(ct)
}

type Controller struct {
	Data   map[string]interface{}
	W      http.ResponseWriter
	R      *http.Request
	Action string //method of controller being callled
	Name   string
}

func (ct *Controller) Init() {
}

func (ct *Controller) Set(w http.ResponseWriter, r *http.Request, action, name string) {
	ct.W = w
	ct.R = r
	ct.Action = action
	ct.Name = name
	ct.Data = make(map[string]interface{})
}
func (ct Controller) Post() {
	http.Error(ct.W, "No such page", http.StatusNotFound)
}

func (ct Controller) Get() {
	http.Error(ct.W, "No such page", http.StatusNotFound)
}

func (ct Controller) Put() {
	http.Error(ct.W, "No such page", http.StatusNotFound)
}

func (ct Controller) Delete() {
	http.Error(ct.W, "No such page", http.StatusNotFound)
}

func (ct Controller) Patch() {
	http.Error(ct.W, "No such page", http.StatusNotFound)
}

func (ct Controller) Head() {
	http.Error(ct.W, "No such page", http.StatusNotFound)
}

func (ct Controller) Options() {
	http.Error(ct.W, "No such page", http.StatusNotFound)
}

func (ct *Controller) SetSession(key string, value string) {
	session := SessionManager.StartSession(ct.W, ct.R)
	session.Set(key, value)
}

func (ct *Controller) GetSession(key string) (value string) {
	session := SessionManager.StartSession(ct.W, ct.R)
	value = session.Get(key)
	return
}

func (ct *Controller) DeleteSession() {
	SessionManager.DeleteSession(ct.W, ct.R)
}

func (c *Controller) RenderTemplate(tplfiles ...string) {
	t, err := ParseFiles(tplfiles...)
	if err == nil {
		err = t.Execute(c.W, c.Data)
	}
	if err != nil {
		http.Error(c.W, "No such page", http.StatusNotFound)
		Logger.Debug(err)
	}
}

func (c Controller) Render() { //auto render
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

func (c Controller) Redirect(url string) {
	http.Redirect(c.W, c.R, url, http.StatusFound)
}
