package userqueryservice

import (
	"context"

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
	res, err := s.uc.getUserByID(ctx, getUserByIDRequest{id: req.GetId()})
	if err != nil {
		return nil, err
	}

	return &definition.GetUserResponse{
		Id:              res.id,
		Name:            res.name,
		Email:           res.email,
		Password:        res.password,
		TelephoneNumber: res.telephoneNumber,
		Gender:          res.gender,
		CreatedAt:       res.createdAt,
		UpdatedAt:       res.updatedAt,
	}, nil
}
