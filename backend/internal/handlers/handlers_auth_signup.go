package handlers

import (
	"diplomka/internal/model"
	"diplomka/pkg/log"
	"encoding/json"
	"net/http"
)

func (a *auth) SignUp(w http.ResponseWriter, r *http.Request) {
	// w.Header().Set("Content-Type", "application/json")
	user := model.User{}
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		log.Printf("json decode: %v", err)
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}

	if err := a.AuthSerivice.SignUp(r.Context(), user); err != nil {
		log.Printf("service SignUp: %v", err)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}
