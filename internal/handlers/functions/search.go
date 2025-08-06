package functions

import (
	"log"
	"net/http"
	"strings"

	"github.com/levifitzpatrick1/blog/internal/utils"
	components "github.com/levifitzpatrick1/blog/web/templates/components/blog"
)

type SearchHandler struct {
	Posts  []utils.Post
	Logger *log.Logger
}

func NewSearchHandler(posts []utils.Post, l *log.Logger) *SearchHandler {
	return &SearchHandler{Posts: posts, Logger: l}
}

func (h *SearchHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	query := strings.ToLower(r.FormValue("search"))

	var filteredPosts []utils.Post

	if query == "" {
		filteredPosts = h.Posts
	} else {
		for _, post := range h.Posts {
			if strings.Contains(strings.ToLower(post.Title), query) {
				filteredPosts = append(filteredPosts, post)
			}
		}
	}

	component := components.PostCardList(filteredPosts)
	err := component.Render(r.Context(), w)
	if err != nil {
		h.Logger.Print("Error rendering search results: ", err)
		http.Error(w, "Failed to render search results.", http.StatusInternalServerError)
	}
}
