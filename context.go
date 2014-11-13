package restweb

import (
	"net/http"
)

type Context struct {
	Requset  *http.Request
	Response http.ResponseWriter
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

func (c *Context) Redirct(urlStr string, statuscode int) {
	http.Redirect(c.Response, c.Requset, urlStr, statuscode)

}
