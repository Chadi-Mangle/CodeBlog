package main

import (
	"embed"
	_ "embed"
	"errors"
	"fmt"
	"html/template"
	"net/http"

	"github.com/Chadi-Mangle/CodeBlog/pkg/controller"
	"github.com/Chadi-Mangle/CodeBlog/pkg/view"
)

//go:embed templates
var tmplFS embed.FS

func main() {
	port := ":8080"
	tmpl := template.Must(template.ParseFS(tmplFS, "templates/*"))

	router := http.NewServeMux()

	index := view.Page{
		Template: tmpl,
		Filename: "index.html",
	}

	IndexView := view.View{
		Index: index,
	}

	IndexController := controller.NewController(&IndexView, nil)

	router.HandleFunc("GET /", IndexController.Index(
		func(r *http.Request) any {
			return struct{ Name string }{Name: "World"}
		}))

	srv := http.Server{
		Addr:    port,
		Handler: router,
	}

	fmt.Printf("Starting website at localhost%s", port)

	if err := srv.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
		fmt.Println("An error occured:", err)
	}
}
