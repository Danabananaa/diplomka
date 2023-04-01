package handlers

import (
	"diplomka/internal/model"
	"diplomka/pkg/log"
	"encoding/json"
	"net/http"
)

type spending_type struct {
	model.SpendingService
}

func NewSpendingHandlers(s model.SpendingService) *spending_type {
	return &spending_type{
		SpendingService: s,
	}
}

func (s *spending_type) AllSpendingTypes(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	obj, err := s.SpendingService.GetAllSpendingRepo(r.Context())
	if err != nil {
		log.Println(err)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	err = json.NewEncoder(w).Encode(obj)
	if err != nil {
		log.Println(err)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
}
