package handlers

import (
	"log"
	"net/http"

	db "github.com/levifitzpatrick1/blog/internal/database/generated"
	templates "github.com/levifitzpatrick1/blog/web/templates/Pages/FRC"
)

type FRCHandler struct {
	Queries *db.Queries
	Logger  *log.Logger
}

func NewFRCHandler(q *db.Queries, l *log.Logger) *FRCHandler {
	return &FRCHandler{Queries: q, Logger: l}
}

func (h *FRCHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	err := templates.FRCPage().Render(r.Context(), w)
	if err != nil {
		h.Logger.Printf("Error rendering FRC page: %v", err)
		http.Error(w, "Failed to render FRC page", http.StatusInternalServerError)
	}
}
