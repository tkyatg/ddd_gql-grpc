package hash

import (
	"errors"

	"github.com/takuya911/project-services/services/auth/shared"
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
func (h *hash) CompareHashAndPass(dbPassword string, formPassword string) error {
	if err := bcrypt.CompareHashAndPassword([]byte(dbPassword), []byte(formPassword)); err != nil {
		return errors.New(shared.PasswordIsIncorrect)
	}
	return nil
}
