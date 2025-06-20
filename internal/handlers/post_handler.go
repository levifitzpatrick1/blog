package handlers

import (
	"database/sql"
	"html/template"
	"log"
	"net/http"
	"strings"

	generatedDB "github.com/levifitzpatrick1/blog/internal/database/generated"
	"github.com/levifitzpatrick1/blog/internal/markdown"
	webTemplates "github.com/levifitzpatrick1/blog/web/templates"
)

type PostHandler struct {
	Queries *generatedDB.Queries
	Logger  *log.Logger
}

func NewPostHandler(q *generatedDB.Queries, l *log.Logger) *PostHandler {
	return &PostHandler{Queries: q, Logger: l}
}

type PostPageData struct {
	Title          string
	HTMLContent    template.HTML
	PublishDate    string
	UpdateDate     string
	ShowUpdateDate bool
}

func (h *PostHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	slug := strings.TrimPrefix(r.URL.Path, "/blog/")
	if slug == "" {
		http.NotFound(w, r)
		return
	}

	post, err := h.Queries.GetPostBySlug(ctx, slug)
	if err != nil {
		if err == sql.ErrNoRows {
			log.Printf("Post with slug '%s' not found.", slug)
			http.NotFound(w, r)
			return
		}
		log.Printf("Error fetching post with slug '%s': %v", slug, err)
		http.Error(w, "Error fetching post", http.StatusInternalServerError)
		return
	}

	renderedHTMLContent, err := markdown.RenderHTML(post.Content)
	if err != nil {
		log.Printf("Error rendering markdown as HTML: %v", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}

	publishDate := post.Published.Time.Format("01/02/2006")
	updateDate := post.Updated.Time.Format("01/02/2006")
	showUpdate := publishDate != updateDate

	pageData := PostPageData{
		Title:          post.Title,
		HTMLContent:    template.HTML(renderedHTMLContent),
		PublishDate:    publishDate,
		UpdateDate:     updateDate,
		ShowUpdateDate: showUpdate,
	}

	webTemplates.RenderTemplate(w, "post_page.html", pageData)
}
