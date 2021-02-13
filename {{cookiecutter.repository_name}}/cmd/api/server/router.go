package main

import (
	"net/http"

	"github.com/go-chi/chi"

	"github.com/{{cookiecutter.organization_name}}/{{cookiecutter.repository_name}}/cmd/api/adapter/invitation"
	"github.com/{{cookiecutter.organization_name}}/{{cookiecutter.repository_name}}/pkg/presenter"
	"github.com/{{cookiecutter.organization_name}}/{{cookiecutter.repository_name}}/pkg/registry"
)

func NewRouter(repo registry.Repository) http.Handler {
	r := chi.NewRouter()

	invitationHandler := invitation.NewHandler(repo)

	r = commonMiddleware(r)

	r.Mount("/invitations", invitation.NewRouter(invitationHandler))
	r.Get("/health", func(w http.ResponseWriter, r *http.Request) {
		presenter.Success(w)
	})

	return r
}
