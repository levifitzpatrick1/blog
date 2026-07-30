package handlers

import (
	"log"
	"net/http"

	"github.com/levifitzpatrick1/levifitzpatrick.page/internal/utils"
	"github.com/levifitzpatrick1/levifitzpatrick.page/internal/utils/markdown"
	home "github.com/levifitzpatrick1/levifitzpatrick.page/web/templates/Pages/Home"
)

type HomeHandler struct {
	Store  *markdown.PostStore
	Logger *log.Logger
}

func NewHomeHandler(store *markdown.PostStore, l *log.Logger) *HomeHandler {
	return &HomeHandler{Store: store, Logger: l}
}

func (h *HomeHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	posts, err := h.Store.GetPosts()
	if err != nil {
		h.Logger.Printf("Error getting posts: %v", err)
		http.Error(w, "Failed to load posts", http.StatusInternalServerError)
		return
	}

	var latestPost utils.Post
	if len(posts) > 0 {
		latestPost = posts[0]
	}

	err = home.Home(latestPost).Render(r.Context(), w)
	if err != nil {
		h.Logger.Printf("Error rendering home page: %v", err)
		http.Error(w, "Failed to render home page", http.StatusInternalServerError)
	}
}
