package main

import (
	"database/sql"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"

	_ "github.com/lib/pq"
)

type application struct {
	staticDir   *string
	infoLogger  *log.Logger
	errorLogger *log.Logger
}

func main() {
	port := flag.String("port", "8080", "Server port")
	staticDir := flag.String("static-dir", "./ui/static", "Path to static assets")
	dsn := flag.String("dsn", "", "PostgreSQL data source name")

	app := &application{
		staticDir:   staticDir,
		infoLogger:  log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime),
		errorLogger: log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile),
	}

	flag.Parse()

	db, err := openDB(*dsn)
	if err != nil {
		app.errorLogger.Fatal(err)
	}
	defer db.Close()

	server := &http.Server{
		Addr:     fmt.Sprintf(":%v", *port),
		ErrorLog: app.errorLogger,
		Handler:  app.routes(),
	}

	app.infoLogger.Println("Starting on port", *port)
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
