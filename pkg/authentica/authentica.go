package authentica

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/matiniiuu/mcommon/pkg/derrors"
	"github.com/matiniiuu/mcommon/pkg/mconfig"
)

type (
	IAuthentica interface {
		Send(*AuthenticaSendRequest) error
	}
	Authentica struct{ authorizationKey string }
)

func New(cfg *mconfig.Authentica) IAuthentica {
	return &Authentica{authorizationKey: cfg.AuthorizationKey}
}

func (a *Authentica) Send(dto *AuthenticaSendRequest) error {
	client := &http.Client{}

	body, _ := json.Marshal(dto)
	req, _ := http.NewRequest("POST", fmt.Sprintf("%s/send-otp", dto.Environment.GetURL()), bytes.NewBuffer(body))

	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Accept", "application/json")
	req.Header.Add("X-Authorization", a.authorizationKey)

	resp, err := client.Do(req)
	if err != nil {
		return err
	}

	defer resp.Body.Close()
	if resp.StatusCode == 200 {
		return nil
	}
	resp_body, _ := io.ReadAll(resp.Body)
	return derrors.New(derrors.KindUnexpected, string(resp_body))
}
