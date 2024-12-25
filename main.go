package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
)

func home(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}

	w.Write([]byte("Hello from Snippetbox"))
}

func snippetView(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil || id < 1 {
		http.NotFound(w, r)
		return
	}

	fmt.Fprintf(w, "A specific snippet with id %d", id)
}

func snippetCreate(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		//                      POST
		w.Header().Add("Allow", http.MethodPost) // Must be sent if the server responds with a 405 Method Not Allowed status code

		// w.WriteHeader(405) // Should be always before Write()
		// w.Write([]byte("Method Not Allowed"))

		//                                  405
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}

	w.Write([]byte("Create new snippet"))
}

func main() {
	const port = 4000

	// It is generally a good idea to avoid using DefaultServeMux
	// and functions like http.Handle() and http.HandleFunc()
	mux := http.NewServeMux()                        // Router (or servemux in Go terminology)
	mux.HandleFunc("/", home)                        // Subtree path (match a single slash, followed by anything)
	mux.HandleFunc("/snippet/view", snippetView)     // Fixed path (don’t end with a trailing slash)
	mux.HandleFunc("/snippet/create", snippetCreate) // Fixed path

	log.Println("Starting on port", port)
	err := http.ListenAndServe(fmt.Sprintf(":%v", port), mux) // addr should be in the format "host:port"
	log.Fatal(err)
}
