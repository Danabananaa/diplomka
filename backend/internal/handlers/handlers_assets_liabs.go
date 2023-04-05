package handlers

import (
	"diplomka/internal/model"
	"diplomka/pkg/log"
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

type assets_liab struct {
	model.AssetsLiabService
}

func NewAssetsHandlers(asl model.AssetsLiabService) *assets_liab {
	return &assets_liab{
		AssetsLiabService: asl,
	}
}

func (asl *assets_liab) AddAssetsLiabs(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	ast := model.Assets_or_Liabs{}

	userID := GetID(w, r)
	ast.UserID = userID
	err := json.NewDecoder(r.Body).Decode(&ast)
	if err != nil {
		log.Printf("json decode: %v", err)
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}
	_, err = asl.AddAssetsLiabsService(r.Context(), ast)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
	}
}

func (asl *assets_liab) GetAssetLiab(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	bet := model.Between{}

	// err := json.NewDecoder(r.Body).Decode(&bet)
	// if err != nil {
	// 	log.Printf("json decode: %v", err)
	// 	http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
	// 	return
	// }

	userID := GetID(w, r)

	bet.UserID = userID

	enddate, err := time.Parse("2006-01-02", time.Now().Format("2006-01-02"))
	if err != nil {
		fmt.Println(err)
		return
	}

	startdate := FindFirstDay(enddate)

	bet.StartDate = model.CustomTime{
		Time: startdate,
	}

	bet.EndDate = model.CustomTime{
		Time: enddate,
	}

	asset, liab, err := asl.AssetsLiabService.GetAssetsLiabsService(r.Context(), bet)
	if err != nil {
		log.Println(err)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
	mergestruct := model.MergeStruct{
		AssetsArr:      asset,
		LiabilitiesArr: liab,
	}

	err = json.NewEncoder(w).Encode(mergestruct)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
}
