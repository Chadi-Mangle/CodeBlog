package router

import (
	"html/template"
	"net/http"
)

type ServeTemplate struct {
	*http.ServeMux
	template *template.Template
}

func NewServeTemplate(template *template.Template) *ServeTemplate {
	return &ServeTemplate{
		ServeMux: http.NewServeMux(),
		template: template,
	}
}

func (s *ServeTemplate) HandleTemplate(pattern string, template *template.Template, filename string, data any) {
	s.HandleFunc(pattern, func(w http.ResponseWriter, r *http.Request) {
		template.ExecuteTemplate(w, filename, data)
	})
}
