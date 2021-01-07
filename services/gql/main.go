package main

import (
	"context"
	"log"
	"net/http"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/takuya911/project-services/services/gql/adapter/env"
	"github.com/takuya911/project-services/services/gql/graph"
	"github.com/takuya911/project-services/services/gql/graph/generated"
	definition "github.com/takuya911/project-user-definition"
	"google.golang.org/grpc"
)

const defaultPort = "8080"

func main() {
	env := env.NewEnv()

	ctx := context.Background()

	opts := []grpc.DialOption{grpc.WithInsecure()}
	conn, err := grpc.DialContext(ctx, env.GetUserServerName()+":"+env.GetUserServerPort(), opts...)
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	userC := definition.NewUserQueryServiceClient(conn)

	// Regist handler
	resolver := graph.NewResolver(userC)
	srv := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: resolver}))

	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", srv)
	log.Printf("connect to http://localhost:%s/ for GraphQL playground", env.GetGraphqlServerPort())
	log.Fatal(http.ListenAndServe(":"+env.GetGraphqlServerPort(), nil))

}
