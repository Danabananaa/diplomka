package handlers

import (
	"diplomka/internal/model"
	"diplomka/pkg/log"
	"encoding/json"
	"net/http"
)

type spending struct {
	model.SpendingService
}

func NewSpendingHandlers(spn model.SpendingService) *spending {
	return &spending{
		SpendingService: spn,
	}
}

func (s *spending) PostSpending(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	inc := model.Spending{}

	temp := r.Context().Value(UserKey)

	userID, ok := temp.(int64)
	if !ok {
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}
	inc.UserID = userID
	err := json.NewDecoder(r.Body).Decode(&inc)
	if err != nil {
		log.Printf("json decode: %v", err)
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}

	_, err = s.InsertSpendingService(r.Context(), inc)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
	}
}
