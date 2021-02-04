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

// Route es un objeto que contiene los datos para crear una ruta en una API REST, puede utilizarse en conjunto
// con el objeto Router de github.com/gorilla/mux.
type Route struct {
	Name       string
	Methods    []string
	Pattern    string
	HandleFunc http.HandlerFunc
}
