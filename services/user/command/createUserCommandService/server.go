package createusercommandservice

import (
	"context"

	definition "github.com/takuya911/project-user-definition"
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
	return &definition.CreateResponse{}, nil
}

func (s *server) Update(context context.Context, req *definition.UpdateRequest) (*definition.UpdateResponse, error) {
	return &definition.UpdateResponse{}, nil
}

func (s *server) Delete(context context.Context, req *definition.DeleteRequest) (*definition.DeleteResponse, error) {
	return &definition.DeleteResponse{}, nil
}
