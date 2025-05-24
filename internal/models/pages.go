package models

import "github.com/a-h/templ"

type PageData struct {
	Title string
	Page  templ.Component
}

type HomeData struct {
	Username string
}
