package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"

	userserviceaccessor "github.com/takuya911/project-services/services/gql/adapter/rpc/userServiceAccessor"
	"github.com/takuya911/project-services/services/gql/graph/generated"
	"github.com/takuya911/project-services/services/gql/graph/model"
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

	// feature: return auth info
	return &model.CreateUserResponse{
		UUID: res,
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

	// feature: return auth info
	return &model.UpdateUserResponse{
		UUID: res,
	}, nil
}

func (r *queryResolver) GetUserByID(ctx context.Context, input model.GetUserByIDRequest) (*model.GetUserByIDResponse, error) {
	return r.userServiceAccessor.GetByID(ctx, input.UUID)
}

// Mutation returns generated.MutationResolver implementation.
func (r *resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *resolver }
type queryResolver struct{ *resolver }
