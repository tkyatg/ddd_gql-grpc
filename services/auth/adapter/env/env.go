package env

import (
	"os"

	"github.com/takuya911/ddd_gql-grpc/services/auth/shared"
)

type (
	environment struct {
		dbHost              string
		dbPort              string
		dbUser              string
		dbPassword          string
		dbName              string
		servicePort         string
		tokenSubject        string
		refreshTokenSubject string
		jwtSignKey          string
	}
)

// NewEnv func
func NewEnv() shared.Env {
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	servicePort := os.Getenv("AUTH_SERVICE_PORT")
	tokenSubject := os.Getenv("")
	refreshTokenSubject := os.Getenv("")
	jwtSignKey := os.Getenv("")

	return &environment{
		dbHost,
		dbPort,
		dbUser,
		dbPassword,
		dbName,
		servicePort,
		tokenSubject,
		refreshTokenSubject,
		jwtSignKey,
	}
}

func (t *environment) GetDBHost() string {
	return t.dbHost
}
func (t *environment) GetDBPort() string {
	return t.dbPort
}
func (t *environment) GetDBUser() string {
	return t.dbUser
}
func (t *environment) GetDBPassword() string {
	return t.dbPassword
}
func (t *environment) GetDBName() string {
	return t.dbName
}
func (t *environment) GetAuthServicePort() string {
	return t.servicePort
}
func (t *environment) GetTokenSubject() string {
	return t.tokenSubject
}
func (t *environment) GetRefreshTokenSubject() string {
	return t.refreshTokenSubject
}
func (t *environment) GetJwtSignKey() string {
	return t.jwtSignKey
}
