package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
)

type config struct {
	port      string
	staticDir string
}

func main() {
	var cfg config

	flag.StringVar(&cfg.port, "port", "8080", "Server port")
	flag.StringVar(&cfg.staticDir, "static-dir", "./ui/static", "Path to static assets")
	flag.Parse()

	infoLogger := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	errorLogger := log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)

	mux := http.NewServeMux()
	fileServer := http.FileServer(http.Dir(cfg.staticDir))

	mux.Handle("/static/", http.StripPrefix("/static", fileServer))
	mux.HandleFunc("/", home)
	mux.HandleFunc("/snippet/view", snippetView)
	mux.HandleFunc("/snippet/create", snippetCreate)

	server := &http.Server{
		Addr:     fmt.Sprintf(":%v", cfg.port),
		ErrorLog: errorLogger,
		Handler:  mux,
	}

	infoLogger.Println("Starting on port", cfg.port)
	err := server.ListenAndServe()
	errorLogger.Fatal(err)
}
