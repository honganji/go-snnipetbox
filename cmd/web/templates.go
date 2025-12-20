package main

import (
	"html/template"
	"path/filepath"
	"time"

	"github.com/honganji/go-snippetbox/internal/models"
)

// templateData holds data sent to HTML templates
type templateData struct {
	CurrentYear     int
	Snippet         models.Snippet
	Snippets        []models.Snippet
	Form            any
	Flash           string
	IsAuthenticated bool
	CSRFToken       string
}

func humanDate(t time.Time) string {
	return t.Format("02 Jan 2006 at 15:04")
}

// template.FuncMap holds custom template functions
var functions = template.FuncMap{
	"humanDate": humanDate,
}

func newTemplateCache() (map[string]*template.Template, error) {
	// initialize a new map to act as the cache
	cache := map[string]*template.Template{}
	pages, err := filepath.Glob("./ui/html/pages/*.tmpl.html")
	if err != nil {
		return nil, err
	}
	// range through all the page templates
	// store the resulting templates in the cache map
	for _, page := range pages {
		name := filepath.Base(page)
		// parse the base template
		ts, err := template.New(name).Funcs(functions).ParseFiles("./ui/html/base.tmpl.html")
		if err != nil {
			return nil, err
		}
		// parse any partials
		ts, err = ts.ParseGlob("./ui/html/partials/*.tmpl.html")
		if err != nil {
			return nil, err
		}
		// parse the page template
		ts, err = ts.ParseFiles(page)
		if err != nil {
			return nil, err
		}
		cache[name] = ts
	}
	return cache, nil
}
