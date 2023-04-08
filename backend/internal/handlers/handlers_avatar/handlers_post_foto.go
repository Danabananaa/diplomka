package handlers_avatar

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"io"
	"net/http"
	"os"
	"path/filepath"

	middleware "diplomka/internal/handlers/handlers_middleware"
	"diplomka/internal/model"

	"github.com/google/uuid"
)

func (a *avatar) UploadFoto(w http.ResponseWriter, r *http.Request) {
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

	user_id := middleware.GetUserID(r)
	imagename := uuid.New().String() + ".jpg"

	info := model.UserImage{
		UserID:    user_id,
		ImageName: imagename,
	}

	_, err = a.AddUserImage(r.Context(), info)
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
