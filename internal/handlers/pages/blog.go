package handlers

import (
	"log"
	"net/http"

	"github.com/levifitzpatrick1/blog/internal/utils/markdown"
	blogs "github.com/levifitzpatrick1/blog/web/templates/Pages/Blogs"
)

type BlogIndexHandler struct {
	Store  *markdown.PostStore
	Logger *log.Logger
}

func NewBlogIndexHandler(store *markdown.PostStore, l *log.Logger) *BlogIndexHandler {
	return &BlogIndexHandler{Store: store, Logger: l}
}

func (h *BlogIndexHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	posts, err := h.Store.GetPosts()
	if err != nil {
		h.Logger.Printf("Error getting posts: %v", err)
		http.Error(w, "Failed to load posts", http.StatusInternalServerError)
		return
	}

	err = blogs.Blogs(posts).Render(r.Context(), w)
	if err != nil {
		h.Logger.Printf("Error rendering blog index page: %v", err)
		http.Error(w, "Failed to render blog index", http.StatusInternalServerError)
	}
}
