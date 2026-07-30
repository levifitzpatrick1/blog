package handlers

import (
	"log"
	"net/http"
	"strings"

	"github.com/levifitzpatrick1/levifitzpatrick.page/internal/utils"
	"github.com/levifitzpatrick1/levifitzpatrick.page/internal/utils/markdown"
	blogs "github.com/levifitzpatrick1/levifitzpatrick.page/web/templates/Pages/Blogs"
)

type PostHandler struct {
	Store  *markdown.PostStore
	Logger *log.Logger
}

func NewPostHandler(store *markdown.PostStore, l *log.Logger) *PostHandler {
	return &PostHandler{Store: store, Logger: l}
}

func (h *PostHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	slug := strings.TrimPrefix(r.URL.Path, "/blog/")
	if slug == "" {
		http.NotFound(w, r)
		return
	}

	posts, err := h.Store.GetPosts()
	if err != nil {
		h.Logger.Printf("Error getting posts: %v", err)
		http.Error(w, "Failed to load posts", http.StatusInternalServerError)
		return
	}

	var foundPost *utils.Post
	for i := range posts {
		if posts[i].Slug == slug {
			foundPost = &posts[i]
			break
		}
	}

	if foundPost == nil {
		log.Printf("Post with slug '%s' not found.", slug)
		http.NotFound(w, r)
		return
	}

	component := blogs.PostPage(*foundPost)
	err = component.Render(r.Context(), w)
	if err != nil {
		h.Logger.Print("Error rendering template: ", err)
		http.Error(w, "Failed to render page.", http.StatusInternalServerError)
	}
}
