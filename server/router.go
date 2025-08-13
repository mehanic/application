package server

import (
	"net/http"

	"github.com/gorilla/mux"
)

var allRoutes = []func(*Context){}

type Router struct {
	router *mux.Router
}

func NewRouter() *Router {
	return &Router{router: mux.NewRouter()}
}

func (r *Router) ListenAndServe(addr string) error {
	srv := &http.Server{
		Handler: r.router,
		Addr:    addr,
	}

	return srv.ListenAndServe()
}

func (r *Router) Group(tpl string) *Router {
	return &Router{router: r.router.PathPrefix(tpl).Subrouter()}
}

func (r *Router) route(path, method string, handler func(*Context)) {
	h := func(w http.ResponseWriter, r *http.Request) {
		c := NewContext(w, r)
		for _, ar := range allRoutes {
			ar(c)
		}
		handler(c)
	}
	if path == "" {
		r.router.Methods(method).HandlerFunc(h)
	} else {
		r.router.Methods(method).Path(path).HandlerFunc(h)
	}
}

func (r *Router) GET(path string, handler func(*Context)) {
	r.route(path, "GET", handler)
}

func (r *Router) POST(path string, handler func(*Context)) {
	r.route(path, "POST", handler)
}

func (r *Router) OPTIONS(path string, handler func(*Context)) {
	r.route(path, "OPTIONS", handler)
}

func (r *Router) PutToAllRoutes(f func(*Context)) {
	allRoutes = append(allRoutes, f)
}

func ConstructRequest(c *Context) {
	c.Header("Access-Control-Allow-Methods", "GET, POST")
	c.Header("Access-Control-Max-Age", "86400")
	c.Header("Access-Control-Allow-Credentials", "true")
	c.Header("Access-Control-Allow-Origin", "*")
	c.Header("Access-Control-Allow-Headers", "Content-Type, Authorization")
}

func DestructRequest(c *Context) {
	c.Header("Access-Control-Allow-Orifin", "*")
}

func (r *Router) Static(PathPrefix, path string) {
	r.router.PathPrefix(PathPrefix).Handler(http.FileServer(http.Dir(path)))
}
