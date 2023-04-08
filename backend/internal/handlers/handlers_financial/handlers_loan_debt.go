package handlers_financial

import (
	"diplomka/internal/model"
	"diplomka/pkg/log"
	"encoding/json"
	"net/http"

	middleware "diplomka/internal/handlers/handlers_middleware"
)

func (f *financial) AddLoanDebt(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	ast := model.Loan_Debt{}

	ast.UserID = middleware.GetUserID(r)
	err := json.NewDecoder(r.Body).Decode(&ast)
	if err != nil {
		log.Printf("json decode: %v", err)
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}

	if err := ast.Validate(); err != nil {
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}

	_, err = f.AddLoanDebtService(r.Context(), ast)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
	}
}

func (f *financial) GetLoanDebt(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	bet, err := getFilter(r)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}

	loan, debt, err := f.GetLoanDebtService(r.Context(), bet)
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
