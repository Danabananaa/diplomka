package handlers

import (
	"diplomka/internal/model"
	"net/http"
)

type Middleware struct {
	model.AuthSerivice
}

func NewMiddleware(a model.AuthSerivice) *Middleware {
	return &Middleware{
		AuthSerivice: a,
	}
}

func (m *Middleware) PanicRecover(next http.Handler) http.Handler {
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

func (m *Middleware) RequireAuthentication(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		token := r.Header.Get("Authorization")
		_, err := m.AuthSerivice.Refresh(r.Context(), token)
		if err != nil {
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			return
		}
	})
}
