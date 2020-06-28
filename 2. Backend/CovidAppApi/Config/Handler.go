package config

import "net/http"

// AddHeaders to Response
func AddHeaders(response *http.ResponseWriter) {
	(*response).Header().Set("Content-Type", "application/json")
	(*response).Header().Set("Access-Control-Allow-Origin", "*")
	(*response).Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	(*response).Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
}
