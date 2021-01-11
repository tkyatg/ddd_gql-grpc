package resolver

import (
	"context"

	userserviceaccessor "github.com/takuya911/project-services/services/gql/adapter/rpc/userServiceAccessor"
	"github.com/takuya911/project-services/services/gql/graph/generated"
	"github.com/takuya911/project-services/services/gql/shared"
	userdefinition "github.com/takuya911/project-user-definition"
	"google.golang.org/grpc"
)

type resolver struct {
	userServiceClient   userdefinition.UserQueryServiceClient
	userServiceAccessor userserviceaccessor.ServiceAccessor
}

// NewResolver function
func NewResolver(ctx context.Context, env shared.Env) generated.ResolverRoot {
	opts := []grpc.DialOption{grpc.WithInsecure()}
	conn, err := grpc.DialContext(ctx, env.GetUserServerName()+":"+env.GetUserServerPort(), opts...)
	if err != nil {
		panic(err)
	}
	// client
	userServiceClient := userdefinition.NewUserQueryServiceClient(conn)

	// accessor
	userServiceAccessor := userserviceaccessor.NewUserServiceAccessor(userServiceClient)

	return &resolver{
		userServiceClient,
		userServiceAccessor,
	}
}
