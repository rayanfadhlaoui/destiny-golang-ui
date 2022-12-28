package handlers

import (
	"github.com/rayanfadhlaoui/destiny-golang-ui/pkg/models"
	"github.com/rayanfadhlaoui/destiny-golang-ui/pkg/render"
	"net/http"
)

type Repository struct {
	templateInformation *render.TemplateInformation
}

func NewRepository(useCache bool) *Repository {
	return &Repository{
		templateInformation: render.NewTemplateInformation(useCache),
	}
}

func (re *Repository) Home(w http.ResponseWriter, r *http.Request) {
	td := &models.TemplateData{StringMap: map[string]string{"name": "rayan"}}
	re.templateInformation.Render(w, "home.page.gohtml", td)
}

func (re *Repository) About(w http.ResponseWriter, r *http.Request) {
	re.templateInformation.Render(w, "about.page.gohtml", &models.TemplateData{})
}
