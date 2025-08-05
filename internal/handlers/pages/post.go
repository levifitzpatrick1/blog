package handlers

import (
	"log"
	"net/http"
	"strings"

	"github.com/levifitzpatrick1/blog/internal/utils"
	blogs "github.com/levifitzpatrick1/blog/web/templates/Pages/Blogs"
)

type PostHandler struct {
	Posts  []utils.Post
	Logger *log.Logger
}

func NewPostHandler(posts []utils.Post, l *log.Logger) *PostHandler {
	return &PostHandler{Posts: posts, Logger: l}
}

func (h *PostHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	slug := strings.TrimPrefix(r.URL.Path, "/blog/")
	if slug == "" {
		http.NotFound(w, r)
		return
	}

	var foundPost *utils.Post
	for i := range h.Posts {
		if h.Posts[i].Slug == slug {
			foundPost = &h.Posts[i]
			break
		}
	}

	if foundPost == nil {
		log.Printf("Post with slug '%s' not found.", slug)
		http.NotFound(w, r)
		return
	}

	component := blogs.PostPage(*foundPost)
	err := component.Render(r.Context(), w)
	if err != nil {
		h.Logger.Print("Error rendering template: ", err)
		http.Error(w, "Failed to render page.", http.StatusInternalServerError)
	}
}
