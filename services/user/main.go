package main

import (
	"log"

	"github.com/takuya911/ddd_gql-grpc/services/user/adapter/env"
	"github.com/takuya911/ddd_gql-grpc/services/user/adapter/rpc"
	"github.com/takuya911/ddd_gql-grpc/services/user/adapter/sql"
)

func main() {
	env := env.NewEnv()
	dbConnection, err := sql.NewGormConnect(env.GetDBUser(), env.GetDBPassword(), env.GetDBName(), env.GetDBHost(), env.GetDBPort())
	if err != nil {
		log.Fatal(err)
	}
	defer dbConnection.Close()

	server, err := rpc.NewServer(env.GetUserServicePort(), dbConnection)
	if err != nil {
		log.Fatal(err)
	}

	if err := server.Serve(); err != nil {
		log.Fatal(err)
	}
}
