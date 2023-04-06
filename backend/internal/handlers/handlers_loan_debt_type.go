package handlers

import (
	"diplomka/internal/model"
	"diplomka/pkg/log"
	"encoding/json"
	"net/http"
)

type loan_debt_type struct {
	model.LoanDebtTypeService
}

func NewLoanDebtTypeHandlers(alt model.LoanDebtTypeService) *loan_debt_type {
	return &loan_debt_type{
		LoanDebtTypeService: alt,
	}
}

func (alt *loan_debt_type) GetAllLoanDebtType(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	obj, err := alt.LoanDebtTypeService.GetAllLoanDebtTypeService(r.Context())
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
