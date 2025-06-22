// internal/handlers/search_handler.go
package handlers

import (
	"log"
	"net/http"

	generatedDB "github.com/levifitzpatrick1/blog/internal/database/generated"
	components "github.com/levifitzpatrick1/blog/web/templates/components"
)

type SearchHandler struct {
	Queries *generatedDB.Queries
	Logger  *log.Logger
}

func NewSearchHandler(q *generatedDB.Queries, l *log.Logger) *SearchHandler {
	return &SearchHandler{Queries: q, Logger: l}
}

func (h *SearchHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	query := r.FormValue("search")

	var posts []generatedDB.ListPostsRow
	var err error

	if query == "" {
		posts, err = h.Queries.ListPosts(ctx)
	} else {
		var searchResults []generatedDB.SearchPostsRow
		searchResults, err = h.Queries.SearchPosts(ctx, generatedDB.SearchPostsParams{
			Title:   query,
			Content: query,
		})

		if err == nil {
			posts = make([]generatedDB.ListPostsRow, len(searchResults))
			for i, p := range searchResults {
				posts[i] = generatedDB.ListPostsRow{
					ID:        p.ID,
					Title:     p.Title,
					Slug:      p.Slug,
					Published: p.Published,
					Excerpt:   p.Excerpt,
				}
			}
		}
	}

	if err != nil {
		log.Printf("Error searching posts with query '%s': %v", query, err)
		http.Error(w, "Error searching posts", http.StatusInternalServerError)
		return
	}

	component := components.PostCardList(posts)
	renderErr := component.Render(ctx, w)
	if renderErr != nil {
		h.Logger.Print("Error rendering search results: ", renderErr)
		http.Error(w, "Failed to render search results.", http.StatusInternalServerError)
	}
}
