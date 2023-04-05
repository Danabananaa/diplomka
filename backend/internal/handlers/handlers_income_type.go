package handlers

import (
	"diplomka/internal/model"
	"diplomka/pkg/log"
	"encoding/json"
	"net/http"
)

type income_spend_type struct {
	model.IncomeSpendTypeService
}

func NewIncomeSpendTypeHandlers(s model.IncomeSpendTypeService) *income_spend_type {
	return &income_spend_type{
		IncomeSpendTypeService: s,
	}
}

func (s *income_spend_type) AllIncomeTypes(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	inc_type, spn_type, err := s.IncomeSpendTypeService.GetAllIncSpendingService(r.Context())
	if err != nil {
		log.Println(err)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
	mergestruct := model.IncomeSpendType{
		Type_income:   inc_type,
		Type_spending: spn_type,
	}

	err = json.NewEncoder(w).Encode(mergestruct)
	if err != nil {
		log.Println(err)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
}
