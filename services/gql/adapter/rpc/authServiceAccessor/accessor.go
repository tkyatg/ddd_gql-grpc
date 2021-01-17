package authserviceaccessor

import (
	"context"

	definition "github.com/takuya911/project-auth-definition"
)

type (
	serviceAccessor struct {
		authQueryClient definition.AuthQueryServiceClient
	}
	// ServiceAccessor interface
	ServiceAccessor interface {
		GenToken(ctx context.Context, req GenTokenRequest) (GenTokenResponse, error)
		Login(ctx context.Context, req LoginRequest) (LoginResponse, error)
	}
)

// NewAuthServiceAccessor func
func NewAuthServiceAccessor(authQueryClient definition.AuthQueryServiceClient) ServiceAccessor {
	return &serviceAccessor{authQueryClient}
}

func (r *serviceAccessor) GenToken(ctx context.Context, req GenTokenRequest) (GenTokenResponse, error) {
	res, err := r.authQueryClient.GenToken(ctx, &definition.GenTokenRequest{
		Uuid: req.UUID,
	})
	if err != nil {
		return GenTokenResponse{}, err
	}

	return GenTokenResponse{
		TokenPair: TokenPair{
			AccessToken:  res.GetToken(),
			RefreshToken: res.GetRefreshToken(),
		},
	}, nil
}

func (r *serviceAccessor) Login(ctx context.Context, req LoginRequest) (LoginResponse, error) {
	return LoginResponse{}, nil
}
