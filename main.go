package main

import (
	"database/sql"
	"log"
	"net/http"
	"os"

	generatedDB "github.com/levifitzpatrick1/blog/internal/database/generated"
	"github.com/levifitzpatrick1/blog/internal/handlers/functions"
	handlers "github.com/levifitzpatrick1/blog/internal/handlers/pages"
	_ "modernc.org/sqlite"
)

type Site struct {
	Queries *generatedDB.Queries
	Logger  *log.Logger
}

func main() {
	addr := ":4000"
	dbPath := "./site.db"

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

	home := handlers.NewHomeHandler(site.Queries, site.Logger)
	resume := handlers.NewResumeHandler(site.Queries, site.Logger)
	blog := handlers.NewBlogIndexHandler(site.Queries, site.Logger)
	post := handlers.NewPostHandler(site.Queries, site.Logger)
	frc := handlers.NewFRCHandler(site.Queries, site.Logger)

	search := functions.NewSearchHandler(site.Queries, site.Logger)

	staticDir := "./web/static"
	mux.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir(staticDir))))

	mux.HandleFunc("/{$}", home.ServeHTTP)
	mux.HandleFunc("/blog", blog.ServeHTTP)
	mux.HandleFunc("/blog/{slug}", post.ServeHTTP)
	mux.HandleFunc("/frc", frc.ServeHTTP)
	mux.HandleFunc("/resume", resume.ServeHTTP)

	mux.HandleFunc("/search", search.ServeHTTP)

	logger.Printf("Starting server @ %s", addr)
	err = http.ListenAndServe(addr, mux)
	if err != nil {
		logger.Fatal("ListenAndServer error: ", err)
	}

}
