package handlers_auth

import (
	"encoding/json"
	"net/http"

	middleware "diplomka/internal/handlers/handlers_middleware"
)

func (a *auth) Profile(w http.ResponseWriter, r *http.Request) {
	userID := middleware.GetUserID(r)

	if userID <= 0 {
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}

	profile, err := a.AuthService.Profile(r.Context(), userID)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	err = json.NewEncoder(w).Encode(profile)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
}
