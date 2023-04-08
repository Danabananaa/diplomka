package handlers_avatar

import (
	"encoding/json"
	"net/http"

	middleware "diplomka/internal/handlers/handlers_middleware"
)

type image struct {
	Name string `json:"name"`
}

func (a *avatar) GetFoto(w http.ResponseWriter, r *http.Request) {
	userID := middleware.GetUserID(r)
	picName, err := a.GetUserImageService(r.Context(), userID)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(image{
		Name: picName.ImageName,
	})
}
