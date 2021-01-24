package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"errors"

	authserviceaccessor "github.com/takuya911/project-services/services/gql/adapter/rpc/authServiceAccessor"
	userserviceaccessor "github.com/takuya911/project-services/services/gql/adapter/rpc/userServiceAccessor"
	"github.com/takuya911/project-services/services/gql/graph/generated"
	"github.com/takuya911/project-services/services/gql/graph/model"
	"github.com/takuya911/project-services/services/gql/shared"
)

func (r *mutationResolver) Login(ctx context.Context, input model.LoginRequest) (*model.LoginResponse, error) {
	res, err := r.userServiceAccessor.GetByEmailAndPassword(ctx, userserviceaccessor.GetByEmailAndPasswordRequest{
		Email:    input.Email,
		Password: input.Password,
	})
	if err != nil {
		return nil, err
	}
	if res.UUID == "" {
		return nil, errors.New(shared.LoginFailed)
	}

	token, err := r.authServiceAccessor.GenToken(ctx, authserviceaccessor.GenTokenRequest{UUID: res.UUID})
	if err != nil {
		return nil, err
	}

	return &model.LoginResponse{
		UUID: res.UUID,
		TokenPair: &model.TokenPair{
			AccessToken:  token.TokenPair.AccessToken,
			RefreshToken: token.TokenPair.RefreshToken,
		},
	}, nil
}

// Mutation returns generated.MutationResolver implementation.
func (r *resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

type mutationResolver struct{ *resolver }
