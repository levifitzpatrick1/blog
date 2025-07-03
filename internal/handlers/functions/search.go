package functions

import (
	"fmt"
	"log"
	"net/http"

	db "github.com/levifitzpatrick1/blog/internal/database/generated"
	"github.com/levifitzpatrick1/blog/internal/utils"
	components "github.com/levifitzpatrick1/blog/web/templates/components/blog"
)

type SearchHandler struct {
	Queries *db.Queries
	Logger  *log.Logger
}

func NewSearchHandler(q *db.Queries, l *log.Logger) *SearchHandler {
	return &SearchHandler{Queries: q, Logger: l}
}

func (h *SearchHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	query := r.FormValue("search")

	var posts []db.Blog
	var err error

	if query == "" {
		posts, err = h.Queries.Listblog(ctx)
	} else {
		searchPattern := fmt.Sprintf("%%%s%%", query)
		posts, err = h.Queries.Searchblog(ctx, db.SearchblogParams{
			Title:   searchPattern,
			Content: searchPattern,
		})
	}

	if err != nil {
		log.Printf("Error searching posts with query '%s': %v", query, err)
		http.Error(w, "Error searching posts", http.StatusInternalServerError)
		return
	}

	var sendable []utils.BlogItem

	for _, p := range posts {
		tags, err := utils.GetTagsBlog(p.ID, h.Queries)
		if err != nil {
			h.Logger.Printf("Error getting tags for post %d: %v", p.ID, err)
			continue
		}
		sendable = append(sendable, utils.BlogItem{
			Blog: p,
			Tags: tags,
		})
	}

	component := components.PostCardList(sendable)
	renderErr := component.Render(ctx, w)
	if renderErr != nil {
		h.Logger.Print("Error rendering search results: ", renderErr)
		http.Error(w, "Failed to render search results.", http.StatusInternalServerError)
	}
}
