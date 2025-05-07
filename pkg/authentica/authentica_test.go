package authentica

import (
	"testing"

	"github.com/matiniiuu/mcommon/pkg/authentica/enums"
	"github.com/matiniiuu/mcommon/pkg/mconfig"
)

func TestSend_Live(t *testing.T) {
	authentica := New(&mconfig.Authentica{
		AuthorizationKey: "test",
	})

	req := &AuthenticaSendRequest{
		TemplateId:     "6",
		Phone:          "+966....",
		Method:         enums.SMS,
		NumberOfDigits: 4,
		OtpFormat:      enums.Numeric,
		IsFallbackOn:   true,
		FallbackMethod: enums.Whatsapp,
		FallbackPhone:  "+966....",
		Otp:            1919,
		Environment:    enums.MockServer,
	}

	err := authentica.Send(req)
	if err != nil {
		t.Fatalf("Real API call failed: %v", err)
	}
	t.Log("OTP sent successfully")
}
