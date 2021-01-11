package userserviceaccessor

import (
	"context"
	"fmt"

	"github.com/takuya911/project-services/services/gql/graph/model"
	definition "github.com/takuya911/project-user-definition"
)

type (
	serviceAccessor struct {
		userQueryClient   definition.UserQueryServiceClient
		userCommnadClient definition.UserCommandServiceClient
	}
	// ServiceAccessor interface
	ServiceAccessor interface {
		GetByID(ctx context.Context, uuid string) (*model.GetUserByIDResponse, error)
		Create(ctx context.Context, req CreateUserRequest) (string, error)
	}
)

// NewUserServiceAccessor func
func NewUserServiceAccessor(userQueryClient definition.UserQueryServiceClient, userCommnadClient definition.UserCommandServiceClient) ServiceAccessor {
	return &serviceAccessor{userQueryClient, userCommnadClient}
}

func (r *serviceAccessor) GetByID(ctx context.Context, uuid string) (*model.GetUserByIDResponse, error) {
	res, err := r.userQueryClient.GetByID(ctx, &definition.GetByIDRequest{
		Uuid: uuid,
	})
	fmt.Print(res)
	if err != nil {
		return nil, err
	}

	return &model.GetUserByIDResponse{User: &model.User{
		UUID:            res.GetUuid(),
		Name:            res.GetName(),
		Email:           res.GetEmail(),
		Password:        res.GetPassword(),
		TelephoneNumber: res.GetTelephoneNumber(),
		Gender:          res.GetGender(),
		CreatedAt:       res.GetCreatedAt(),
		UpdatedAt:       res.GetUpdatedAt(),
	}}, nil
}

func (r *serviceAccessor) Create(ctx context.Context, req CreateUserRequest) (string, error) {
	res, err := r.userCommnadClient.Create(ctx, &definition.CreateRequest{
		Name:            req.Name,
		Email:           req.Email,
		Password:        req.Password,
		TelephoneNumber: req.TelephoneNumber,
		Gender:          req.Gender,
	})
	if err != nil {
		return "", err
	}
	return res.Uuid, err
}
