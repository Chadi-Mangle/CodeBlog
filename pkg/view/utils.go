package view

import (
	"html/template"
	"net/http"
)

type Page struct {
	Template *template.Template
	filename string
}

type View struct {
	Index Page
	Show  Page
	New   Page
	Edit  Page
}

func (self *Page) Render(w http.ResponseWriter, data any) error {
	return self.Template.ExecuteTemplate(w, self.filename, data)
}
