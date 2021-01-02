package invitation

import (
	"github.com/{{cookiecutter.organization_name}}/{{cookiecutter.repository_name}}/pkg/domain/model"
)

type FindCodeInput struct {
	ID int64
}

type AddCodeInput struct {
	Invitation model.Invitation
}
