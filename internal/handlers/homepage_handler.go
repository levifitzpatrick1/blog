package handlers

import (
	"database/sql"
	"log"
	"net/http"

	generatedDB "github.com/levifitzpatrick1/blog/internal/database/generated"
	pages "github.com/levifitzpatrick1/blog/web/templates/Pages/Home"
	templates "github.com/levifitzpatrick1/blog/web/templates/Pages/Home"
)

type HomeHandler struct {
	Queries *generatedDB.Queries
	Logger  *log.Logger
}

func NewHomeHandler(q *generatedDB.Queries, l *log.Logger) *HomeHandler {
	return &HomeHandler{Queries: q, Logger: l}
}

func (h *HomeHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	latestPost, err := h.Queries.GetLatestPost(ctx)
	if err != nil {
		if err != sql.ErrNoRows {
			h.Logger.Printf("Error fetching latest post: %v", err)
		}
		latestPost = generatedDB.GetLatestPostRow{}
	}

	pageData := pages.HomeData{
		LatestPost: latestPost,
	}

	err = templates.Home(pageData).Render(r.Context(), w)
	if err != nil {
		h.Logger.Printf("Error rendering home page: %v", err)
		http.Error(w, "Failed to render home page", http.StatusInternalServerError)
	}
}
