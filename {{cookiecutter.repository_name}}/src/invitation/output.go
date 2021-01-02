package invitation

import (
	"github.com/{{cookiecutter.organization_name}}/{{cookiecutter.repository_name}}/pkg/domain/model"
)

type FindCodeOutput struct {
	Invitation model.Invitation
}

type AddCodeOutput struct {
	Invitation model.Invitation
}
