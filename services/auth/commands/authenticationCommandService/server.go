package authenticationcommandservice

import (
	"context"

	definition "github.com/takuya911/project-auth-definition"
)

type (
	server struct {
		uc Usecase
	}
)

// NewServer はコンストラクタです
func NewServer(uc Usecase) definition.AuthCommandServiceServer {
	return &server{uc}
}
func (s *server) Login(context context.Context, req *definition.LoginRequest) (*definition.LoginResponse, error) {
	res, err := s.uc.login(loginRequest{
		email:    req.Email,
		password: req.Password,
	})
	if err != nil {
		return nil, err
	}

	return &definition.LoginResponse{
		LoginResult:  res.loginResult,
		AccessToken:  res.accessToken,
		RefreshToken: res.refreshToken,
	}, nil
}
