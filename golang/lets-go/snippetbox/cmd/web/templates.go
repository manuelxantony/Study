package main

import (
	"html/template"
	"path/filepath"
	"time"

	"blue.study.golang/internal/models"
)

// in book templateData is in main
type templateData struct {
	CurrentYear 	int
	Snippet 		*models.Snippet
	Snippets 		[]*models.Snippet
	Form 			any
	Flash  			string
	IsAuthenticated bool
	CSRFToken 		string
}

func humanDate(t time.Time) string {
	return t.Format("02 Jan 2006 at 15:04")
}

var functions = template.FuncMap{
	"humanDate" : humanDate,
}

func newTemplateCache() (map[string]*template.Template, error) {
	// template do have fields 
	// so even thought it looks like empty
	// it do have some values 
	// and hence no need for make
	cache := map[string]*template.Template{} 

	pages, err := filepath.Glob("./ui/html/pages/*.html")
	if err != nil {
		return nil, err
	}

	for _, page:= range pages {
		// Extract the file name (like 'home.tmpl.html') from the full filepath 
		// and assign it to the name variable.
		name := filepath.Base(page)
		// The template.FuncMap must be registered with the template set before you
		// call the ParseFiles() method. This means we have to use template.New() to
		// create an empty template set, use the Funcs() method to register the
		// template.FuncMap, and then parse the file as normal.
		ts := template.New(name).Funcs(functions)

		ts, err = ts.ParseFiles("./ui/html/base.tmpl.html")
		if err != nil {
			return nil, err
		}

		ts, err = ts.ParseGlob("./ui/html/partials/*.html")
		if err != nil {
			return nil, err
		}

		ts, err = ts.ParseFiles(page)
		if err != nil {
			return nil, err
		}

		cache[name] = ts
	}

	return cache, nil
}

