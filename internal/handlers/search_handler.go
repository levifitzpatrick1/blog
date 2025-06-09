package handlers

import (
	"log"
	"net/http"

	generatedDB "github.com/levifitzpatrick1/blog/internal/database/generated"
	webTemplates "github.com/levifitzpatrick1/blog/web/templates"
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
	query := r.URL.Query().Get("query")

	var posts []generatedDB.ListPostsRow
	var err error

	if query == "" {
		posts, err = h.Queries.ListPosts(ctx)
	} else {
		tp, te := h.Queries.SearchPosts(ctx, generatedDB.SearchPostsParams{
			Title:   query,
			Content: query,
		})
		if te == nil {
			posts = make([]generatedDB.ListPostsRow, len(tp))
			for i, p := range tp {
				posts[i] = generatedDB.ListPostsRow{
					ID:        p.ID,
					Title:     p.Title,
					Slug:      p.Slug,
					Published: p.Published,
					Excerpt:   p.Excerpt,
				}
			}
		}
		err = te
	}

	if err != nil {
		log.Printf("Error searching posts with query '%s': %v", query, err)
		http.Error(w, "Error searching posts", http.StatusInternalServerError)
		return
	}
	postsData := struct {
		Posts []generatedDB.ListPostsRow
	}{
		Posts: posts,
	}
	webTemplates.RenderTemplate(w, "post_card_list.html", postsData)
}
