package handlers

import (
	"diplomka/internal/model"
	"diplomka/pkg/log"
	"encoding/json"
	"net/http"
)

type assliatype struct {
	model.AssLiaTypeService
}

func NewAssLiaTypeHandlers(alt model.AssLiaTypeService) *assliatype {
	return &assliatype{
		AssLiaTypeService: alt,
	}
}

func (alt *assliatype) GetAllAssLiaType(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	obj, err := alt.AssLiaTypeService.GetAllAssLiaTypeService(r.Context())
	if err != nil {
		log.Println(err)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	err = json.NewEncoder(w).Encode(obj)
	if err != nil {
		log.Println(err)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
}
