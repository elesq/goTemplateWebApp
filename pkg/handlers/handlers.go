package handlers

import (
	"net/http"

	"github.com/elesq/gotemplatewebapp/pkg/config"
	"github.com/elesq/gotemplatewebapp/pkg/models"
	"github.com/elesq/gotemplatewebapp/pkg/render"
)

// Repo is the repository used by the handlers
var Repo *Repository

// Repository defines the respository type
type Repository struct {
	App *config.AppConfig
}

// NewRepo creates a new repository
func NewRepo(a *config.AppConfig) *Repository {
	return &Repository{
		App: a,
	}
}

// NewHandlers sets the repository for the handlers
func NewHandlers(r *Repository) {
	Repo = r
}

// Home function handles requests to the home page
// and responds with a welcome message.
func (rr *Repository) Home(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, "home.page.tmpl", &models.TemplateData{})

}

// About function handles requests to the about page
// and responds with a message including the sum of two numbers.
func (rr *Repository) About(w http.ResponseWriter, r *http.Request) {
	stringMap := make(map[string]string)
	stringMap["test"] = "testdata"

	render.RenderTemplate(w, "about.page.tmpl", &models.TemplateData{
		StringMap: stringMap,
	})
}
