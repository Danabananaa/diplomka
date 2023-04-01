package handlers

import (
	"diplomka/internal/model"
	"diplomka/pkg/log"
	"encoding/json"
	"net/http"
)

type income_type struct {
	model.IncomeTypeService
}

func NewIncomeTypeHandlers(s model.IncomeTypeService) *income_type {
	return &income_type{
		IncomeTypeService: s,
	}
}

func (s *income_type) AllIncomeTypes(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	obj, err := s.IncomeTypeService.GetAllIncomeRepo(r.Context())
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
