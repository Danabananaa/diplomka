package handlers

import (
	"bytes"
	"diplomka/internal/model"
	"encoding/base64"
	"encoding/json"
	"io"
	"net/http"
	"os"
	"path/filepath"

	"github.com/google/uuid"
)

func (a *auth) GetUserInfo(w http.ResponseWriter, r *http.Request) {
	user_id := GetID(w, r)
	user, err := a.AuthSerivice.GetUserInfoService(r.Context(), int(user_id))
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
	err = json.NewEncoder(w).Encode(user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (a *auth) PostImage(w http.ResponseWriter, r *http.Request) {
	var imageData model.ImageData
	err := json.NewDecoder(r.Body).Decode(&imageData)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	decoded, err := base64.StdEncoding.DecodeString(imageData.Data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	user_id := GetID(w, r)
	imagename := uuid.New().String() + ".jpg"

	info := model.UserImage{
		UserID:    user_id,
		ImageName: imagename,
	}

	_, err = a.AuthSerivice.AddUserImage(r.Context(), info)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	file, err := os.Create(filepath.Join("./temp", imagename))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer file.Close()

	// Write the decoded image data to the file
	_, err = io.Copy(file, bytes.NewReader(decoded))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Image uploaded successfully"))
}

func (a *auth) GetImage(w http.ResponseWriter, r *http.Request) {
	id := GetID(w, r)

	info, err := a.AuthSerivice.GetUserImageService(r.Context(), int(id))
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	fileName := "./temp/" + info.ImageName
	dir := http.Dir("./temp/")
	fs := http.FileServer(dir)

	if _, err := os.Stat(fileName); os.IsNotExist(err) {
		http.NotFound(w, r)
		return
	}

	fs.ServeHTTP(w, r)
}
