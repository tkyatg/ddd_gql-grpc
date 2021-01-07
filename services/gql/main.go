package main

import (
	"context"
	"log"
	"net/http"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/takuya911/project-services/services/gql/adapter/env"
	"github.com/takuya911/project-services/services/gql/graph/generated"
	"github.com/takuya911/project-services/services/gql/graph/resolver"
)

const defaultPort = "8080"

func main() {
	env := env.NewEnv()
	ctx := context.Background()

	srv := handler.NewDefaultServer(
		generated.NewExecutableSchema(
			generated.Config{
				Resolvers: resolver.NewResolver(ctx, env),
			},
		),
	)

	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", srv)

	log.Fatal(http.ListenAndServe(":"+env.GetGraphqlServerPort(), nil))

}
