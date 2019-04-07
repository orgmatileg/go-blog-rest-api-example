package middleware

import "net/http"

// CheckAuth func Middleware
func CheckAuth(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		// w.Header().Add("Access-Control-Allow-Origin", "*")

		next.ServeHTTP(w, r)
	})
}
