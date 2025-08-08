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

func (s *ServeTemplate) HandleTemplate(pattern string, filename string, query func(r *http.Request) any) {
	s.HandleFunc(pattern, func(w http.ResponseWriter, r *http.Request) {
		data := query(r)
		if err := s.template.ExecuteTemplate(w, filename, data); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	})
}
