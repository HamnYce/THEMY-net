package apiv1middleware

import (
	"fmt"
	"net/http"
)

func WrapperMiddleWare(handlerFunc http.Handler) http.Handler {
	return MethodCheckMiddleware(http.MethodPost, JSONResponseMiddleware(CORSMiddleware(handlerFunc)))
}

func JSONResponseMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		next.ServeHTTP(w, r)
	})
}

func CORSMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// TODO: Fix this for security reasons to avoid CSRF attacks
		w.Header().Set("Access-Control-Allow-Origin", "*")
		next.ServeHTTP(w, r)
	})
}

func MethodCheckMiddleware(method string, next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != method {
			w.Write([]byte(
				fmt.Sprintf("\"error\": \"Only %s method is allowed\"", method),
			))
			return
		}
		next.ServeHTTP(w, r)
	})
}

func LoggerMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Printf("Request URI: %s\n", r.RequestURI)
		next.ServeHTTP(w, r)
	})
}
