package sproute

import "net/http"

type Middleware func(handler http.HandlerFunc) http.HandlerFunc
