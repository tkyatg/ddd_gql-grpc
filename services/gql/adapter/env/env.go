package env

import (
	"os"

	"github.com/tkyatg/ddd_gql-grpc/services/gql/shared"
)

type (
	environment struct {
		graphqlServerPort string
		userServerName    string
		userServerPort    string
		authServerName    string
		authServerPort    string
	}
)

// NewEnv はコンストラクタです
func NewEnv() shared.Env {
	graphqlServerPort := os.Getenv("GRAPHQL_SERVICE_PORT")
	userServerName := os.Getenv("USER_SERVICE_NAME")
	userServerPort := os.Getenv("USER_SERVICE_PORT")
	authServerName := os.Getenv("AUTH_SERVICE_NAME")
	authServerPort := os.Getenv("AUTH_SERVICE_PORT")
	return &environment{
		graphqlServerPort,
		userServerName,
		userServerPort,
		authServerName,
		authServerPort,
	}
}

func (t *environment) GetGraphqlServerPort() string {
	return t.graphqlServerPort
}
func (t *environment) GetUserServerName() string {
	return t.userServerName
}
func (t *environment) GetUserServerPort() string {
	return t.userServerPort
}
func (t *environment) GetAuthServerName() string {
	return t.authServerName
}
func (t *environment) GetAuthServerPort() string {
	return t.authServerPort
}
