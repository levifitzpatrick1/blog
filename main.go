package main

import (
	"log"
	"net/http"
	"os"

	"github.com/levifitzpatrick1/blog/internal/handlers/functions"
	handlers "github.com/levifitzpatrick1/blog/internal/handlers/pages"
	"github.com/levifitzpatrick1/blog/internal/utils"
	"github.com/levifitzpatrick1/blog/internal/utils/markdown"
)

type Site struct {
	Posts  []utils.Post
	Logger *log.Logger
}

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "4000"
	}
	addr := ":" + port

	logger := log.New(os.Stdout, "", log.Ldate|log.Ltime|log.Lshortfile)

	posts, err := markdown.LoadPosts("BlogMarkdowns/Blog", logger)
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
	frc := handlers.NewFRCHandler(site.Posts, site.Logger)

	search := functions.NewSearchHandler(site.Posts, site.Logger)

	staticDir := "./web/static"
	mux.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir(staticDir))))

	mux.HandleFunc("/{$}", home.ServeHTTP)
	mux.HandleFunc("/blog", blog.ServeHTTP)
	mux.HandleFunc("/blog/{slug}", post.ServeHTTP)
	mux.HandleFunc("/frc", frc.ServeHTTP)

	mux.HandleFunc("/search", search.ServeHTTP)

	logger.Printf("Starting server @ %s", addr)
	err = http.ListenAndServe(addr, mux)
	if err != nil {
		logger.Fatal("ListenAndServer error: ", err)
	}

}
