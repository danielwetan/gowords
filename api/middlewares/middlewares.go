package middlewares

import (
	"errors"
	"net/http"

	"github.com/danielwetan/gowords/auth"
	"github.com/danielwetan/gowords/responses"
)

func SetMiddlewareJSON(next http.HandleFunc) http.HandleFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		next(w, r)
	}
}

funct SetMiddlewareAuthentification(next http.HandleFunc) http.HandleFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		err := auth.TokenValid(r)
		if err != nil {
			responses.ERROR(w, http.StatusUnauthorized, errros.New("Unauthorized"))
			return
		}
		next(w, r)
	}
}
