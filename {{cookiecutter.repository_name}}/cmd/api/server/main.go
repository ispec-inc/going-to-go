package main

import (
	"net/http"

	"github.com/{{cookiecutter.organization_name}}/{{cookiecutter.repository_name}}/pkg/config"
	"github.com/{{cookiecutter.organization_name}}/{{cookiecutter.repository_name}}/pkg/registry"
)

func main() {
	config.Init()

	repo, cleanup := registry.NewRepository()
	defer cleanup()

	r := NewRouter(repo)
	http.ListenAndServe(":9000", r)
}
