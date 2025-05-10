package paytabs

import "github.com/matiniiuu/mcommon/pkg/paytabs/enums"

type (
	Details struct {
		Name    string `json:"name"`
		Email   string `json:"email"`
		Phone   string `json:"phone"`
		Country string `json:"country"`
		State   string `json:"state"`
		City    string `json:"city"`
		Street1 string `json:"street1"`
		Zip     string `json:"zip"`
	}
	CreateTransactionResponse struct {
		TransactionRef string `json:"tran_ref"`
		RedirectUrl    string `json:"redirect_url"`
	}
	CreateTransactionRequest struct {
		ProfileId        string                        `json:"profile_id"`
		TransactionType  enums.PaytabsTransactionType  `json:"tran_type"`
		TransactionClass enums.PaytabsTransactionClass `json:"tran_class"`

		CartId          string                `json:"cart_id"`
		CartCurrency    enums.PaytabsCurrency `json:"cart_currency"`
		CartAmount      float64               `json:"cart_amount"`
		CartDescription string                `json:"cart_description"`

		CallbackUrl string `json:"callback"`
		ReturnUrl   string `json:"return"`

		//? Optional Parameters
		Language *string `json:"paypage_lang,omitempty"`

		CustomerDetails *Details `json:"customer_details,omitempty"`
		ShippingDetails *Details `json:"shipping_details,omitempty"`

		HideShipping *bool `json:"hide_shipping,omitempty"`

		Framed             *bool `json:"framed,omitempty"`
		FramedReturnTop    *bool `json:"framed_return_top,omitempty"`
		FramedReturnParent *bool `json:"framed_return_parent,omitempty"`
		ForceFullUi        *bool `json:"force_full_ui,omitempty"`
		//@ when this item is available return url must be 'None'
		FramedMessageContent *string `json:"framed_message_target,omitempty"`

		PaymentMethods []*enums.PaytabsPaymentMethod `json:"payment_methods,omitempty"`

		ThemeId *uint `json:"config_id,omitempty"`

		AlternateCurrency *string `json:"alt_currency"`
	}

	ValidatePaymentRequest struct {
		ProfileId      string `json:"profile_id"`
		TransactionRef string `json:"tran_ref"`
	}
	ValidatePaymentResponse struct {
		ProfileId      string `json:"profile_id"`
		TransactionRef string `json:"tran_ref"`
		PaymentResult  struct {
			ResponseStatus  string `json:"response_status"`
			ResponseCode    string `json:"response_code"`
			ResponseMessage string `json:"response_message"`
			TransactionTime string `json:"transaction_time"`
		} `json:"payment_result"`
		PaymentInfo struct {
			PaymentMethod      string `json:"payment_method"`
			CartType           string `json:"card_type"`
			CardScheme         string `json:"card_scheme"`
			PaymentDescription string `json:"payment_description"`
			ExpiryMonth        uint   `json:"expiryMonth"`
			ExpiryYear         uint   `json:"expiryYear"`
			IssuerCountry      string `json:"issuerCountry"`
			IssuerName         string `json:"issuerName"`
		} `json:"payment_info"`
		Status          enums.PaymentStatus `json:"status"`
		CustomerDetails Details             `json:"customer_details"`
		MerchantId      uint                `json:"merchantId"`
		Trace           string              `json:"trace"`
	}
)

// https://support.paytabs.com/en/support/solutions/articles/60000711358-what-is-response-code-vs-the-response-status-
