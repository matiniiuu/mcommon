package hash

import (
	"errors"

	"github.com/matiniiuu/mcommon/pkg/derrors"
	"github.com/matiniiuu/mcommon/pkg/translator/messages"
	"golang.org/x/crypto/bcrypt"
)

func Password(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", derrors.New(derrors.KindUnexpected, messages.GeneralError)
	}

	return string(bytes), nil
}

func CheckPassword(password, hash string) error {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	if err != nil {
		if errors.Is(err, bcrypt.ErrMismatchedHashAndPassword) {
			return derrors.New(derrors.KindInvalid, messages.EmailOrPasswordIsIncorrect)
		}

		return derrors.New(derrors.KindUnexpected, messages.GeneralError)
	}

	return nil
}
