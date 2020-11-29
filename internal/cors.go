package internal

import "net/http"

// CORS middleware handles a gorilla middleware for implementing CORS functionality
func CORS(next http.Handler) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		// set headers
		rw.Header().Set("Access-Control-Allow-Headers:", "*")
		rw.Header().Set("Access-Control-Allow-Origin", "*")
		rw.Header().Set("Access-Control-Allow-Methods", "*")

		if r.Method == http.MethodOptions {
			rw.WriteHeader(http.StatusOK)
			return
		}

		next.ServeHTTP(rw, r)
	})
}
