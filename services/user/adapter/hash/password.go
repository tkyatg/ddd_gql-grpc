package hash

import "golang.org/x/crypto/bcrypt"

// GenEncryptedPass func
func GenEncryptedPass(password string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hash), nil
}
