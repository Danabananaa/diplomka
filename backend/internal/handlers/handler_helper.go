package handlers

import (
	"net/http"
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
