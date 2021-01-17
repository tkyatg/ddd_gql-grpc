package rpc

import (
	"github.com/jinzhu/gorm"
	"github.com/takuya911/project-services/services/user/adapter/hash"
	usercommandservice "github.com/takuya911/project-services/services/user/commands/userCommandService"
	"github.com/takuya911/project-services/services/user/domain"
	userqueryservice "github.com/takuya911/project-services/services/user/queries/userQueryService"
	definition "github.com/takuya911/project-user-definition"
)

func (s *server) registerServices(dbConnection *gorm.DB) {
	hash := hash.NewHash()
	// user query service
	userQueryDataAccessor := userqueryservice.NewDataAccessor(dbConnection)
	userQueryUsecase := userqueryservice.NewUsecase(userQueryDataAccessor, hash)
	userQueryServer := userqueryservice.NewServer(userQueryUsecase)

	// user command service
	userCommandDataAccessor := domain.NewUserDataAccessor(dbConnection)
	userCommandRepository := domain.NewUserRepository(userCommandDataAccessor)
	userCommandUsecase := usercommandservice.NewUsecase(userCommandRepository, hash)
	userCommandServer := usercommandservice.NewServer(userCommandUsecase)

	definition.RegisterUserQueryServiceServer(s.rpc, userQueryServer)
	definition.RegisterUserCommandServiceServer(s.rpc, userCommandServer)
}
