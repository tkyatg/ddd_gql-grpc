package rpc

import (
	"github.com/jinzhu/gorm"

	"github.com/takuya911/ddd_gql-grpc/services/auth/adapter/env"
	"github.com/takuya911/ddd_gql-grpc/services/auth/adapter/jwt"
	authqueryservice "github.com/takuya911/ddd_gql-grpc/services/auth/queries/authQueryService"
	definition "github.com/takuya911/project-auth-definition"
)

func (s *server) registerServices(dbConnection *gorm.DB) {
	env := env.NewEnv()
	token := jwt.NewToken(env)

	// auth query service
	authQueryUsecase := authqueryservice.NewUsecase(token)
	authQueryServer := authqueryservice.NewServer(authQueryUsecase)

	definition.RegisterAuthQueryServiceServer(s.rpc, authQueryServer)
}
