package handlers

import (
	"net/http"

	"github.com/a-h/templ"
	"github.com/jmarren/go-web/internal/components"
	"github.com/jmarren/go-web/internal/models"
)

type Home struct{}

func NewHome() *Page {
	return &Page{
		Title:     "home",
		Component: &Home{},
	}
}

func (h *Home) GetComponent(w http.ResponseWriter, r *http.Request) templ.Component {
	return components.Home(&models.HomeData{
		Username: "nickthebrick",
	})
}
