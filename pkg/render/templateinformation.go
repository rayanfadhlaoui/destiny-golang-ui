package render

import (
	"html/template"
	"log"
	"path/filepath"
)

type TemplateInformation struct {
	partials []string
	useCache bool
	tc       map[string]*template.Template
}

func NewTemplateInformation(useCache bool) *TemplateInformation {
	log.Println("init")
	var tc map[string]*template.Template
	partials := createPartials()
	if useCache {
		tc = populateCache(partials)
	}
	return &TemplateInformation{
		useCache: useCache,
		partials: partials,
		tc:       tc,
	}
}

func populateCache(partials []string) map[string]*template.Template {
	tc := make(map[string]*template.Template)
	matches, err := filepath.Glob("./templates/*.page.gohtml")
	if err != nil {
		log.Fatalf(err.Error())
	}

	for _, page := range matches {
		name := filepath.Base(page)
		ts, err := template.New(name).ParseFiles(page)
		if err != nil {
			log.Fatalf(err.Error())
		}

		_, err = ts.ParseFiles(partials...)
		if err != nil {
			log.Fatalf(err.Error())
		}

		tc[name] = ts
	}
	return tc
}

func createPartials() []string {
	layoutMatches, err := filepath.Glob("./templates/*.layout.gohtml")
	if err != nil {
		log.Fatal(err)
	}

	partialMatches, err := filepath.Glob("./templates/*.partial.gohtml")

	if err != nil {
		log.Fatal(err)
	}

	return append(layoutMatches, partialMatches...)
}
