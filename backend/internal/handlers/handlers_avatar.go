package handlers

import (
	"net/http"

	"diplomka/internal/model"

	"github.com/gorilla/mux"
)

func NewAvatarHandlers(r *mux.Router) {
	fs := http.FileServer(model.NewFileSystem(http.Dir("./static/images/")))

	r.Handle("/images/", http.StripPrefix("/images", fs))
}
