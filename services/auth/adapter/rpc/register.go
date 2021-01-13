package rpc

import (
	"github.com/jinzhu/gorm"

	definition "github.com/takuya911/project-auth-definition"
	authqueryservice "github.com/takuya911/project-services/services/auth/queries/authQueryService"
)

func (s *server) registerServices(dbConnection *gorm.DB) {
	// auth query service
	authQueryUsecase := authqueryservice.NewUsecase()
	authQueryServer := authqueryservice.NewServer(authQueryUsecase)

	definition.RegisterAuthQueryServiceServer(s.rpc, authQueryServer)
}
