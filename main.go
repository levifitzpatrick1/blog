package main

import (
	"database/sql"
	"log"
	"net/http"
	"os"

	generatedDB "github.com/levifitzpatrick1/blog/internal/database/generated"
	"github.com/levifitzpatrick1/blog/internal/handlers"
	webTemplates "github.com/levifitzpatrick1/blog/web/templates"
	_ "modernc.org/sqlite"
)

type Site struct {
	Queries *generatedDB.Queries
	Logger  *log.Logger
}

func main() {
	addr := ":4000"
	dbPath := "./blog.db"

	logger := log.New(os.Stdout, "", log.Ldate|log.Ltime|log.Lshortfile)

	sqlDB, err := sql.Open("sqlite", dbPath)
	if err != nil {
		logger.Fatal("Cannot connect to database: ", err)
	}
	defer sqlDB.Close()

	queries := generatedDB.New(sqlDB)

	webTemplates.LoadTemplates("web/templates")

	site := &Site{
		Queries: queries,
		Logger:  logger,
	}

	mux := http.NewServeMux()

	homeHandler := handlers.NewHomeHandler(site.Queries, site.Logger)
	postHandler := handlers.NewPostHandler(site.Queries, site.Logger)
	searchHandler := handlers.NewSearchHandler(site.Queries, site.Logger)

	staticDir := "./web/static"
	if _, statErr := os.Stat(staticDir); os.IsNotExist(statErr) {
		logger.Println("Warning: Static directory './web/static' not found. CSS might not load.")
	} else {
		logger.Printf("Serving static files from: %s", staticDir)
	}
	logger.Printf("Serving static files from: %s", staticDir)
	mux.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir(staticDir))))

	mux.HandleFunc("GET /{$}", homeHandler.ServeHTTP)
	mux.HandleFunc("GET /search", searchHandler.ServeHTTP)
	mux.HandleFunc("GET /blog/{slug}", postHandler.ServeHTTP)
	mux.HandleFunc("/request", func(w http.ResponseWriter, r *http.Request) {
		println("got request")
	})

	logger.Printf("Starting server @ %s", addr)
	err = http.ListenAndServe(addr, mux)
	if err != nil {
		logger.Fatal("ListenAndServer error: ", err)
	}

}
