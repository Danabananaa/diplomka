package handlers_auth

import (
	"net/http"

	middleware "diplomka/internal/handlers/handlers_middleware"
	"diplomka/pkg/log"
)

func (a *auth) Delete(w http.ResponseWriter, r *http.Request) {
	userID := middleware.GetUserID(r)

	if userID <= 0 {
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}

	err := a.AuthService.Delete(r.Context(), userID)
	if err != nil {
		log.Println(err)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
}
