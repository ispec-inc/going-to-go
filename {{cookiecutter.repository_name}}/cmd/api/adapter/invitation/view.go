package invitation

import (
	"github.com/{{cookiecutter.organization_name}}/{{cookiecutter.repository_name}}/pkg/view"
)

type GetCodeResponse struct {
	InvitationCode view.InvitationCode `json:"invitation_code"`
}

type AddCodeResponse struct {
	InvitationCode view.InvitationCode `json:"invitation_code"`
}
