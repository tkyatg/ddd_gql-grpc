package hash

import (
	"errors"

	"github.com/tkyatg/ddd_gql-grpc/services/user/shared"
	"golang.org/x/crypto/bcrypt"
)

type (
	hash struct {
	}
)

// NewHash func
func NewHash() shared.Hash {
	return &hash{}
}

// GenEncryptedPass func
func (h *hash) GenEncryptedPassword(password string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hash), nil
}

// CompareHashAndPass func
func (h *hash) CompareHashAndPass(dbPassword string, formPassword string) error {
	if err := bcrypt.CompareHashAndPassword([]byte(dbPassword), []byte(formPassword)); err != nil {
		return errors.New(shared.PasswordIsIncorrect)
	}
	return nil
}
