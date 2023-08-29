package main

import "01/internal/models"

type templateData struct {
	Snippet *models.Snippet
	Snippets []*models.Snippet
}
