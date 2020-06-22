package middleware

import (
	"log"
	"net/http"
)

func LoggerMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		go log.Println(r.Method +"\t"+ r.RequestURI)
		next.ServeHTTP(w, r)
	})
}
