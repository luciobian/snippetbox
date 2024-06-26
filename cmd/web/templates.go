package main

import (
	"html/template"
	"lucio/snippetbox/pkg/forms"
	"lucio/snippetbox/pkg/models"
	"path/filepath"
	"time"
)

type templateData struct {
	AuthenticatedUser int
	CurrentYear       int
	Form              *forms.Form
	Snippet           *models.Snippet
	Flash             string
	Snippets          []*models.Snippet
}

func humanDate(t time.Time) string {
	return t.Format("02 Jan 2006 15:04:05")
}

var functions = template.FuncMap{
	"humanDate": humanDate,
}

func NewTemplateCache(dir string) (map[string]*template.Template, error) {
	cache := map[string]*template.Template{}

	pages, err := filepath.Glob(filepath.Join(dir, "*.page.gohtml"))
	if err != nil {
		return nil, err
	}

	for _, page := range pages {
		name := filepath.Base(page)

		ts, err := template.New(name).Funcs(functions).ParseFiles(page)
		if err != nil {
			return nil, err
		}

		ts, err = ts.ParseGlob(filepath.Join(dir, "*.layout.gohtml"))
		if err != nil {
			return nil, err
		}

		ts, err = ts.ParseGlob(filepath.Join(dir, "*.partial.gohtml"))
		if err != nil {
			return nil, err
		}

		cache[name] = ts
	}
	return cache, nil
}
