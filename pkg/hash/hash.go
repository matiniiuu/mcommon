package hash

import (
	"crypto/rand"
	"encoding/base64"
	"errors"
	"fmt"
	"strings"

	"github.com/matiniiuu/mcommon/pkg/derrors"
	"github.com/matiniiuu/mcommon/pkg/translator/messages"
	"golang.org/x/crypto/argon2"
	"golang.org/x/crypto/bcrypt"
)

type HashMethod string

const (
	Bcrypt HashMethod = "bcrypt"
	Argon2 HashMethod = "argon2"
)

type HashedPassword struct {
	Method HashMethod
	Hash   string
}

func GenerateSalt(size int) (string, error) {
	salt := make([]byte, size)
	if _, err := rand.Read(salt); err != nil {
		return "", derrors.New(derrors.KindUnexpected, messages.GeneralError)
	}
	return base64.RawStdEncoding.EncodeToString(salt), nil
}

func HashPassword(password string, method HashMethod) (string, error) {
	switch method {
	case Bcrypt:
		bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
		if err != nil {
			return "", derrors.New(derrors.KindUnexpected, messages.GeneralError)
		}
		return fmt.Sprintf("%s:%s", Bcrypt, string(bytes)), nil
	case Argon2:
		salt, err := GenerateSalt(16)
		if err != nil {
			return "", err
		}
		hash := argon2.IDKey([]byte(password), []byte(salt), 1, 64*1024, 4, 32)
		return fmt.Sprintf("%s:%s:%s", Argon2, salt, base64.RawStdEncoding.EncodeToString(hash)), nil

	default:
		return "", derrors.New(derrors.KindInvalid, messages.UnsupportedHashingMethod)
	}
}

func CheckPassword(password, hashed string) error {
	parts := strings.Split(hashed, ":")
	if len(parts) < 2 {
		return derrors.New(derrors.KindInvalid, messages.GeneralError)
	}
	method := parts[0]
	switch HashMethod(method) {
	case Bcrypt:
		err := bcrypt.CompareHashAndPassword([]byte(parts[1]), []byte(password))
		if err != nil {
			if errors.Is(err, bcrypt.ErrMismatchedHashAndPassword) {
				return derrors.New(derrors.KindInvalid, messages.EmailOrPasswordIsIncorrect)
			}
			return derrors.New(derrors.KindUnexpected, messages.GeneralError)
		}
		return nil
	case Argon2:
		if len(parts) < 3 {
			return derrors.New(derrors.KindInvalid, messages.GeneralError)
		}
		salt := parts[1]
		hash := parts[2]
		computedHash := argon2.IDKey([]byte(password), []byte(salt), 1, 64*1024, 4, 32)
		if base64.RawStdEncoding.EncodeToString(computedHash) != hash {
			return derrors.New(derrors.KindInvalid, messages.EmailOrPasswordIsIncorrect)
		}
		return nil
	default:
		return derrors.New(derrors.KindInvalid, messages.UnsupportedHashingMethod)
	}
}
