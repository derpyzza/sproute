package sproute

import (
	"net/http"
)

type Route struct {
	Path string `json:"path"`

	Method string `json:"method"`

	Handler http.HandlerFunc `json:"handler"`
}

func (rt Route) match(req *http.Request) bool {
	if rt.Method != req.Method {
		return false
	}

	if rt.Path != req.URL.Path {
		return false
	}

	return true
}
