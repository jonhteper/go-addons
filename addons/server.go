package addons

import "net/http"

// AllOrigin es equivalente a "Access-Control-Allow-Origin: *"
func AllOrigin(w *http.ResponseWriter) {
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
}

// AllHeaders es equivalente a "Access-Control-Allow-Headers: *"
func AllHeaders(w *http.ResponseWriter) {
	(*w).Header().Set("Access-Control-Allow-Headers", "*")
}

// AllMethods es equivalente a "Access-Control-Allow-Methods: *"
func AllMethods(w *http.ResponseWriter) {
	(*w).Header().Set("Access-Control-Allow-Methods", "*")
}

// ApplicationJson es equivalente a "Content-Type: application/json; charset=utf-8*"
func ApplicationJson(w *http.ResponseWriter) {
	(*w).Header().Set("Content-Type", "application/json; charset=utf-8")
}
