package shared

import (
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/google/uuid"
)

const (
	tokenSubject        = "id_token"
	refreshTokenSubject = "refresh_token"
	signKey             = "-----BEGIN EC PRIVATE KEY-----\nMIHcAgEBBEIAGT88ebOnAbgmS9Idbns1+VqWV9UN2dvzqiXMmxvAyKNnpoFQxYEL\nLrvmL9uqZaCcbR7EOz5OQyyozKyfqxNiMcigBwYFK4EEACOhgYkDgYYABAB/PCXh\nMMmfHGuR2vm7NLtaa1Jg25CuldjD3LlpFAbrQ0tkfnvskJYRkuFJcbbMGEDLKwvz\nH/HCCi/k8lmynF/DlwH4EAVQTUhkoHO2AUS5zK5oDTKxPN8v86BXBBtbbdVEjZaL\na6hVSC8VOiQD+NeSCWwdN2pY0gYCQHcvxyrCqvAX9Q==\n-----END EC PRIVATE KEY-----"
)

type (
	// TokenPair struct
	TokenPair struct {
		AccessToken  string
		RefreshToken string
	}
	// JwtClaims struct
	JwtClaims struct {
		User struct {
			UUID string `json:"id"`
		} `json:"user"`
		jwt.StandardClaims
	}
	// GenToken interface
	GenToken interface {
		GenTokenPair(userUUID string) (string, string, error)
	}
)

// GenTokenPair func
func GenTokenPair(userUUID string) (TokenPair, error) {
	accessToken, err := genToken(userUUID, tokenSubject, 3600)
	if err != nil {
		return TokenPair{}, err
	}
	refreshToken, err := genToken(userUUID, refreshTokenSubject, 3600)
	if err != nil {
		return TokenPair{}, err

	}
	return TokenPair{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}, nil
}

func genToken(userUUID string, sub string, expSec int64) (string, error) {
	expTime := time.Now().Add(time.Duration(expSec) * time.Second)
	claims := &JwtClaims{
		struct {
			UUID string `json:"id"`
		}{
			UUID: userUUID,
		},
		jwt.StandardClaims{
			Id:        uuid.New().String(),
			Subject:   sub,
			ExpiresAt: expTime.Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodES512, claims)

	privateKey, err := jwt.ParseECPrivateKeyFromPEM([]byte(signKey))
	if err != nil {
		return "", err
	}

	tokenString, err := token.SignedString(privateKey)
	if err != nil {
		return "", err
	}
	return tokenString, nil
}
