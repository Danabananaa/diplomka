package handlers

import (
	"encoding/json"
	"net/http"

	"diplomka/internal/model"
	"diplomka/pkg/log"
)

type spending struct {
	model.SpendingService
}

func NewSpendingHandlers(s model.SpendingService) *spending {
	return &spending{
		SpendingService: s,
	}
}

func (s *spending) AllSpendingTypes(w http.ResponseWriter, r *http.Request) {
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