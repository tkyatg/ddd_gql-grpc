package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"

	userserviceaccessor "github.com/takuya911/ddd_gql-grpc/services/gql/adapter/rpc/userServiceAccessor"
	"github.com/takuya911/ddd_gql-grpc/services/gql/graph/generated"
	"github.com/takuya911/ddd_gql-grpc/services/gql/graph/model"
)

func (r *mutationResolver) CreateUser(ctx context.Context, input model.CreateUserRequest) (*model.CreateUserResponse, error) {
	res, err := r.userServiceAccessor.Create(ctx, userserviceaccessor.CreateUserRequest{
		Name:            input.Name,
		Email:           input.Email,
		Password:        input.Password,
		TelephoneNumber: input.TelephoneNumber,
		Gender:          input.Gender,
	})
	if err != nil {
		return nil, err
	}

	// token, err := r.authServiceAccessor.GenToken(ctx, authserviceaccessor.GenTokenRequest{UUID: res.UUID})
	// if err != nil {
	// 	return nil, err
	// }

	return &model.CreateUserResponse{
		UUID:      res.UUID,
		TokenPair: &model.TokenPair{
			// AccessToken:  token.TokenPair.AccessToken,
			// RefreshToken: token.TokenPair.RefreshToken,
		},
	}, nil
}

func (r *mutationResolver) UpdateUser(ctx context.Context, input model.UpdateUserRequest) (*model.UpdateUserResponse, error) {
	res, err := r.userServiceAccessor.Update(ctx, userserviceaccessor.UpdateUserRequest{
		UUID:            input.UUID,
		Name:            input.Name,
		Email:           input.Email,
		Password:        input.Password,
		TelephoneNumber: input.TelephoneNumber,
		Gender:          input.Gender,
	})
	if err != nil {
		return nil, err
	}

	// token, err := r.authServiceAccessor.GenToken(ctx, authserviceaccessor.GenTokenRequest{UUID: res.UUID})
	// if err != nil {
	// 	return nil, err
	// }

	return &model.UpdateUserResponse{
		UUID:      res.UUID,
		TokenPair: &model.TokenPair{
			// AccessToken:  token.TokenPair.AccessToken,
			// RefreshToken: token.TokenPair.RefreshToken,
		},
	}, nil
}

func (r *mutationResolver) DeleteUser(ctx context.Context, input model.DeleteUserRequest) (*model.DeleteUserResponse, error) {
	res, err := r.userServiceAccessor.Delete(ctx, userserviceaccessor.DeleteUserRequest{
		UUID: input.UUID,
	})
	if err != nil {
		return nil, err
	}

	return &model.DeleteUserResponse{
		UUID: res.UUID,
	}, nil
}

func (r *queryResolver) GetUserByID(ctx context.Context, input model.GetUserByIDRequest) (*model.GetUserByIDResponse, error) {
	res, err := r.userServiceAccessor.GetByID(ctx, userserviceaccessor.GetByIDRequest{
		UUID: input.UUID,
	})
	if err != nil {
		return nil, err
	}

	return &model.GetUserByIDResponse{
		User: &model.User{
			UUID:            res.UUID,
			Name:            res.Name,
			Email:           res.Email,
			Password:        res.Password,
			TelephoneNumber: res.TelephoneNumber,
			Gender:          res.Gender,
			CreatedAt:       res.CreatedAt,
			UpdatedAt:       res.UpdatedAt,
		},
	}, nil
}

// Query returns generated.QueryResolver implementation.
func (r *resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type queryResolver struct{ *resolver }
