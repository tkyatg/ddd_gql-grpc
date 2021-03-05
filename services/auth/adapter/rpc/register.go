package rpc

import (
	"github.com/jinzhu/gorm"
	definition "github.com/tkyatg/auth-definition"
	"github.com/tkyatg/ddd_gql-grpc/services/auth/adapter/env"
	"github.com/tkyatg/ddd_gql-grpc/services/auth/adapter/jwt"
	authqueryservice "github.com/tkyatg/ddd_gql-grpc/services/auth/queries/authQueryService"
)

func (s *server) registerServices(dbConnection *gorm.DB) {
	env := env.NewEnv()
	token := jwt.NewToken(env)

	// auth query service
	authQueryUsecase := authqueryservice.NewUsecase(token)
	authQueryServer := authqueryservice.NewServer(authQueryUsecase)

	definition.RegisterAuthQueryServiceServer(s.rpc, authQueryServer)
}
