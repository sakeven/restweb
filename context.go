package restweb

import (
	"net/http"
	"net/url"
)

type Context struct {
	Requset  *http.Request
	Response http.ResponseWriter
	Input    url.Values
	Output   map[string]interface{}
}

func NewContext(w http.ResponseWriter, r *http.Request) (c *Context) {
	c = &Context{Requset: r, Response: w}
	c.Requset.ParseForm()
	c.Input = c.Requset.Form
	c.Output = make(map[string]interface{})
	return
}

func (c *Context) SetSession(key string, value string) {
	session := SessionManager.StartSession(c.Response, c.Requset)
	session.Set(key, value)
}

func (c *Context) GetSession(key string) string {
	sess := SessionManager.StartSession(c.Response, c.Requset)
	return sess.Get(key)
}

func (c *Context) DeleteSession() {
	SessionManager.DeleteSession(c.Response, c.Requset)
}

func (c *Context) Redirect(urlStr string, code int) {
	http.Redirect(c.Response, c.Requset, urlStr, code)
}

func (c *Context) Error(err string, code int) {
	http.Error(c.Response, err, code)
}
