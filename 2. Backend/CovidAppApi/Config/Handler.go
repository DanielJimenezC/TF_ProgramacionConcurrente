package config

import "net/http"

// AddHeaders to Response
func AddHeaders(response *http.ResponseWriter) {
	(*response).Header().Set("Content-Type", "application/json")
	(*response).Header().Set("Access-Control-Allow-Origin", "*")
}
