package main

import (
	"log"
	"net/http"
	"os"

	"github.com/levifitzpatrick1/blog/internal/handlers/functions"
	handlers "github.com/levifitzpatrick1/blog/internal/handlers/pages"
	"github.com/levifitzpatrick1/blog/internal/utils/markdown"
)

type Site struct {
	Store  *markdown.PostStore
	Logger *log.Logger
}

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "4000"
	}
	addr := ":" + port

	logger := log.New(os.Stdout, "", log.Ldate|log.Ltime|log.Lshortfile)

	store := markdown.NewPostStore("BlogMarkdowns", logger)

	site := &Site{
		Store:  store,
		Logger: logger,
	}

	mux := http.NewServeMux()

	home := handlers.NewHomeHandler(site.Store, site.Logger)
	blog := handlers.NewBlogIndexHandler(site.Store, site.Logger)
	post := handlers.NewPostHandler(site.Store, site.Logger)
	frc := handlers.NewFRCHandler(site.Logger)

	search := functions.NewSearchHandler(site.Store, site.Logger)

	staticDir := "./web/static"
	mux.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir(staticDir))))

	mux.HandleFunc("/{$}", home.ServeHTTP)
	mux.HandleFunc("/blog", blog.ServeHTTP)
	mux.HandleFunc("/blog/{slug}", post.ServeHTTP)
	mux.HandleFunc("/frc", frc.ServeHTTP)

	mux.HandleFunc("/search", search.ServeHTTP)

	logger.Printf("Starting server @ %s", addr)
	err := http.ListenAndServe(addr, mux)
	if err != nil {
		logger.Fatal("ListenAndServe error: ", err)
	}

}
