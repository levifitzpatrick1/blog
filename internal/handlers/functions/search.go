package functions

import (
	"log"
	"net/http"
	"strings"

	"github.com/levifitzpatrick1/blog/internal/utils"
	components "github.com/levifitzpatrick1/blog/web/templates/components/blog"
)

// SearchHandler serves search requests for blog posts.
// It performs case-insensitive matching against post titles and tags.
// Partial matches are supported (substring matching).
type SearchHandler struct {
	Posts  []utils.Post
	Logger *log.Logger
}

func NewSearchHandler(posts []utils.Post, l *log.Logger) *SearchHandler {
	return &SearchHandler{Posts: posts, Logger: l}
}

func (h *SearchHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// Read and normalize the query.
	raw := r.FormValue("search")
	query := strings.TrimSpace(strings.ToLower(raw))

	// If empty, return all posts.
	if query == "" {
		component := components.PostCardList(h.Posts)
		if err := component.Render(r.Context(), w); err != nil {
			h.Logger.Print("Error rendering search results: ", err)
			http.Error(w, "Failed to render search results.", http.StatusInternalServerError)
		}
		return
	}

	// Support multiple tokens (space-separated). We will include a post if
	// any token matches the title or any tag (OR semantics between tokens).
	tokens := strings.Fields(query)

	var filteredPosts []utils.Post
	for _, post := range h.Posts {
		// Precompute lowercase title and tag strings for efficient checks.
		title := strings.ToLower(post.Title)

		// Build a single lowercase string of all tag names for quick partial checks,
		// and also keep individual tag names available for more precise checks.
		var tagNames []string
		for _, t := range post.Tags {
			tagNames = append(tagNames, strings.ToLower(t.Name))
		}

		matched := false
		for _, tok := range tokens {
			// Allow token to match title or any tag (partial match).
			if strings.Contains(title, tok) {
				matched = true
				break
			}
			for _, tn := range tagNames {
				if strings.Contains(tn, tok) {
					matched = true
					break
				}
			}
			if matched {
				break
			}
		}

		if matched {
			filteredPosts = append(filteredPosts, post)
		}
	}

	// Render the filtered list component.
	component := components.PostCardList(filteredPosts)
	if err := component.Render(r.Context(), w); err != nil {
		h.Logger.Print("Error rendering search results: ", err)
		http.Error(w, "Failed to render search results.", http.StatusInternalServerError)
	}
}
