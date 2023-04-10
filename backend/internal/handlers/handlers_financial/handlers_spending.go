package handlers_financial

import (
	"encoding/json"
	"net/http"

	middleware "diplomka/internal/handlers/handlers_middleware"
	"diplomka/internal/model"
)

func (s *financial) PostSpending(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	spen := model.Spending{}

	spen.UserID = middleware.GetUserID(r)

	err := json.NewDecoder(r.Body).Decode(&spen)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}

	if err := spen.Validate(); err != nil {
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}

	_, err = s.InsertSpendingService(r.Context(), spen)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
	}
}
