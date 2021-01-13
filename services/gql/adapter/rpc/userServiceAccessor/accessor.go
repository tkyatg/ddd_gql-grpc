package userserviceaccessor

import (
	"context"

	authdefinition "github.com/takuya911/project-auth-definition"
	"github.com/takuya911/project-services/services/gql/graph/model"
	userdefinition "github.com/takuya911/project-user-definition"
)

type (
	serviceAccessor struct {
		userQueryClient   userdefinition.UserQueryServiceClient
		userCommnadClient userdefinition.UserCommandServiceClient
		authdefinition    authdefinition.AuthQueryServiceClient
	}
	// ServiceAccessor interface
	ServiceAccessor interface {
		GetByID(ctx context.Context, uuid string) (*model.GetUserByIDResponse, error)
		Create(ctx context.Context, req CreateUserRequest) (string, error)
		Update(ctx context.Context, req UpdateUserRequest) (string, error)
		Delete(ctx context.Context, req DeleteUserRequest) (string, error)
	}
)

// NewUserServiceAccessor func
func NewUserServiceAccessor(
	userQueryClient userdefinition.UserQueryServiceClient,
	userCommnadClient userdefinition.UserCommandServiceClient,
	authQueryClient authdefinition.AuthQueryServiceClient) ServiceAccessor {
	return &serviceAccessor{userQueryClient, userCommnadClient, authQueryClient}
}

func (r *serviceAccessor) GetByID(ctx context.Context, uuid string) (*model.GetUserByIDResponse, error) {
	res, err := r.userQueryClient.GetByID(ctx, &userdefinition.GetByIDRequest{
		Uuid: uuid,
	})
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
	res, err := r.userCommnadClient.Create(ctx, &userdefinition.CreateRequest{
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

func (r *serviceAccessor) Update(ctx context.Context, req UpdateUserRequest) (string, error) {
	res, err := r.userCommnadClient.Update(ctx, &userdefinition.UpdateRequest{
		Uuid:            req.UUID,
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

func (r *serviceAccessor) Delete(ctx context.Context, req DeleteUserRequest) (string, error) {
	res, err := r.userCommnadClient.Delete(ctx, &userdefinition.DeleteRequest{
		Uuid: req.UUID,
	})
	if err != nil {
		return "", err
	}
	return res.Uuid, err
}
