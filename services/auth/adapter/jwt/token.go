package jwt

import (
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/google/uuid"
	"github.com/takuya911/ddd_gql-grpc/services/auth/shared"
)

type (
	token struct {
		env shared.Env
	}
	jwtClaims struct {
		User struct {
			UUID string `json:"id"`
		} `json:"user"`
		jwt.StandardClaims
	}
)

// NewToken はインスタンスを生成します
func NewToken(env shared.Env) shared.Token {
	return &token{env}
}

// GenTokenPair func
func (t *token) GenTokenPair(userUUID string) (string, string, error) {
	tokenSubject := t.env.GetTokenSubject()
	refreshTokenSubject := t.env.GetRefreshTokenSubject()
	jwtSignKey := t.env.GetJwtSignKey()

	accessToken, err := genToken(userUUID, tokenSubject, jwtSignKey, 3600)
	if err != nil {
		return "", "", err
	}
	refreshToken, err := genToken(userUUID, refreshTokenSubject, jwtSignKey, 3600)
	if err != nil {
		return "", "", err

	}
	return accessToken, refreshToken, nil
}

func genToken(userUUID string, sub string, jwtSignKey string, expSec int64) (string, error) {

	expTime := time.Now().Add(time.Duration(expSec) * time.Second)
	claims := &jwtClaims{
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

	privateKey, err := jwt.ParseECPrivateKeyFromPEM([]byte(jwtSignKey))
	if err != nil {
		return "", err
	}

	tokenString, err := token.SignedString(privateKey)
	if err != nil {
		return "", err
	}
	return tokenString, nil
}
