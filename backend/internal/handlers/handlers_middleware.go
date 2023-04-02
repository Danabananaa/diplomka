package handlers

import (
	"context"
	"net/http"

	"diplomka/internal/model"
	"diplomka/pkg/log"
)

type key string

var UserKey key = "userID"

type middleware struct {
	authService model.AuthSerivice
}

func NewMiddleware(a model.AuthSerivice) *middleware {
	return &middleware{
		authService: a,
	}
}

func (m *middleware) PanicRecover(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if err := recover(); err != nil {
				w.Header().Set("Connection", "close")
				http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
				return
			}
		}()
		next.ServeHTTP(w, r)
	})
}

func (m *middleware) RequireAuthentication(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		token, ok := r.Header["Authorization"]
		if !ok {
			http.Error(w, http.StatusText(http.StatusUnauthorized), http.StatusUnauthorized)
			return
		}

		userID, err := m.authService.Verification(token[0])
		if err != nil {
			log.Println(err)
			http.Error(w, http.StatusText(http.StatusUnauthorized), http.StatusUnauthorized)
			return
		}

		ctx := context.WithValue(r.Context(), UserKey, userID)

		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
