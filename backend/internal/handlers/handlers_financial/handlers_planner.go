package handlers_financial

import (
	"encoding/json"
	"log"
	"net/http"

	middleware "diplomka/internal/handlers/handlers_middleware"
	"diplomka/internal/model"
)

func (f *financial) GetPlanner(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
}

func (f *financial) PostPlanner(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	planner := model.SpendingPlanner{}

	planner.UserID = middleware.GetUserID(r)

	err := json.NewDecoder(r.Body).Decode(&planner)
	if err != nil {
		log.Println(err)
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}

	if err := planner.Validate(); err != nil {
		log.Println(err)
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}

	err = f.AddSpendingPLanner(r.Context(), planner)
	if err != nil {
		log.Println(err)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
}
