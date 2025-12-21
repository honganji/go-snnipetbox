package main

import (
	"html/template"
	"io/fs"
	"path/filepath"
	"time"

	"github.com/honganji/go-snippetbox/internal/models"
	"github.com/honganji/go-snippetbox/ui"
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
	pages, err := fs.Glob(ui.Files, "html/pages/*.tmpl.html")
	if err != nil {
		return nil, err
	}
	// range through all the page templates
	// store the resulting templates in the cache map
	for _, page := range pages {
		name := filepath.Base(page)
		// create slice of all the files to be parsed for the template
		patterns := []string{
			"html/base.tmpl.html",
			"html/partials/*.tmpl.html",
			page,
		}
		ts, err := template.New(name).Funcs(functions).ParseFS(ui.Files, patterns...)
		if err != nil {
			return nil, err
		}
		cache[name] = ts
	}
	return cache, nil
}
