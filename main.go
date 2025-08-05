package main

import (
	"bytes"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"sort"
	"strings"

	"github.com/levifitzpatrick1/blog/internal/handlers/functions"
	handlers "github.com/levifitzpatrick1/blog/internal/handlers/pages"
	"github.com/levifitzpatrick1/blog/internal/utils"
	"github.com/levifitzpatrick1/blog/internal/utils/markdown"
	"gopkg.in/yaml.v3"
	_ "modernc.org/sqlite"
)

type Site struct {
	Posts  []utils.Post
	Logger *log.Logger
}

func main() {
	addr := ":4000"

	logger := log.New(os.Stdout, "", log.Ldate|log.Ltime|log.Lshortfile)

	posts, err := loadPosts("Blog Markdowns/Blog", logger)
	if err != nil {
		logger.Fatalf("Failed to load posts: %v", err)
	}

	site := &Site{
		Posts:  posts,
		Logger: logger,
	}

	mux := http.NewServeMux()

	home := handlers.NewHomeHandler(site.Posts, site.Logger)
	//resume := handlers.NewResumeHandler(site.Posts, site.Logger)
	blog := handlers.NewBlogIndexHandler(site.Posts, site.Logger)
	post := handlers.NewPostHandler(site.Posts, site.Logger)
	//frc := handlers.NewFRCHandler(site.Posts, site.Logger)

	search := functions.NewSearchHandler(site.Posts, site.Logger)

	staticDir := "./web/static"
	mux.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir(staticDir))))

	mux.HandleFunc("/{$}", home.ServeHTTP)
	mux.HandleFunc("/blog", blog.ServeHTTP)
	mux.HandleFunc("/blog/{slug}", post.ServeHTTP)
	//mux.HandleFunc("/frc", frc.ServeHTTP)
	//mux.HandleFunc("/resume", resume.ServeHTTP)

	mux.HandleFunc("/search", search.ServeHTTP)

	logger.Printf("Starting server @ %s", addr)
	err = http.ListenAndServe(addr, mux)
	if err != nil {
		logger.Fatal("ListenAndServer error: ", err)
	}

}

func loadPosts(dir string, logger *log.Logger) ([]utils.Post, error) {
	var posts []utils.Post

	err := filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !info.IsDir() && strings.HasSuffix(info.Name(), ".md") {
			content, err := os.ReadFile(path)
			if err != nil {
				return fmt.Errorf("Failed to read md format in %s: %w", path, err)
			}

			parts := bytes.SplitN(content, []byte("---"), 3)
			if len(parts) < 3 {
				return fmt.Errorf("Invalid md format in %s: missing header", path)
			}

			var tmpPost utils.TempPost
			var post utils.Post
			if err := yaml.Unmarshal(parts[1], &tmpPost); err != nil {
				return fmt.Errorf("Failed to parse header in %s: %w", path, err)
			}
			post.Title = tmpPost.Title
			post.Blurb = tmpPost.Blurb
			post.Slug = tmpPost.Slug
			post.CreateDate = tmpPost.CreateDate
			post.ModifyDate = tmpPost.ModifyDate
			post.Content = string(parts[2])

			post.GetTags(tmpPost.Tags)

			rendered, err := markdown.RenderHTML(post.Content)
			if err != nil {
				return fmt.Errorf("Failed to render html in %s: %w", path, err)
			}
			post.HTMLContent = template.HTML(rendered)

			posts = append(posts, post)
			logger.Printf("Loaded post: %s", post.Title)
		}
		return nil
	})
	if err != nil {
		return nil, err
	}
	sort.Slice(posts, func(i, j int) bool {
		return posts[i].ModifyDate.After(posts[j].ModifyDate)
	})

	return posts, nil
}
