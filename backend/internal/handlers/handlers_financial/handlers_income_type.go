package handlers_financial

import (
	"encoding/json"
	"net/http"

	"diplomka/internal/model"
	"diplomka/pkg/log"
)

func (f *financial) AllIncomeTypes(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	inc_type, spn_type, err := f.GetAllIncomeSpendingService(r.Context())
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
