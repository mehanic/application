package server

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

type Context struct {
	ResponseWriter http.ResponseWriter
	Request        *http.Request
}

func NewContext(w http.ResponseWriter, r *http.Request) *Context {
	return &Context{w, r}
}

func (c *Context) Header(name, value string) {
	c.ResponseWriter.Header().Set(name, value)
}

func (c *Context) Param(name string) string {
	params := mux.Vars(c.Request)
	return params[name]
}

func (c *Context) RenderError(status int, err error) {
	http.Error(c.ResponseWriter, err.Error(), status)
}

func (c *Context) setStatus(s int) {
	c.ResponseWriter.WriteHeader(s)
}

func (c *Context) RenderJSON(status int, j interface{}) {
	c.Header("Content-Type", "application/json")
	c.setStatus(status)
	data, err := json.Marshal(j)
	if err != nil {
		c.setStatus(http.StatusInternalServerError)
		return
	}
	c.ResponseWriter.Write(data)
}
