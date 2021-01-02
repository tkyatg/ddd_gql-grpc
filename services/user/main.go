package main

import (
	"log"
	"net"
	"os"

	"github.com/takuya911/project-services/services/user/adapter/sql"
	usercommandservice "github.com/takuya911/project-services/services/user/commands/userCommandService"
	domain "github.com/takuya911/project-services/services/user/domain"
	userqueryservice "github.com/takuya911/project-services/services/user/queries/userQueryService"
	definition "github.com/takuya911/project-user-definition"
	"google.golang.org/grpc"
)

func main() {
	// db connect
	dbConn, err := sql.NewGormConnect()
	if err != nil {
		log.Fatal(err)
	}
	defer dbConn.Close()

	lis, err := net.Listen("tcp", ":"+os.Getenv("USER_SERVICE_PORT"))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	// user query service
	userQueryDataAccessor := userqueryservice.NewDataAccessor(dbConn)
	userQueryUsecase := userqueryservice.NewUsecase(userQueryDataAccessor)
	userQueryServer := userqueryservice.NewServer(userQueryUsecase)
	// user command service
	userCommandDataAccessor := domain.NewUserDataAccessor(dbConn)
	userCommandRepository := domain.NewUserRepository(userCommandDataAccessor)
	userCommandUsecase := usercommandservice.NewUsecase(userCommandRepository)
	userCommandServer := usercommandservice.NewServer(userCommandUsecase)

	server := grpc.NewServer()
	definition.RegisterUserQueryServiceServer(server, userQueryServer)
	definition.RegisterUserCommandServiceServer(server, userCommandServer)

	if err := server.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
