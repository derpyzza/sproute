package sproute

import (
	"net/http"
)

type Router struct {
	Routes []Route

	middlewares []Middleware

	CheckNotFound bool
	NotFound      http.HandlerFunc
}

func New() *Router {
	return &Router{}
}

func (rtr *Router) ServeHTTP(res http.ResponseWriter, req *http.Request) {

	for _, rt := range rtr.Routes {
		match := rt.match(req)
		if !match {

			continue
		}

		rt.Handler.ServeHTTP(res, req)
		return
	}

	if rtr.CheckNotFound {
		rtr.NotFound(res, req)
	}

}

func (rtr *Router) Handle(method, path string, handler http.HandlerFunc) {
	rt := Route{
		Path:    path,
		Method:  method,
		Handler: handler,
	}

	rtr.Routes = append(rtr.Routes, rt)
}

func (rtr *Router) Get(path string, handler http.HandlerFunc) {
	rtr.Handle("GET", path, handler)
}

func (rtr *Router) Post(path string, handler http.HandlerFunc) {
	rtr.Handle("POST", path, handler)
}

func (rtr *Router) Put(path string, handler http.HandlerFunc) {
	rtr.Handle("PUT", path, handler)
}

func (rtr *Router) Delete(path string, handler http.HandlerFunc) {
	rtr.Handle("DELETE", path, handler)
}
