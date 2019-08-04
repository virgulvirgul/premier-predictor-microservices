package auth

import (
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func (a *authenticator) HttpMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		p, _ := mux.CurrentRoute(r).GetPathTemplate()
		if p == "/health" || p == "/legacy" {
			next.ServeHTTP(w, r)
			return
		}

		token := r.Header.Get("Authorization")

		err := a.doAuth(token)
		if err != nil {
			log.Printf("auth error: %s", err)
			w.WriteHeader(http.StatusUnauthorized)
			return
		}

		next.ServeHTTP(w, r)
	})
}
