package main

import (
	"fmt"
	"log"
	"net"

	"github.com/takuya911/project-services/services/user/adapter/env"
	"github.com/takuya911/project-services/services/user/adapter/sql"
	usercommandservice "github.com/takuya911/project-services/services/user/commands/userCommandService"
	"github.com/takuya911/project-services/services/user/domain"
	userqueryservice "github.com/takuya911/project-services/services/user/queries/userQueryService"
	definition "github.com/takuya911/project-user-definition"
	"google.golang.org/grpc"
)

func main() {
	env := env.NewEnv()
	dbConnection, err := sql.NewGormConnect(env.GetDBUser(), env.GetDBPassword(), env.GetDBName(), env.GetDBHost(), env.GetDBPort())
	if err != nil {
		log.Fatal(err)
	}
	defer dbConnection.Close()

	listenPort, err := net.Listen("tcp", fmt.Sprintf(":%s", env.GetUserServicePort()))
	if err != nil {
		log.Fatal(err)
	}

	server := grpc.NewServer()

	// user query service
	userQueryDataAccessor := userqueryservice.NewDataAccessor(dbConnection)
	userQueryUsecase := userqueryservice.NewUsecase(userQueryDataAccessor)
	userQueryServer := userqueryservice.NewServer(userQueryUsecase)
	// user command service
	userCommandDataAccessor := domain.NewUserDataAccessor(dbConnection)
	userCommandRepository := domain.NewUserRepository(userCommandDataAccessor)
	userCommandUsecase := usercommandservice.NewUsecase(userCommandRepository)
	userCommandServer := usercommandservice.NewServer(userCommandUsecase)
	definition.RegisterUserQueryServiceServer(server, userQueryServer)
	definition.RegisterUserCommandServiceServer(server, userCommandServer)

	if err := server.Serve(listenPort); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
