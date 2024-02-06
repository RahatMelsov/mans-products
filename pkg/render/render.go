package render

import (
	"bytes"
	"fmt"
	"log"
	"net/http"
	"path/filepath"
	"text/template"

	"github.com/rtmelsov/mansProducts/pkg/config"
	"github.com/rtmelsov/mansProducts/pkg/models"
)

var functions = template.FuncMap{}
var app *config.AppConfig

// NewTemplates gets new templates from the application config
func NewTemplates(a *config.AppConfig) {
	app = a
}

func GetDefaultTemplateData(td *models.TemplateData) *models.TemplateData {
	return td
}

// RenderTemplate render templates using html/templates
func RenderTemplate(w http.ResponseWriter, path string, td *models.TemplateData) {
	var tc map[string]*template.Template
	if app.UseCache {
		tc = app.TemplateCache
	} else {
		var err error
		tc, err = CreateTemplateCache()
		if err != nil {
			log.Fatal(err)
		}
	}
	t, ok := tc[path]
	if !ok {
		fmt.Println("Parsing template error")
		return
	}
	buf := new(bytes.Buffer)
	td = GetDefaultTemplateData(td)
	_ = t.Execute(buf, td)

	_, err := buf.WriteTo(w)

	if err != nil {
		fmt.Println("Error write tempalate to browser", err)
	}
}

// CreateTemplateCache creates a template cache as a map
func CreateTemplateCache() (map[string]*template.Template, error) {
	myCache := map[string]*template.Template{}

	pages, err := filepath.Glob("./templates/*.page.tmpl")
	if err != nil {
		return myCache, err
	}

	for _, page := range pages {
		name := filepath.Base(page)
		ts, err := template.New(name).Funcs(functions).ParseFiles(page)
		if err != nil {
			return myCache, err
		}
		matches, err := filepath.Glob("./templates/*.layout.tmpl")
		if err != nil {
			return myCache, err
		}
		if len(matches) > 0 {
			ts, err = ts.ParseGlob("./templates/*.layout.tmpl")
			if err != nil {
				return myCache, err
			}
		}

		myCache[name] = ts
	}
	return myCache, nil
}
