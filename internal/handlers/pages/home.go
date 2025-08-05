package handlers

import (
	"log"
	"net/http"

	"github.com/levifitzpatrick1/blog/internal/utils"
	home "github.com/levifitzpatrick1/blog/web/templates/Pages/Home"
)

type HomeHandler struct {
	Posts  []utils.Post
	Logger *log.Logger
}

func NewHomeHandler(posts []utils.Post, l *log.Logger) *HomeHandler {
	return &HomeHandler{Posts: posts, Logger: l}
}

func (h *HomeHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	var latestPost utils.Post
	if len(h.Posts) > 0 {
		latestPost = h.Posts[0]
	}

	err := home.Home(latestPost).Render(r.Context(), w)
	if err != nil {
		h.Logger.Printf("Error rendering home page: %v", err)
		http.Error(w, "Failed to render home page", http.StatusInternalServerError)
	}
}
