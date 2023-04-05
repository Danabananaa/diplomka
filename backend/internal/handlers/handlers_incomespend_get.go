package handlers

import (
	"diplomka/internal/model"
	"diplomka/pkg/log"
	"encoding/json"
	"fmt"
	"net/http"
)

type income_spending struct {
	model.IncomeService
	model.SpendingService
}

func NewIncomeSpendingHandlers(inc model.IncomeService, spn model.SpendingService) *income_spending {
	return &income_spending{
		IncomeService:   inc,
		SpendingService: spn,
	}
}

func (i *income_spending) GetIncomeSpending(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	bet := model.Between{}

	err := json.NewDecoder(r.Body).Decode(&bet)
	if err != nil {
		log.Printf("json decode: %v", err)
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}

	userID := GetID(w, r)

	bet.UserID = userID

	inc, err := i.IncomeService.GetIncomeService(r.Context(), bet)
	if err != nil {
		log.Println(err)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
	spen, err := i.SpendingService.GetSpendingService(r.Context(), bet)
	if err != nil {
		log.Println(err)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
	mergerstruct := model.Income_Spending{
		Income:   inc,
		Spending: spen,
	}
	fmt.Println(spen, inc)

	err = json.NewEncoder(w).Encode(mergerstruct)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
}
