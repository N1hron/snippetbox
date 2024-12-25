package main

import (
	"fmt"
	"log"
	"net/http"
)

func home(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello from Snippetbox"))
}

func main() {
	const port = 4000

	mux := http.NewServeMux() // Router (or servemux in Go terminology)
	mux.HandleFunc("/", home) // Go's servemux treats the URL pattern "/" like a catch-all

	log.Println("Starting on port", port)
	err := http.ListenAndServe(fmt.Sprintf(":%v", port), mux) // addr should be in the format "host:port"
	log.Fatal(err)
}
