package handlers

import (
	"log"
	"net/http"

	db "github.com/levifitzpatrick1/blog/internal/database/generated"
	Pages "github.com/levifitzpatrick1/blog/web/templates/Pages/Resume"
)

type ResumeHandler struct {
	Queries *db.Queries
	Logger  *log.Logger
}

func NewResumeHandler(q *db.Queries, l *log.Logger) *ResumeHandler {
	return &ResumeHandler{Queries: q, Logger: l}
}

func (h *ResumeHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	err := Pages.ResumePage().Render(r.Context(), w)
	if err != nil {
		h.Logger.Printf("Error rendering home page: %v", err)
		http.Error(w, "Failed to render home page", http.StatusInternalServerError)
	}
}
