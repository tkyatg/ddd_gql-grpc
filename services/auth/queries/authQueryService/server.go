package authqueryservice

import (
	"context"
	"errors"

	"github.com/takuya911/ddd_gql-grpc/services/auth/domain"
	"github.com/takuya911/ddd_gql-grpc/services/auth/shared"
	definition "github.com/takuya911/project-auth-definition"
)

type server struct {
	uc Usecase
}

// NewServer function
func NewServer(uc Usecase) definition.AuthQueryServiceServer {
	return &server{uc}
}

func (s *server) GenToken(ctx context.Context, req *definition.GenTokenRequest) (*definition.GenTokenResponse, error) {
	uuid := req.GetUuid()
	if uuid == "" {
		return nil, errors.New(shared.RequiredUserUUID)
	}
	if _, err := domain.ParseUserUUID(uuid); err != nil {
		return nil, err
	}
	tokenPair, err := s.uc.genToken(genTokenRequest{userUUID: uuid})
	if err != nil {
		return nil, err
	}
	return &definition.GenTokenResponse{
		AccessToken:  tokenPair.accessToken,
		RefreshToken: tokenPair.refreshToken,
	}, nil
}
