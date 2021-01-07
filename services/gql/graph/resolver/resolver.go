package resolver

import (
	"context"

	"github.com/takuya911/project-services/services/gql/graph/generated"
	"github.com/takuya911/project-services/services/gql/shared"
	definition "github.com/takuya911/project-user-definition"
	"google.golang.org/grpc"
)

type resolver struct {
	userClient definition.UserQueryServiceClient
}

// NewResolver function
func NewResolver(ctx context.Context, env shared.Env) generated.ResolverRoot {
	opts := []grpc.DialOption{grpc.WithInsecure()}
	conn, err := grpc.DialContext(ctx, env.GetUserServerName()+":"+env.GetUserServerPort(), opts...)
	if err != nil {
		panic(err)
	}
	defer conn.Close()
	userClient := definition.NewUserQueryServiceClient(conn)

	return &resolver{
		userClient,
	}
}
