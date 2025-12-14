package main

import "github.com/honganji/go-snippetbox/internal/models"

// templateData holds data sent to HTML templates
type templateData struct {
	Snippet models.Snippet
}
