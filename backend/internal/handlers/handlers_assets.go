package handlers

import (
	"diplomka/internal/model"
	"diplomka/pkg/log"
	"encoding/json"
	"net/http"
)

type assets_liab struct {
	model.AssetsLiabService
}

func NewAssetsHandlers(ass model.AssetsLiabService) *assets_liab {
	return &assets_liab{
		AssetsLiabService: ass,
	}
}

func (ass *assets_liab) AddAssets(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	ast := model.Assets{}

	temp := r.Context().Value(UserKey)

	userID, ok := temp.(int64)
	if !ok {
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}
	ast.UserID = userID
	err := json.NewDecoder(r.Body).Decode(&ast)
	if err != nil {
		log.Printf("json decode: %v", err)
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}

	_, err = ass.AddAssetsService(r.Context(), ast)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
	}
}

func (ass *assets_liab) AddLiab(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	lia := model.Liabilities{}

	temp := r.Context().Value(UserKey)

	userID, ok := temp.(int64)
	if !ok {
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}
	lia.UserID = userID
	err := json.NewDecoder(r.Body).Decode(&lia)
	if err != nil {
		log.Printf("json decode: %v", err)
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}

	_, err = ass.AddLiabService(r.Context(), lia)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
	}
}
