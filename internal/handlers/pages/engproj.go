package handlers

import (
	"log"
	"net/http"

	"github.com/levifitzpatrick1/blog/web/templates/Pages/atlas"
)

type EngProjHandler struct {
	Logger *log.Logger
}

func NewEngProjHandler(l *log.Logger) *EngProjHandler {
	return &EngProjHandler{Logger: l}
}

func (h *EngProjHandler) Index(w http.ResponseWriter, r *http.Request) {
	err := atlas.Index().Render(r.Context(), w)
	if err != nil {
		h.Logger.Printf("Error rendering engproj index page: %v", err)
		http.Error(w, "Failed to render engproj index", http.StatusInternalServerError)
	}
}

func (h *EngProjHandler) About(w http.ResponseWriter, r *http.Request) {
	err := atlas.About().Render(r.Context(), w)
	if err != nil {
		h.Logger.Printf("Error rendering engproj about page: %v", err)
		http.Error(w, "Failed to render engproj about", http.StatusInternalServerError)
	}
}

func (h *EngProjHandler) Contact(w http.ResponseWriter, r *http.Request) {
	err := atlas.Contact().Render(r.Context(), w)
	if err != nil {
		h.Logger.Printf("Error rendering engproj contact page: %v", err)
		http.Error(w, "Failed to render engproj contact", http.StatusInternalServerError)
	}
}
