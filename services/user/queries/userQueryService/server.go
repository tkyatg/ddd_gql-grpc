package userqueryservice

import (
	"context"
	"errors"

	"github.com/takuya911/project-services/services/user/domain"
	"github.com/takuya911/project-services/services/user/shared"
	definition "github.com/takuya911/project-user-definition"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type server struct {
	uc Usecase
}

// NewServer function
func NewServer(uc Usecase) definition.UserQueryServiceServer {
	return &server{uc}
}

func (s *server) GetByID(ctx context.Context, req *definition.GetByIDRequest) (*definition.GetByIDResponse, error) {
	uuid := req.GetUuid()
	if uuid == "" {
		return nil, errors.New(shared.RequiredUserUUID)
	}
	if _, err := domain.ParseUserUUID(uuid); err != nil {
		return nil, err
	}
	res, err := s.uc.getByID(getByIDRequest{userUUID: uuid})
	if err != nil {
		return nil, err
	}
	return &definition.GetByIDResponse{
		Uuid:            res.userUUID,
		Name:            res.name,
		Email:           res.email,
		Password:        res.password,
		TelephoneNumber: res.telephoneNumber,
		Gender:          res.gender,
		CreatedAt:       timestamppb.New(res.createdAt),
		UpdatedAt:       timestamppb.New(res.updatedAt),
	}, nil
}

func (s *server) GetByEmailAndPassword(ctx context.Context, req *definition.GetByEmailAndPasswordRequest) (*definition.GetByEmailAndPasswordResponse, error) {
	email := req.GetEmail()
	password := req.GetPassword()
	if email == "" || password == "" {
		return nil, errors.New(shared.RequiredLoginArgs)
	}
	res, err := s.uc.getByEmailAndPassword(getByEmailAndPasswordRequest{email: email, password: password})
	if err != nil {
		return nil, err
	}
	return &definition.GetByEmailAndPasswordResponse{
		Uuid:            res.userUUID,
		Name:            res.name,
		Email:           res.email,
		Password:        res.password,
		TelephoneNumber: res.telephoneNumber,
		Gender:          res.gender,
		CreatedAt:       timestamppb.New(res.createdAt),
		UpdatedAt:       timestamppb.New(res.updatedAt),
	}, nil
}
