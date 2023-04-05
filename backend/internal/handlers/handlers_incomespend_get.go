package handlers

import (
	"diplomka/internal/model"
	"diplomka/pkg/log"
	"encoding/json"
	"fmt"
	"net/http"
	"time"
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

	// err := json.NewDecoder(r.Body).Decode(&bet)
	// if err != nil {
	// 	log.Printf("json decode: %v", err)
	// 	http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
	// 	return
	// }

	userID := GetID(w, r)

	bet.UserID = userID

	enddate, err := time.Parse("2006-01-02", time.Now().Format("2006-01-02"))
	if err != nil {
		fmt.Println(err)
		return
	}

	startdate := FindFirstDay(enddate)

	bet.StartDate = model.CustomTime{
		Time: startdate,
	}

	bet.EndDate = model.CustomTime{
		Time: enddate,
	}

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

	err = json.NewEncoder(w).Encode(mergerstruct)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
}

func FindFirstDay(t time.Time) time.Time {
	x := t.Month()
	g := time.Date(int(t.Year()), x, 1, 0, 0, 0, 0, time.UTC)
	return g
}
