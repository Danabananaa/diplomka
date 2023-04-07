package handlers_financial

import (
	"encoding/json"
	"net/http"

	"diplomka/internal/model"
	"diplomka/pkg/log"
)

func (f *financial) GetIncomeSpending(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	bet, err := getFilter(r)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}

	inc, err := f.GetIncomeService(r.Context(), bet)
	if err != nil {
		log.Println(err)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
	spen, err := f.GetSpendingService(r.Context(), bet)
	if err != nil {
		log.Println(err)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
	mergerstruct := model.Income_Spending{
		Income:   inc,
		Spending: spen,
	}

	err = json.NewEncoder(w).Encode(mergerstruct)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
}
