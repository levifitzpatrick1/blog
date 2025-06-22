package handlers

import (
	"log"
	"net/http"

	generatedDB "github.com/levifitzpatrick1/blog/internal/database/generated"
	templates "github.com/levifitzpatrick1/blog/web/templates/Pages/Blogs"
)

type BlogIndexHandler struct {
	Queries *generatedDB.Queries
	Logger  *log.Logger
}

func NewBlogIndexHandler(q *generatedDB.Queries, l *log.Logger) *BlogIndexHandler {
	return &BlogIndexHandler{Queries: q, Logger: l}
}

func (h *BlogIndexHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	s := r.URL.Query().Get("sort")
	posts, err := h.Queries.ListPosts(ctx)
	if err != nil {
		h.Logger.Print("Error fetching posts: ", err)
		http.Error(w, "Whoops! I Couldn't Find The Posts!", http.StatusInternalServerError)
		return
	}

	pageData := templates.BlogsData{
		Posts:       posts,
		CurrentSort: s,
		CurrentPath: r.URL.Path,
	}

	err = templates.Blogs(pageData).Render(ctx, w)
	if err != nil {
		h.Logger.Printf("Error rendering blog index page: %v", err)
	}
}
