package router

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

type muxRouter struct{}

var (
	muxDispacher = mux.NewRouter()
)

// MuxRouter Implementation
func MuxRouter() Router {
	return &muxRouter{}
}

func (*muxRouter) GET(uri string, f func(w http.ResponseWriter, r *http.Request)) {
	muxDispacher.HandleFunc(uri, f).Methods("GET")
}

func (*muxRouter) POST(uri string, f func(w http.ResponseWriter, r *http.Request)) {
	muxDispacher.HandleFunc(uri, f).Methods("POST", "OPTIONS")
}

func (*muxRouter) PUT(uri string, f func(w http.ResponseWriter, r *http.Request)) {
	muxDispacher.HandleFunc(uri, f).Methods("PUT", "OPTIONS")
}

func (*muxRouter) DELETE(uri string, f func(w http.ResponseWriter, r *http.Request)) {
	muxDispacher.HandleFunc(uri, f).Methods("DELETE", "OPTIONS")
}

func (*muxRouter) SERVE(port string) {
	fmt.Printf("Mux Http server running on port %v", port)
	http.ListenAndServe(port, muxDispacher)
}
