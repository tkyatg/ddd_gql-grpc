package usercommandservice

import (
	"context"

	definition "github.com/tkyatg/user-definition"
)

type (
	server struct {
		uc Usecase
	}
)

// NewServer はコンストラクタです
func NewServer(uc Usecase) definition.UserCommandServiceServer {
	return &server{uc}
}
func (s *server) Create(context context.Context, req *definition.CreateRequest) (*definition.CreateResponse, error) {
	res, err := s.uc.create(createRequest{
		name:            req.Name,
		email:           req.Email,
		password:        req.Password,
		telephoneNumber: req.TelephoneNumber,
		gender:          req.Gender,
	})
	if err != nil {
		return nil, err
	}
	return &definition.CreateResponse{
		Uuid: res.userUUID,
	}, nil
}

func (s *server) Update(context context.Context, req *definition.UpdateRequest) (*definition.UpdateResponse, error) {
	if err := s.uc.update(updateRequest{
		userUUID:        req.Uuid,
		name:            req.Name,
		email:           req.Email,
		password:        req.Password,
		telephoneNumber: req.TelephoneNumber,
		gender:          req.Gender,
	}); err != nil {
		return nil, err
	}
	return &definition.UpdateResponse{
		Uuid: req.Uuid,
	}, nil
}

func (s *server) Delete(context context.Context, req *definition.DeleteRequest) (*definition.DeleteResponse, error) {
	if err := s.uc.delete(deleteRequest{
		userUUID: req.Uuid,
	}); err != nil {
		return &definition.DeleteResponse{}, err
	}
	return &definition.DeleteResponse{
		Uuid: req.Uuid,
	}, nil
}
