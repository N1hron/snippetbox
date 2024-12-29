package main

import (
	"database/sql"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"github.com/n1hron/snippetbox/internal/models"
)

type application struct {
	port        *string
	staticDir   *string
	infoLogger  *log.Logger
	errorLogger *log.Logger
	snippets    *models.SnippetModel
}

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	port := flag.String("port", "8080", "Server port")
	staticDir := flag.String("static-dir", "./ui/static", "Path to static assets")
	dsn := flag.String("dsn", os.Getenv("DSN"), "PostgreSQL data source name")

	flag.Parse()

	infoLogger := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	errorLogger := log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)

	db, err := openDB(*dsn)
	if err != nil {
		errorLogger.Fatal(err)
	}
	defer db.Close()

	app := &application{
		port:        port,
		staticDir:   staticDir,
		infoLogger:  infoLogger,
		errorLogger: errorLogger,
		snippets:    &models.SnippetModel{DB: db},
	}

	server := &http.Server{
		Addr:     fmt.Sprintf(":%v", *app.port),
		ErrorLog: app.errorLogger,
		Handler:  app.routes(),
	}

	app.infoLogger.Println("Starting on port", *app.port)
	err = server.ListenAndServe()
	app.errorLogger.Fatal(err)
}

func openDB(dsn string) (*sql.DB, error) {
	db, err := sql.Open("postgres", dsn)
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	return db, nil
}
