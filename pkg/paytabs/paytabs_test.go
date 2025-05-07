package paytabs

import (
	"testing"

	"github.com/matiniiuu/mcommon/pkg/mconfig"
	"github.com/matiniiuu/mcommon/pkg/mutils"
	"github.com/matiniiuu/mcommon/pkg/paytabs/enums"
)

func TestCreateTransactionRequest(t *testing.T) {
	auth := New(&mconfig.Paytabs{
		ProfileId: 0,
		ServerKey: "",
		Region:    enums.SAU,
	})
	req := &CreateTransactionRequest{
		TransactionType:  enums.Sale,
		TransactionClass: enums.Ecom,
		CartId:           "uniqueid2",
		CartCurrency:     enums.SAR,
		CartAmount:       12.22,
		CartDescription:  "this is my description",
		CallbackUrl:      "https://webhook.site/",
		ReturnUrl:        "https://trip360.sa",
		HideShipping:     mutils.BoolPointer(true),
		CustomerDetails: &Details{
			Name:    "Matin Mansouri",
			Email:   "matinniiuu@gmail.com",
			Phone:   "+989194788031",
			Country: "Saudi arabia",
			State:   "Riyadh",
			City:    "Riyadh",
			Street1: "Riyadh king aziz",
			Zip:     "123123",
		},
	}

	result, err := auth.CreateTransaction(req)
	if err != nil {
		t.Fatalf("Real API call failed: %v", err)
	}
	t.Log(result)
}
func TestValidatePayment(t *testing.T) {
	auth := New(&mconfig.Paytabs{
		ProfileId: 0,
		ServerKey: "",
		Region:    enums.SAU,
	})
	req := &ValidatePaymentRequest{
		TransactionRef: "",
	}

	result, err := auth.ValidatePayment(req)
	if err != nil {
		t.Fatalf("Real API call failed: %v", err)
	}
	t.Log(result, result.Status)
}
