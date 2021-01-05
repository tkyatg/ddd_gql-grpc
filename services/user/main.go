package main

import (
	"log"

	"github.com/takuya911/project-services/services/user/adapter/env"
	"github.com/takuya911/project-services/services/user/adapter/rpc"
	"github.com/takuya911/project-services/services/user/adapter/sql"
	usercommandservice "github.com/takuya911/project-services/services/user/commands/userCommandService"
	"github.com/takuya911/project-services/services/user/domain"
	userqueryservice "github.com/takuya911/project-services/services/user/queries/userQueryService"
)

func main() {
	env := env.NewEnv()
	dbConnection, err := sql.NewGormConnect(env.GetDBUser(), env.GetDBPassword(), env.GetDBName(), env.GetDBHost(), env.GetDBPort())
	if err != nil {
		log.Fatal(err)
	}
	defer dbConnection.Close()

	server, err := rpc.NewServer(env.GetUserServicePort())
	if err != nil {
		log.Fatal(err)
	}

	// user query service
	userQueryDataAccessor := userqueryservice.NewDataAccessor(dbConnection)
	userQueryUsecase := userqueryservice.NewUsecase(userQueryDataAccessor)
	userQueryServer := userqueryservice.NewServer(userQueryUsecase)
	// user command service
	userCommandDataAccessor := domain.NewUserDataAccessor(dbConnection)
	userCommandRepository := domain.NewUserRepository(userCommandDataAccessor)
	userCommandUsecase := usercommandservice.NewUsecase(userCommandRepository)
	userCommandServer := usercommandservice.NewServer(userCommandUsecase)

	server.RegisterService(userQueryServer, userCommandServer)
	if err := server.Serve(); err != nil {
		log.Fatal(err)
	}
}
