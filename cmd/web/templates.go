package main

import (
	"path/filepath"
	"text/template"
	"time"

	"github.com/n1hron/snippetbox/internal/models"
)

type templateData struct {
	CurrentYear int
	Snippet     *models.Snippet
	Snippets    []*models.Snippet
}

func (app *application) newTemplateData() *templateData {
	return &templateData{
		CurrentYear: time.Now().Year(),
	}
}

func newTemplateCache() (map[string]*template.Template, error) {
	cache := map[string]*template.Template{}

	pages, err := filepath.Glob("./ui/html/pages/*html")
	if err != nil {
		return nil, err
	}

	for _, page := range pages {
		ts, err := template.ParseFiles("./ui/html/base.html")
		if err != nil {
			return nil, err
		}

		ts, err = ts.ParseGlob("./ui/html/partials/*html")
		if err != nil {
			return nil, err
		}

		ts, err = ts.ParseFiles(page)
		if err != nil {
			return nil, err
		}

		name := filepath.Base(page)
		cache[name] = ts
	}

	return cache, nil
}
