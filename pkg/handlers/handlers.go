package handlers

import (
	"net/http"

	"github.com/rtmelsov/mansProducts/pkg/config"
	"github.com/rtmelsov/mansProducts/pkg/models"
	"github.com/rtmelsov/mansProducts/pkg/render"
)

var Repo *Repository

type Repository struct {
	App *config.AppConfig
}

func GetNewRepository(app *config.AppConfig) *Repository {
	return &Repository{
		App: app,
	}
}

func NewHandler(r *Repository) {
	Repo = r
}

// Home is the home page handler
func (m *Repository) Home(w http.ResponseWriter, r *http.Request) {
	remoteIp := r.RemoteAddr
	m.App.Session.Put(r.Context(), "remote_ip", remoteIp)
	render.RenderTemplate(w, "home.page.tmpl", &models.TemplateData{})
}

// About is the about page handler
func (m *Repository) About(w http.ResponseWriter, r *http.Request) {
	testMap := make(map[string]string)
	testMap["test"] = "Hello World"
	remoteIp := m.App.Session.GetString(r.Context(), "remote_ip")
	testMap["remote_ip"] = remoteIp
	render.RenderTemplate(w, "about.page.tmpl", &models.TemplateData{
		StringMap: testMap,
	})
}
