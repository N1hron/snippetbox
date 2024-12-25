package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
)

type config struct {
	port      string
	staticDir string
}

func main() {
	var cfg config

	flag.StringVar(&cfg.port, "port", "8080", "Server port") // Define a new command-line flag
	flag.StringVar(&cfg.staticDir, "static-dir", "./ui/static", "Path to static assets")
	flag.Parse() // Parse the command-line flag

	mux := http.NewServeMux()
	fileServer := http.FileServer(http.Dir(cfg.staticDir))

	mux.Handle("/static/", http.StripPrefix("/static", fileServer))
	mux.HandleFunc("/", home)
	mux.HandleFunc("/snippet/view", snippetView)
	mux.HandleFunc("/snippet/create", snippetCreate)

	log.Println("Starting on port", cfg.port)
	err := http.ListenAndServe(fmt.Sprintf(":%v", cfg.port), mux)
	log.Fatal(err)
}
