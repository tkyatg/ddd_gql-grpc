package main

import (
	"log"
	"net"
	"os"

	"github.com/takuya911/project-services/services/user/adapter/env"
	userqueryservice "github.com/takuya911/project-services/services/user/queries/userQueryService"
	definition "github.com/takuya911/project-user-definition"
	"google.golang.org/grpc"
)

func main() {
	// db connect
	dbConn, err := env.NewGormConnect()
	if err != nil {
		log.Fatal(err)
	}
	defer dbConn.Close()

	lis, err := net.Listen("tcp", ":"+os.Getenv("USER_SERVICE_PORT"))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	// userservice
	userQueryDataAccessor := userqueryservice.NewDataAccessor(dbConn)
	userQueryUsecase := userqueryservice.NewUsecase(userQueryDataAccessor)
	userQueryServer := userqueryservice.NewServer(userQueryUsecase)

	server := grpc.NewServer()
	definition.RegisterUserServiceServer(server, userQueryServer)

	if err := server.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
