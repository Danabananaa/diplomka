package handlers

import (
	"net/http"

	"diplomka/internal/handlers/handlers_avatar"
	"diplomka/internal/model"

	"github.com/gorilla/mux"
)

func InitAvatarHandlers(r *mux.Router) {
	handlers := handlers_avatar.NewAvatarHandlers()

	fs := http.FileServer(model.NewFileSystem(http.Dir("./static/images/")))

	r.Handle("/images/", http.StripPrefix("/images", fs)).Methods(http.MethodGet)
	r.HandleFunc("/images/upload", handlers.UploadFoto).Methods(http.MethodPost)
	r.HandleFunc("/images/upload", handlers.DeleteFoto).Methods(http.MethodDelete)
}
