package hash

import (
	"github.com/takuya911/project-services/services/user/shared"
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
