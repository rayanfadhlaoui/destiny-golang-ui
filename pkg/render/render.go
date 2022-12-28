package render

import (
	"destiny-go-ui/pkg/models"
	"fmt"
	"html/template"
	"log"
	"net/http"
)

func AddDefaultData(td *models.TemplateData) *models.TemplateData {
	return td
}

func (t *TemplateInformation) Render(w http.ResponseWriter, templateName string, td *models.TemplateData) {
	if t.useCache {
		log.Println("using cache")
		executeTemplate(w, t.tc[templateName], td)
	} else {
		log.Println("not using cache")
		renderWithoutCache(w, templateName, t.partials, td)
	}
}

func renderWithoutCache(w http.ResponseWriter, t string, partials []string, td *models.TemplateData) {
	var templateSlice []string
	templateSlice = append(templateSlice, fmt.Sprintf("./templates/%s", t))

	for _, x := range partials {
		templateSlice = append(templateSlice, x)
	}

	tmpl, err := template.ParseFiles(templateSlice...)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	executeTemplate(w, tmpl, td)
}

func executeTemplate(w http.ResponseWriter, tmpl *template.Template, td *models.TemplateData) {
	td = AddDefaultData(td)
	if err := tmpl.Execute(w, td); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
