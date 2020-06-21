package router

import "net/http"

// Router interface
type Router interface {
	GET(uri string, f func(w http.ResponseWriter, r *http.Request))
	GETBY(uri string, f func(w http.ResponseWriter, r *http.Request))
	POST(uri string, f func(w http.ResponseWriter, r *http.Request))
	SERVE(port string)
}
