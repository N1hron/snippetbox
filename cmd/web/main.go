package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	const port = 4000

	mux := http.NewServeMux()
	mux.HandleFunc("/", home)
	mux.HandleFunc("/snippet/view", snippetView)
	mux.HandleFunc("/snippet/create", snippetCreate)

	log.Println("Starting on port", port)
	err := http.ListenAndServe(fmt.Sprintf(":%v", port), mux)
	log.Fatal(err)
}
