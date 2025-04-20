package authentica

import (
	"testing"

	"github.com/matiniiuu/mcommon/pkg/authentica/enums"
)

func TestSend_Live(t *testing.T) {
	auth := &authentica{
		authorizationKey: "test",
	}

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

	err := auth.Send(req)
	if err != nil {
		t.Fatalf("Real API call failed: %v", err)
	}
	t.Log("OTP sent successfully")
}
