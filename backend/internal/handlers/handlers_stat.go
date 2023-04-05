package handlers

import (
	"encoding/json"
	"net/http"

	"diplomka/internal/model"
)

type stat struct {
	model.StatisticsService
}

func NewStatisticsHandlers(s model.StatisticsService) *stat {
	return &stat{
		StatisticsService: s,
	}
}

func (s *stat) GetStat(w http.ResponseWriter, r *http.Request) {
	b := model.Between{}
	err := json.NewDecoder(r.Body).Decode(&b)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	b.UserID = 2

	stat, err := s.StatisticsService.GetStatistics(r.Context(), b)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = json.NewEncoder(w).Encode(stat)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
