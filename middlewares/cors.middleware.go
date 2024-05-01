package middlewares

import "net/http"

func Cors(handler http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		w.Header().Set("Allow-Control-Access-Origin", "*")
		handler(w, req)
	}
}
