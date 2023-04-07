package handlers

import (
	"fmt"
	"net/http"
	"time"

	"diplomka/internal/model"
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

func GetTime() time.Time {
	var zeroTime time.Time
	time, err := time.Parse("2006-01-02", time.Now().Format("2006-01-02"))
	if err != nil {
		fmt.Errorf("Error was ocured during time parsing %v", err)
		return zeroTime
	}
	return time
}
