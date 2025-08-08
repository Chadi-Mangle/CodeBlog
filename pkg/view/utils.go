package view

import (
	"html/template"
	"net/http"
)

type Page struct {
	Template *template.Template
	Filename string
}

type View struct {
	Index Page
	Show  Page
	New   Page
	Edit  Page
}

func (p *Page) Render(w http.ResponseWriter, data any) error {
	return p.Template.ExecuteTemplate(w, p.Filename, data)
}
