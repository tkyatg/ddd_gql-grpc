package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"

	authserviceaccessor "github.com/takuya911/project-services/services/gql/adapter/rpc/authServiceAccessor"
	"github.com/takuya911/project-services/services/gql/graph/generated"
	"github.com/takuya911/project-services/services/gql/graph/model"
)

func (r *mutationResolver) Login(ctx context.Context, input model.LoginRequest) (*model.LoginResponse, error) {
	res, err := r.authServiceAccessor.Login(ctx, authserviceaccessor.LoginRequest{
		Email:    input.Email,
		Password: input.Password,
	})
	if err != nil {
		return nil, err
	}
	return &model.LoginResponse{
		TokenPair: &model.TokenPair{
			AccessToken:  res.TokenPair.AccessToken,
			RefreshToken: res.TokenPair.RefreshToken,
		},
	}, nil
}

// Mutation returns generated.MutationResolver implementation.
func (r *resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

type mutationResolver struct{ *resolver }
