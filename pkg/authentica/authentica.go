package authentica

import (
	"github.com/matiniiuu/mcommon/pkg/authentica/enums"
	"github.com/matiniiuu/mcommon/pkg/mconfig"
)

type (
	AuthenticaSendRequest struct {
		TemplateId     string            `json:"template_id,omitempty"`
		Phone          string            `json:"phone,omitempty"`
		Method         enums.Method      `json:"method,omitempty"`
		NumberOfDigits int               `json:"number_of_digits,omitempty"`
		OtpFormat      enums.OtpFormat   `json:"otp_format,omitempty"`
		IsFallbackOn   bool              `json:"is_fallback_on,omitempty"`
		FallbackMethod enums.Method      `json:"fallback_method,omitempty"`
		FallbackPhone  string            `json:"fallback_phone,omitempty"`
		Otp            int               `json:"otp,omitempty"`
		SenderName     string            `json:"sender_name,omitempty"`
		Environment    enums.Environment `json:"environment"`
	}
	Authentica interface {
		Send(*AuthenticaSendRequest) error
	}
)

func New(cfg *mconfig.Authentica) (Authentica, error) {
	return &authentica{authorizationKey: cfg.AuthorizationKey}, nil
}
