package main

import (
	"embed"
	_ "embed"
	"errors"
	"fmt"
	"html/template"
	"net/http"

	"github.com/Chadi-Mangle/CodeBlog/pkg/router"
)

//go:embed templates
var tmplFS embed.FS

func main() {
	port := ":8080"
	tmpl := template.Must(template.ParseFS(tmplFS, "templates/*"))

	router := router.NewServeTemplate(tmpl)

	type Data struct{ Name string }

	router.HandleTemplate("GET /{$}", tmpl, "index.html", Data{Name: "World !"})

	srv := http.Server{
		Addr:    port,
		Handler: router,
	}

	fmt.Printf("Starting website at localhost%s", port)

	if err := srv.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
		fmt.Println("An error occured:", err)
	}
}
