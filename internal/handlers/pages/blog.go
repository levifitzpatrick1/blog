package handlers

import (
	"log"
	"net/http"

	"github.com/levifitzpatrick1/blog/internal/utils"
	blogs "github.com/levifitzpatrick1/blog/web/templates/Pages/Blogs"
)

type BlogIndexHandler struct {
	Posts  []utils.Post
	Logger *log.Logger
}

func NewBlogIndexHandler(posts []utils.Post, l *log.Logger) *BlogIndexHandler {
	return &BlogIndexHandler{Posts: posts, Logger: l}
}

func (h *BlogIndexHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	err := blogs.Blogs(h.Posts).Render(r.Context(), w)
	if err != nil {
		h.Logger.Printf("Error rendering blog index page: %v", err)
		http.Error(w, "Failed to render blog index", http.StatusInternalServerError)
	}
}
