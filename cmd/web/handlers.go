package main

import (
	"errors"
	"fmt"
	"net/http"
	"strconv"

	"github.com/honganji/go-snippetbox/internal/models"
)

// renders the home page with the latest snippets
func (app *application) home(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Server", "Go")
	// get the latest snippets from the database
	snippets, err := app.snippets.Latest()
	if err != nil {
		app.serverError(w, r, err)
		return
	}
	// render the home page template, passing in the latest snippets
	app.render(w, r, http.StatusOK, "home.tmpl.html", templateData{Snippets: snippets})
}

// renders the snippet view page for a given ID
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

	// render the snippet view page template, passing in the snippet data
	app.render(w, r, http.StatusOK, "view.tmpl.html", templateData{Snippet: snippet})
}

// renders the snippet creation form
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
