package env

import (
	"os"

	"github.com/takuya911/project-services/services/gql/shared"
)

type (
	environment struct {
		graphqlServerPort string
		userServerName    string
		userServerPort    string
	}
)

// NewEnv はコンストラクタです
func NewEnv() shared.Env {
	graphqlServerPort := os.Getenv("GRAPHQL_SERVICE_PORT")
	userServerName := os.Getenv("USER_SERVICE_NAME")
	userServerPort := os.Getenv("USER_SERVICE_PORT")
	return &environment{
		graphqlServerPort,
		userServerName,
		userServerPort,
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
