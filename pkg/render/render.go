package render

import (
	"bytes"
	"html/template"
	"log"
	"net/http"
	"path/filepath"

	"github.com/elesq/gotemplatewebapp/pkg/config"
	"github.com/elesq/gotemplatewebapp/pkg/models"
)

var app *config.AppConfig

// NewTemplates sets the package wide config for the templates
// that will be used in the application
func NewTemplates(a *config.AppConfig) {
	app = a
}

// AddDefaultData is a utility that allows us to add any default data
// that needs to be available to every page. This should be called in
// advance of executing the template in the RenderTemplate function
func AddDefaultData(td *models.TemplateData) *models.TemplateData {
	return td
}

// RenderTemplate renders templates using the html/template
func RenderTemplate(w http.ResponseWriter, tmpl string, data *models.TemplateData) {
	var tc map[string]*template.Template

	if app.UseCache {
		// get the template cache from the app config
		tc = app.TemplateCache
	} else {
		// create the template cache
		tc, _ = CreateTemplateCache()
	}

	// get the requested template from cache
	t, ok := tc[tmpl]
	if !ok {
		log.Fatal("could not retrieve template from templates cache")
	}

	buf := new(bytes.Buffer)
	err := t.Execute(buf, data)
	if err != nil {
		log.Println(err)
	}

	// render the template
	_, err = buf.WriteTo(w)
	if err != nil {
		log.Println(err)
	}
}

// CreateTemplateCache creates a templates cache to be used by the
// application. It will first take the pages matching the '.page.tmpl'
// pattern and creste the cache. For each of the pages in the template
// cache the function will add the '.layout.tmpl' to be applied to
// the page. Any error encountered in the generation will return the
// pagesCache and trigger an error that the template cache could not
// be successfully created.
func CreateTemplateCache() (map[string]*template.Template, error) {
	pagesCache := map[string]*template.Template{}

	// get all files *.page.tmpl
	pages, err := filepath.Glob("./templates/*.page.tmpl")
	if err != nil {
		return pagesCache, err
	}

	// range files
	for _, page := range pages {
		name := filepath.Base(page)
		ts, err := template.New(name).ParseFiles(page)
		if err != nil {
			return pagesCache, err
		}

		matches, err := filepath.Glob("./templates/*.layout.tmpl")
		if err != nil {
			return pagesCache, err
		}

		if len(matches) > 0 {
			ts, err = ts.ParseGlob("./templates/*.layout.tmpl")
			if err != nil {
				return pagesCache, err
			}
		}

		pagesCache[name] = ts
	}

	return pagesCache, nil
}
