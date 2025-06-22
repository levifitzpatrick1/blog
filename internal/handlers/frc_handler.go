package handlers

import (
	"log"
	"net/http"

	templates "github.com/levifitzpatrick1/blog/web/templates/Pages/FRC"
)

type FRCHandler struct {
	Logger *log.Logger
}

func NewFRCHandler(l *log.Logger) *FRCHandler {
	return &FRCHandler{Logger: l}
}

func (h *FRCHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	err := templates.FRCPage().Render(r.Context(), w)
	if err != nil {
		h.Logger.Printf("Error rendering FRC page: %v", err)
		http.Error(w, "Failed to render FRC page", http.StatusInternalServerError)
	}
}
