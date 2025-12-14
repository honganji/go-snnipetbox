package main

import (
	"errors"
	"fmt"
	"html/template"
	"log/slog"
	"net/http"
	"strconv"

	"github.com/honganji/go-snippetbox/internal/models"
)

// return the latest snippets
// for now, this just writes a placeholder message
func (app *application) home(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Server", "Go")
	snippets, err := app.snippets.Latest()
	if err != nil {
		app.serverError(w, r, err)
		return
	}
	for _, snippet := range snippets {
		fmt.Fprintf(w, "%+v\n", snippet)
	}
	// files := []string{
	// 	"./ui/html/base.tmpl.html",
	// 	"./ui/html/pages/home.tmpl.html",
	// 	"./ui/html/partials/nav.tmpl.html",
	// }
	// ts, err := template.ParseFiles(files...)
	// if err != nil {
	// 	app.serverError(w, r, err)
	// 	return
	// }
	// err = ts.ExecuteTemplate(w, "base", nil)
	// if err != nil {
	// 	app.logger.Error(err.Error(), slog.String("method", r.Method), slog.String("url", r.URL.String()))
	// 	app.serverError(w, r, err)
	// }
}

// returns the snippet for a given ID
func (app *application) snippetView(w http.ResponseWriter, r *http.Request) {
	// get the snippet by its id
	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil || id < 1 {
		http.NotFound(w, r)
		return
	}
	snippet, err := app.snippets.Get(id)
	if err != nil {
		if errors.Is(err, models.ErrNoRecord) {
			http.NotFound(w, r)
		} else {
			app.serverError(w, r, err)
		}
		return
	}

	// load html files and parse templates
	files := []string{
		"./ui/html/base.tmpl.html",
		"./ui/html/pages/view.tmpl.html",
		"./ui/html/partials/nav.tmpl.html",
	}
	ts, err := template.ParseFiles(files...)
	if err != nil {
		app.serverError(w, r, err)
		return
	}
	// with the snippet data, write the html to the response writer
	data := templateData{
		Snippet: snippet,
	}
	err = ts.ExecuteTemplate(w, "base", data)
	if err != nil {
		app.logger.Error(err.Error(), slog.String("method", r.Method), slog.String("url", r.URL.String()))
		app.serverError(w, r, err)
	}
}

// return a form for creating a new snippet
func (app *application) snippetCreate(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Display a form for creating a new snippet..."))
}

// process the form submission
func (app *application) snippetCreatePost(w http.ResponseWriter, r *http.Request) {
	title := "0 snail"
	content := "O snail\nClimb Mount Fuji,\nBut slowly, slowly!\n\n– Kobayashi Issa"
	expires := 7
	id, err := app.snippets.Insert(title, content, expires)
	if err != nil {
		app.serverError(w, r, err)
		return
	}
	http.Redirect(w, r, fmt.Sprintf("/snippet/view/%d", id), http.StatusSeeOther)
}
