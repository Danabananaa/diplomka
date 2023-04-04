package handlers

import (
	"encoding/json"
	"net/http"

	"diplomka/internal/model"
	"diplomka/pkg/log"
)

func (i *income) GetIncome(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	bet := model.Between{}

	err := json.NewDecoder(r.Body).Decode(&bet)
	if err != nil {
		log.Printf("json decode: %v", err)
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}

	temp := r.Context().Value(UserKey)

	userID, ok := temp.(int64)
	if !ok {
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}

	bet.UserID = userID

	obj, err := i.IncomeService.GetIncomeService(r.Context(), bet)
	if err != nil {
		log.Println(err)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	err = json.NewEncoder(w).Encode(obj)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
}
