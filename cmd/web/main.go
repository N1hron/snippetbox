package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
)

type application struct {
	port        *string
	staticDir   *string
	infoLogger  *log.Logger
	errorLogger *log.Logger
}

func main() {
	app := &application{
		port:        flag.String("port", "8080", "Server port"),
		staticDir:   flag.String("static-dir", "./ui/static", "Path to static assets"),
		infoLogger:  log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime),
		errorLogger: log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile),
	}

	flag.Parse()

	server := &http.Server{
		Addr:     fmt.Sprintf(":%v", *app.port),
		ErrorLog: app.errorLogger,
		Handler:  app.routes(),
	}

	app.infoLogger.Println("Starting on port", *app.port)
	err := server.ListenAndServe()
	app.errorLogger.Fatal(err)
}
