package restweb

import (
	"net/http"
	"net/url"
)

type Context struct {
	R      *http.Request
	W      http.ResponseWriter
	Input  url.Values
	Output map[string]interface{}
}

func NewContext(w http.ResponseWriter, r *http.Request) (c *Context) {
	c = &Context{R: r, W: w}
	c.R.ParseForm()
	c.Input = c.R.Form
	c.Output = make(map[string]interface{})
	return
}

func (c *Context) SetSession(key string, value string) {
	session := SessionManager.StartSession(c.W, c.R)
	session.Set(key, value)
}

func (c *Context) GetSession(key string) string {
	sess := SessionManager.StartSession(c.W, c.R)
	return sess.Get(key)
}

func (c *Context) DeleteSession() {
	SessionManager.DeleteSession(c.W, c.R)
}

func (c *Context) Redirect(urlStr string, code int) {
	http.Redirect(c.W, c.R, urlStr, code)
}

func (c *Context) Error(err string, code int) {
	http.Error(c.W, err, code)
}
