package graph

import (
	"github.com/takuya911/project-services/services/gql/graph/generated"
	definition "github.com/takuya911/project-user-definition"
)

type resolver struct {
	userClient definition.UserQueryServiceClient
}

// NewResolver function
func NewResolver(userClient definition.UserQueryServiceClient) generated.ResolverRoot {
	return &resolver{
		userClient,
	}
}
