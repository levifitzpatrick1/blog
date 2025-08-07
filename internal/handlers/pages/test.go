package handlers

import (
	"log"
	"net/http"

	"github.com/levifitzpatrick1/blog/internal/utils"
	"github.com/levifitzpatrick1/blog/web/templates/Pages/test"
)

type TestHandler struct {
	Posts  []utils.Post
	Logger *log.Logger
}

func NewTestHandler(posts []utils.Post, l *log.Logger) *TestHandler {
	return &TestHandler{Posts: posts, Logger: l}
}

func (h *TestHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	component := test.TestPage()
	err := component.Render(r.Context(), w)
	if err != nil {
		h.Logger.Print("Error rendering template: ", err)
		http.Error(w, "Failed to render page.", http.StatusInternalServerError)
	}
}
