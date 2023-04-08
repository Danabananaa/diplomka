package handlers

import (
	"net/http"

	"diplomka/internal/handlers/handlers_avatar"
	"diplomka/internal/model"

	"github.com/gorilla/mux"
)

func InitAvatarHandlers(r *mux.Router, ava model.AvatarService) {
	handlers := handlers_avatar.NewAvatarHandlers(ava)

	// fs := http.FileServer(model.NewFileSystem(http.Dir("./static/images/")))
	fs := http.FileServer(http.Dir("./static/images/"))

	r.Handle("/images/", http.StripPrefix("/images", fs))

	r.HandleFunc("/images", handlers.UploadFoto).Methods(http.MethodPost)
	r.HandleFunc("/images", handlers.DeleteFoto).Methods(http.MethodDelete)
}
