package userserviceaccessor

import (
	"context"

	definition "github.com/tkyatg/user-definition"
)

type (
	serviceAccessor struct {
		userQueryClient   definition.UserQueryServiceClient
		userCommnadClient definition.UserCommandServiceClient
	}
	// ServiceAccessor interface
	ServiceAccessor interface {
		GetByID(ctx context.Context, req GetByIDRequest) (GetByIDResponse, error)
		GetByEmailAndPassword(ctx context.Context, req GetByEmailAndPasswordRequest) (GetByEmailAndPasswordResponse, error)
		Create(ctx context.Context, req CreateUserRequest) (CreateUserResponse, error)
		Update(ctx context.Context, req UpdateUserRequest) (UpdateUserResponse, error)
		Delete(ctx context.Context, req DeleteUserRequest) (DeleteUserResponse, error)
	}
)

// NewUserServiceAccessor func
func NewUserServiceAccessor(
	userQueryClient definition.UserQueryServiceClient,
	userCommnadClient definition.UserCommandServiceClient) ServiceAccessor {
	return &serviceAccessor{userQueryClient, userCommnadClient}
}

func (r *serviceAccessor) GetByID(ctx context.Context, req GetByIDRequest) (GetByIDResponse, error) {
	res, err := r.userQueryClient.GetByID(ctx, &definition.GetByIDRequest{
		Uuid: req.UUID,
	})
	if err != nil {
		return GetByIDResponse{}, err
	}

	return GetByIDResponse{
		UUID:            res.GetUuid(),
		Name:            res.GetName(),
		Email:           res.GetEmail(),
		Password:        res.GetPassword(),
		TelephoneNumber: res.GetTelephoneNumber(),
		Gender:          res.GetGender(),
		CreatedAt:       res.GetCreatedAt(),
		UpdatedAt:       res.GetUpdatedAt(),
	}, nil
}

func (r *serviceAccessor) GetByEmailAndPassword(ctx context.Context, req GetByEmailAndPasswordRequest) (GetByEmailAndPasswordResponse, error) {

	res, err := r.userQueryClient.GetByEmailAndPassword(ctx, &definition.GetByEmailAndPasswordRequest{
		Email:    req.Email,
		Password: req.Password,
	})
	if err != nil {
		return GetByEmailAndPasswordResponse{}, err
	}

	return GetByEmailAndPasswordResponse{
		UUID:            res.GetUuid(),
		Name:            res.GetName(),
		Email:           res.GetEmail(),
		Password:        res.GetPassword(),
		TelephoneNumber: res.GetTelephoneNumber(),
		Gender:          res.GetGender(),
		CreatedAt:       res.GetCreatedAt(),
		UpdatedAt:       res.GetUpdatedAt(),
	}, nil
}

func (r *serviceAccessor) Create(ctx context.Context, req CreateUserRequest) (CreateUserResponse, error) {
	res, err := r.userCommnadClient.Create(ctx, &definition.CreateRequest{
		Name:            req.Name,
		Email:           req.Email,
		Password:        req.Password,
		TelephoneNumber: req.TelephoneNumber,
		Gender:          req.Gender,
	})
	if err != nil {
		return CreateUserResponse{}, err
	}
	return CreateUserResponse{
		UUID: res.Uuid,
	}, err
}

func (r *serviceAccessor) Update(ctx context.Context, req UpdateUserRequest) (UpdateUserResponse, error) {
	res, err := r.userCommnadClient.Update(ctx, &definition.UpdateRequest{
		Uuid:            req.UUID,
		Name:            req.Name,
		Email:           req.Email,
		Password:        req.Password,
		TelephoneNumber: req.TelephoneNumber,
		Gender:          req.Gender,
	})
	if err != nil {
		return UpdateUserResponse{}, err
	}
	return UpdateUserResponse{
		UUID: res.Uuid,
	}, err
}

func (r *serviceAccessor) Delete(ctx context.Context, req DeleteUserRequest) (DeleteUserResponse, error) {
	res, err := r.userCommnadClient.Delete(ctx, &definition.DeleteRequest{
		Uuid: req.UUID,
	})
	if err != nil {
		return DeleteUserResponse{}, err
	}
	return DeleteUserResponse{
		UUID: res.Uuid,
	}, err
}
