package middleware

import (
	"context"
	"net/http"

	"diplomka/internal/model"
	"diplomka/pkg/log"
)

type key string

var UserKey key = "userID"

type middleware struct {
	model.AuthService
}

func NewMiddleware(a model.AuthService) *middleware {
	return &middleware{
		AuthService: a,
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

		userID, err := m.AuthService.Verification(token[0])
		if err != nil {
			log.Println(err)
			http.Error(w, http.StatusText(http.StatusUnauthorized), http.StatusUnauthorized)
			return
		}

		ctx := context.WithValue(r.Context(), UserKey, userID)

		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func GetUserID(r *http.Request) int64 {
	userID, ok := r.Context().Value(UserKey).(int64)
	if !ok {
		return 0
	}
	return userID
}
