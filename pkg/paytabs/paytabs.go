package paytabs

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/matiniiuu/mcommon/pkg/derrors"
	"github.com/matiniiuu/mcommon/pkg/mconfig"
	"github.com/matiniiuu/mcommon/pkg/paytabs/enums"
)

/**
 * @description Paytabs rest api integration for golang
 * @author Matin Mansouri <matiniiuu@gmail.com>
 */

type (
	IPaytabs interface {
		CreateTransaction(dto *CreateTransactionRequest) (*CreateTransactionResponse, error)
		ValidatePayment(dto *ValidatePaymentRequest) (*ValidatePaymentResponse, error)
	}
	Paytabs struct {
		profileId string
		serverKey string
		region    enums.PaytabsRegion
	}
)

func New(cfg *mconfig.Paytabs) IPaytabs {
	return Paytabs{profileId: cfg.ProfileId, serverKey: cfg.ServerKey, region: cfg.Region}
}

func (p Paytabs) CreateTransaction(dto *CreateTransactionRequest) (*CreateTransactionResponse, error) {
	client := &http.Client{}
	dto.ProfileId = p.profileId
	body, _ := json.Marshal(dto)
	req, _ := http.NewRequest("POST", fmt.Sprintf("%s/payment/request", p.region.GetURL()), bytes.NewBuffer(body))

	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Accept", "application/json")
	req.Header.Add("authorization", p.serverKey)

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()
	resp_body, _ := io.ReadAll(resp.Body)
	if resp.StatusCode == 200 {
		response := new(CreateTransactionResponse)
		if err := json.Unmarshal(resp_body, &response); err != nil {
			return nil, err
		}
		return response, nil
	}
	return nil, derrors.New(derrors.KindUnexpected, string(resp_body))
}

func (p Paytabs) ValidatePayment(dto *ValidatePaymentRequest) (*ValidatePaymentResponse, error) {
	client := &http.Client{}
	dto.ProfileId = p.profileId
	body, _ := json.Marshal(dto)
	req, _ := http.NewRequest("POST", fmt.Sprintf("%s/payment/query", p.region.GetURL()), bytes.NewBuffer(body))

	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Accept", "application/json")
	req.Header.Add("authorization", p.serverKey)

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()
	resp_body, _ := io.ReadAll(resp.Body)
	if resp.StatusCode != 200 {
		return nil, derrors.New(derrors.KindUnexpected, string(resp_body))
	}
	response := new(ValidatePaymentResponse)
	if err := json.Unmarshal(resp_body, &response); err != nil {
		return nil, err
	}
	switch response.PaymentResult.ResponseStatus {
	case "A":
		response.Status = enums.PaymentStatusAuthorized
	case "H":
		response.Status = enums.PaymentStatusHold
	case "P":
		response.Status = enums.PaymentStatusPending
	case "V":
		response.Status = enums.PaymentStatusVoided
	case "E":
		response.Status = enums.PaymentStatusError
	case "D":
		response.Status = enums.PaymentStatusDeclined
	case "X":
		response.Status = enums.PaymentStatusExpired
	}

	return response, nil
}
