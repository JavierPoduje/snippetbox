package main

import "snippetbox.javier.net/internal/models"

type templateData struct {
	Snippet  models.Snippet
	Snippets []models.Snippet
}
