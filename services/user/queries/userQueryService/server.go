package userqueryservice

import (
	"context"
	"errors"

	"github.com/takuya911/project-services/services/user/shared"
	definition "github.com/takuya911/project-user-definition"
)

type server struct {
	uc Usecase
}

// NewServer function
func NewServer(uc Usecase) definition.UserServiceServer {
	return &server{uc}
}

func (s *server) GetUserByID(ctx context.Context, req *definition.GetUserRequest) (*definition.GetUserResponse, error) {
	uuid := req.GetUuid()
	if uuid == "" {
		return nil, errors.New(shared.RequiredUserID)
	}
	res, err := s.uc.getUserByID(getUserByIDRequest{userUUID: req.GetUuid()})
	if err != nil {
		return nil, err
	}

	return &definition.GetUserResponse{
		Uuid:            res.userUUID,
		Name:            res.name,
		Email:           res.email,
		Password:        res.password,
		TelephoneNumber: res.telephoneNumber,
		Gender:          res.gender,
		CreatedAt:       res.createdAt,
		UpdatedAt:       res.updatedAt,
	}, nil
}
