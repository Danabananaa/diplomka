package handlers

import (
	"diplomka/internal/model"
	"diplomka/pkg/log"
	"encoding/json"
	"net/http"
)

type loan_debt struct {
	model.LoanDebtService
}

func NewLoanDebtHandlers(ld model.LoanDebtService) *loan_debt {
	return &loan_debt{
		LoanDebtService: ld,
	}
}

func (ld *loan_debt) AddLoanDebt(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	ast := model.Loan_Debt{}

	userID := GetID(w, r)
	ast.UserID = userID
	err := json.NewDecoder(r.Body).Decode(&ast)
	if err != nil {
		log.Printf("json decode: %v", err)
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}

	_, err = ld.AddLoanDebtService(r.Context(), ast)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
	}
}

func (ld *loan_debt) GetLoanDebt(w http.ResponseWriter, r *http.Request) {
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

	enddate := GetTime()

	startdate := FindFirstDay(enddate)

	bet.StartDate = ConvertTime(startdate)

	bet.EndDate = ConvertTime(enddate)

	loan, debt, err := ld.LoanDebtService.GetLoanDebtService(r.Context(), bet)
	if err != nil {
		log.Println(err)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
	mergestruct := model.MergeStruct{
		LoanArr: loan,
		DebtArr: debt,
	}

	err = json.NewEncoder(w).Encode(mergestruct)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
}
