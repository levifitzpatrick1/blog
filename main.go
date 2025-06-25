package main

import (
	"database/sql"
	"log"
	"net/http"
	"os"

	generatedDB "github.com/levifitzpatrick1/blog/internal/database/generated"
	"github.com/levifitzpatrick1/blog/internal/handlers"
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

	site := &Site{
		Queries: queries,
		Logger:  logger,
	}

	mux := http.NewServeMux()

	homeHandler := handlers.NewHomeHandler(site.Queries, site.Logger)
	resumeHandler := handlers.NewResumeHandler(site.Queries, site.Logger)
	blogIndexHandler := handlers.NewBlogIndexHandler(site.Queries, site.Logger)
	postHandler := handlers.NewPostHandler(site.Queries, site.Logger)
	searchHandler := handlers.NewSearchHandler(site.Queries, site.Logger)
	frcHandler := handlers.NewFRCHandler(site.Logger)

	staticDir := "./web/static"
	mux.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir(staticDir))))

	mux.HandleFunc("/{$}", homeHandler.ServeHTTP)
	mux.HandleFunc("/blog", blogIndexHandler.ServeHTTP)
	mux.HandleFunc("/blog/{slug}", postHandler.ServeHTTP)
	mux.HandleFunc("/search", searchHandler.ServeHTTP)
	mux.HandleFunc("/frc", frcHandler.ServeHTTP)
	mux.HandleFunc("/resume", resumeHandler.ServeHTTP)

	logger.Printf("Starting server @ %s", addr)
	err = http.ListenAndServe(addr, mux)
	if err != nil {
		logger.Fatal("ListenAndServer error: ", err)
	}

}
