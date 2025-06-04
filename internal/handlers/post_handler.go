package handlers

import (
	"database/sql"
	"html/template"
	"log"
	"net/http"
	"strings"

	"github.com/gomarkdown/markdown"
	"github.com/gomarkdown/markdown/html"
	"github.com/gomarkdown/markdown/parser"
	generatedDB "github.com/levifitzpatrick1/blog/internal/database/generated"
	webTemplates "github.com/levifitzpatrick1/blog/web/templates"
)

type PostHandler struct {
	Queries *generatedDB.Queries
	Logger  *log.Logger
}

func NewPostHandler(q *generatedDB.Queries, l *log.Logger) *PostHandler {
	return &PostHandler{Queries: q, Logger: l}
}

func markdownToHTML(md []byte) string {
	extensions := parser.CommonExtensions | parser.AutoHeadingIDs | parser.NoEmptyLineBeforeBlock
	p := parser.NewWithExtensions(extensions)
	doc := p.Parse(md)

	htmlFlags := html.CommonFlags | html.HrefTargetBlank
	opts := html.RendererOptions{Flags: htmlFlags}
	renderer := html.NewRenderer(opts)

	return string(markdown.Render(doc, renderer))
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

	extensions := parser.CommonExtensions | parser.AutoHeadingIDs
	p := parser.NewWithExtensions(extensions)
	doc := p.Parse([]byte(post.Content))

	htmlFlags := html.CommonFlags | html.HrefTargetBlank
	opts := html.RendererOptions{Flags: htmlFlags}
	renderer := html.NewRenderer(opts)
	renderedHTMLContent := markdown.Render(doc, renderer)

	pageData := struct {
		Post                generatedDB.Post
		RenderedHTMLContent template.HTML
	}{
		Post:                post,
		RenderedHTMLContent: template.HTML(renderedHTMLContent),
	}

	webTemplates.RenderTemplate(w, "post_page.html", pageData)
}
