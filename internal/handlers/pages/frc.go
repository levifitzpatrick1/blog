package handlers

import (
	"log"
	"net/http"

	"github.com/levifitzpatrick1/levifitzpatrick.page/web/templates/Pages/frc"
)

type FRCHandler struct {
	Logger *log.Logger
}

func NewFRCHandler(l *log.Logger) *FRCHandler {
	return &FRCHandler{Logger: l}
}

func (h *FRCHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	err := frc.Frc().Render(r.Context(), w)
	if err != nil {
		h.Logger.Printf("Error rendering blog index page: %v", err)
		http.Error(w, "Failed to render blog index", http.StatusInternalServerError)
	}
}
