package handlers_financial

import (
	"encoding/json"
	"net/http"

	middleware "diplomka/internal/handlers/handlers_middleware"
	"diplomka/internal/model"
	"diplomka/pkg/log"
)

func (a *financial) PostIncome(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	inc := model.Income{}

	userID := middleware.GetUserID(r)
	inc.UserID = userID
	err := json.NewDecoder(r.Body).Decode(&inc)
	if err != nil {
		log.Printf("json decode: %v", err)
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}

	if err := inc.Validate(); err != nil {
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}

	_, err = a.InsertIncomeService(r.Context(), inc)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
	}
}
