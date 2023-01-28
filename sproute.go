package sproute

import (
	"errors"
	"fmt"
	"net/http"
)

type Router struct {
	Routes []Route

	CheckNotFound bool
	NotFound      http.HandlerFunc

	EnableLogging bool
}

func New() *Router {
	r := &Router{}
	r.NotFound = func(w http.ResponseWriter, r *http.Request) {
		http.Error(w, errors.New("route not found :(").Error(), http.StatusNotFound)
	}
	return r
}

/*
take path, check for path variables, check if requested path matches with
*/
// func ParsePaths(requested, path string) string {
// return "ok"
// }

func (rtr *Router) ServeHTTP(res http.ResponseWriter, req *http.Request) {

	for _, rt := range rtr.Routes {

		match := rt.match(req)
		if !match {
			continue
		}
		fmt.Printf("Matched! %s\n", rt.Path)

		rt.Handler.ServeHTTP(res, req)
		return
	}

	if rtr.CheckNotFound {
		fmt.Println("Not Found :(")
		rtr.NotFound(res, req)
	}

}

func (rtr *Router) Handle(method, path string, handler http.Handler) {

	rt := Route{
		Path:    path,
		Method:  method,
		Handler: handler,
	}

	if rtr.EnableLogging {
		fmt.Println("> Adding path: " + path + " [" + method + "]")
	}

	rtr.Routes = append(rtr.Routes, rt)
}

func (rtr *Router) Get(path string, handler http.Handler) {
	rtr.Handle("GET", path, handler)
}

func (rtr *Router) Post(path string, handler http.Handler) {
	rtr.Handle("POST", path, handler)
}

func (rtr *Router) Put(path string, handler http.Handler) {
	rtr.Handle("PUT", path, handler)
}

func (rtr *Router) Delete(path string, handler http.Handler) {
	rtr.Handle("DELETE", path, handler)
}
