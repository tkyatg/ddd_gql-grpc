package userserviceaccessor

import (
	"context"

	"github.com/takuya911/project-services/services/gql/graph/model"
	definition "github.com/takuya911/project-user-definition"
)

type (
	serviceAccessor struct {
		userClient definition.UserQueryServiceClient
	}
	// ServiceAccessor interface
	ServiceAccessor interface {
		GetByID(ctx context.Context, id string) (*model.GetUserByIDResponse, error)
	}
)

// NewUserServiceAccessor func
func NewUserServiceAccessor(userClient definition.UserQueryServiceClient) ServiceAccessor {
	return &serviceAccessor{userClient}
}

func (r *serviceAccessor) GetByID(ctx context.Context, id string) (*model.GetUserByIDResponse, error) {
	res, err := r.userClient.GetByID(ctx, &definition.GetByIDRequest{
		Uuid: id,
	})
	if err != nil {
		return nil, err
	}

	return &model.GetUserByIDResponse{User: &model.User{
		ID:              res.GetUuid(),
		Name:            res.GetName(),
		Email:           res.GetEmail(),
		Password:        res.GetPassword(),
		TelephoneNumber: res.GetTelephoneNumber(),
		Gender:          res.GetGender(),
		CreatedAt:       res.GetCreatedAt(),
		UpdatedAt:       res.GetUpdatedAt(),
	}}, nil
}
