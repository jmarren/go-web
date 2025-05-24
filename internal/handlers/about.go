package handlers

import (
	"net/http"

	"github.com/a-h/templ"
	"github.com/jmarren/go-web/internal/components"
	"github.com/jmarren/go-web/internal/models"
)

type About struct{}

func NewAbout() *Page {
	return &Page{
		Title:     "home",
		Component: &About{},
	}
}

func (h *About) GetComponent(w http.ResponseWriter, r *http.Request) templ.Component {
	return components.About(&models.AboutData{})
}
