package invitation

import (
	"github.com/ispec-inc/going-to-go/example/pkg/view"
)

type GetCodeResponse struct {
	InvitationCode view.InvitationCode `json:"invitation_code"`
}

type AddCodeResponse struct {
	InvitationCode view.InvitationCode `json:"invitation_code"`
}
