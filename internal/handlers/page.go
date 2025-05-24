package handlers

import (
	"net/http"

	"github.com/a-h/templ"
	"github.com/jmarren/go-web/internal/components"
	"github.com/jmarren/go-web/internal/models"
)

type ComponentHandler interface {
	GetComponent(w http.ResponseWriter, r *http.Request) templ.Component
}

type Page struct {
	Title     string
	Component ComponentHandler
}

func (p *Page) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// get component for page
	pageComponent := p.Component.GetComponent(w, r)

	// if it is an HX-request Render the component only
	if r.Header.Get("HX-Request") == "true" {
		pageComponent.Render(r.Context(), w)
		return
	}

	// render in root
	components.Root(&models.PageData{
		Title: p.Title,
		Page:  pageComponent,
	}).Render(r.Context(), w)
}
