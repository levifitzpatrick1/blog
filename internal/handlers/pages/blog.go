package handlers

import (
	"log"
	"net/http"

	db "github.com/levifitzpatrick1/blog/internal/database/generated"
	"github.com/levifitzpatrick1/blog/internal/utils"
	templates "github.com/levifitzpatrick1/blog/web/templates/Pages/Blogs"
)

type BlogIndexHandler struct {
	Queries *db.Queries
	Logger  *log.Logger
}

func NewBlogIndexHandler(q *db.Queries, l *log.Logger) *BlogIndexHandler {
	return &BlogIndexHandler{Queries: q, Logger: l}
}

func (h *BlogIndexHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	s := r.URL.Query().Get("sort")
	posts, err := h.Queries.Listblog(ctx)
	if err != nil {
		h.Logger.Print("Error fetching posts: ", err)
		http.Error(w, "Whoops! I Couldn't Find The Posts!", http.StatusInternalServerError)
		return
	}

	var sendable []utils.BlogItem

	for _, p := range posts {
		tags, err := utils.GetTagsBlog(p.ID, h.Queries)
		if err != nil {
			h.Logger.Print("Error fetching tags for post: ", err)
			continue
		}
		sendable = append(sendable, utils.BlogItem{
			Blog: p,
			Tags: tags,
		})
	}

	pageData := utils.BlogsData{
		Posts:       sendable,
		CurrentSort: s,
		CurrentPath: r.URL.Path,
	}

	err = templates.Blogs(pageData).Render(ctx, w)
	if err != nil {
		h.Logger.Printf("Error rendering blog index page: %v", err)
	}
}
