package rpc

import (
	"github.com/jinzhu/gorm"

	definition "github.com/takuya911/project-auth-definition"
	"github.com/takuya911/project-services/services/auth/adapter/env"
	"github.com/takuya911/project-services/services/auth/adapter/hash"
	"github.com/takuya911/project-services/services/auth/adapter/jwt"
	authenticationcommandservice "github.com/takuya911/project-services/services/auth/commands/authenticationCommandService"
	"github.com/takuya911/project-services/services/auth/domain"
	authqueryservice "github.com/takuya911/project-services/services/auth/queries/authQueryService"
)

func (s *server) registerServices(dbConnection *gorm.DB) {
	env := env.NewEnv()
	token := jwt.NewToken(env)
	hash := hash.NewHash()

	// auth query service
	authQueryUsecase := authqueryservice.NewUsecase(token)
	authQueryServer := authqueryservice.NewServer(authQueryUsecase)

	// auth command service
	authenticationDataAccessor := domain.NewAuthenticationDataAccessor(dbConnection)
	authenticationRepository := domain.NewAuthenticationRepository(authenticationDataAccessor)
	authenticationCommandUsecase := authenticationcommandservice.NewUsecase(authenticationRepository, token, hash)
	authenticationCommandServer := authenticationcommandservice.NewServer(authenticationCommandUsecase)

	definition.RegisterAuthQueryServiceServer(s.rpc, authQueryServer)
	definition.RegisterAuthCommandServiceServer(s.rpc, authenticationCommandServer)
}
