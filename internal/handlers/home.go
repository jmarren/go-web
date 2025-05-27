package handlers

import (
	"fmt"
	"net/http"

	"github.com/a-h/templ"
	"github.com/jmarren/go-web/internal/components"
	"github.com/jmarren/go-web/internal/db"
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
	users, err := db.Query.GetAllUsers(r.Context())
	if err != nil {
		fmt.Printf("error getting users: %s\n", err)
		return nil
	}
	fmt.Printf("users: %v\n", users)

	return components.Home(&models.HomeData{
		Username: "nickthebrick",
	})
}
