package handlers

import (
	"log"
	"net/http"

	generatedDB "github.com/levifitzpatrick1/blog/internal/database/generated"
	webTemplates "github.com/levifitzpatrick1/blog/web/templates"
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
	s := r.URL.Query().Get("sort")
	var posts []generatedDB.ListPostsRow
	var err error

	switch s {
	case "oldest":
		tp, te := h.Queries.ListPostsOldestFirst(ctx)
		if te == nil {
			posts = make([]generatedDB.ListPostsRow, len(tp))
			for i, p := range tp {
				posts[i] = generatedDB.ListPostsRow{
					ID:        p.ID,
					Title:     p.Title,
					Slug:      p.Slug,
					Published: p.Published,
					Excerpt:   p.Excerpt,
				}
			}
		}
		err = te
	default:
		posts, err = h.Queries.ListPosts(ctx)
	}
	if err != nil {
		h.Logger.Print("Error fetching posts: ", err)
		http.Error(w, "Whoops! I Couldn't Find The Posts!", http.StatusInternalServerError)
	}

	postData := struct {
		Posts       []generatedDB.ListPostsRow
		CurrentSort string
	}{
		Posts:       posts,
		CurrentSort: s,
	}

	if r.Header.Get("HX-Request") == "true" && r.Header.Get("HX-Target") == "#posts-list-container" {
		webTemplates.RenderTemplate(w, "post_list_items.html", postData)
		return
	}

	pageData := struct {
		CurrentSort string
		PostsData   any
	}{
		CurrentSort: s,
		PostsData:   postData,
	}
	webTemplates.RenderTemplate(w, "home.html", pageData)
}
