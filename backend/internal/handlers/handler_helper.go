package handlers

import (
	"diplomka/internal/model"
	"net/http"
	"time"
)

func GetID(w http.ResponseWriter, r *http.Request) int64 {
	temp := r.Context().Value(UserKey)

	userID, ok := temp.(int64)
	if !ok {
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return 0
	}
	return userID
}

func FindFirstDay(t time.Time) time.Time {
	x := t.Month()
	g := time.Date(int(t.Year()), x, 1, 0, 0, 0, 0, time.UTC)
	return g
}

func ConvertTime(t time.Time) model.CustomTime {
	return model.CustomTime{
		Time: t,
	}
}
