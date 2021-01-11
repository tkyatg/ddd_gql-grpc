package userqueryservice

import (
	"context"
	"errors"

	"github.com/takuya911/project-services/services/user/domain"
	"github.com/takuya911/project-services/services/user/shared"
	definition "github.com/takuya911/project-user-definition"
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
		return nil, errors.New(shared.RequiredUserID)
	}
	if _, err := domain.ParseUserUUID(uuid); err != nil {
		return nil, err
	}
	res, err := s.uc.getByID(getUserByIDRequest{userUUID: req.GetUuid()})
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
		CreatedAt:       res.createdAt,
		UpdatedAt:       res.updatedAt,
	}, nil
}
