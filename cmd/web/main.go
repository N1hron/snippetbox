package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
)

type application struct {
	infoLogger  *log.Logger
	errorLogger *log.Logger
}

func main() {
	port := flag.String("port", "8080", "Server port")
	staticDir := flag.String("static-dir", "./ui/static", "Path to static assets")
	flag.Parse()

	infoLogger := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	errorLogger := log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)

	app := &application{
		infoLogger:  infoLogger,
		errorLogger: errorLogger,
	}

	mux := http.NewServeMux()
	fileServer := http.FileServer(http.Dir(*staticDir))

	mux.Handle("/static/", http.StripPrefix("/static", fileServer))
	mux.HandleFunc("/", app.home)
	mux.HandleFunc("/snippet/view", app.snippetView)
	mux.HandleFunc("/snippet/create", app.snippetCreate)

	server := &http.Server{
		Addr:     fmt.Sprintf(":%v", *port),
		ErrorLog: errorLogger,
		Handler:  mux,
	}

	infoLogger.Println("Starting on port", *port)
	err := server.ListenAndServe()
	errorLogger.Fatal(err)
}
