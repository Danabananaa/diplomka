package handlers_avatar

import (
	"fmt"
	"net/http"

	middleware "diplomka/internal/handlers/handlers_middleware"
)

func (a *avatar) DeleteFoto(w http.ResponseWriter, r *http.Request) {
	userID := middleware.GetUserID(r)
	if userID < 1 {
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}
	err := a.AvatarService.DeleteImageService(r.Context(), userID)
	if err != nil {
		fmt.Println(err)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
}
