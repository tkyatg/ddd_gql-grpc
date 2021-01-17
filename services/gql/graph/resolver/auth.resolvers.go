package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"

	"github.com/takuya911/project-services/services/gql/graph/generated"
	"github.com/takuya911/project-services/services/gql/graph/model"
)

func (r *mutationResolver) Login(ctx context.Context, input model.LoginRequest) (*model.LoginResponse, error) {
	return &model.LoginResponse{}, nil
}

// Mutation returns generated.MutationResolver implementation.
func (r *resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

type mutationResolver struct{ *resolver }
