package authserviceaccessor

import (
	"context"

	definition "github.com/takuya911/project-auth-definition"
)

type (
	serviceAccessor struct {
		authQuery             definition.AuthQueryServiceClient
		authenticationCommand definition.AuthenticationCommandServiceClient
	}
	// ServiceAccessor interface
	ServiceAccessor interface {
		GenToken(ctx context.Context, req GenTokenRequest) (GenTokenResponse, error)
		Login(ctx context.Context, req LoginRequest) (LoginResponse, error)
	}
)

// NewAuthServiceAccessor func
func NewAuthServiceAccessor(authQueryClient definition.AuthQueryServiceClient, authenticationCommandServiceClient definition.AuthenticationCommandServiceClient) ServiceAccessor {
	return &serviceAccessor{
		authQuery:             authQueryClient,
		authenticationCommand: authenticationCommandServiceClient}
}

func (r *serviceAccessor) GenToken(ctx context.Context, req GenTokenRequest) (GenTokenResponse, error) {
	res, err := r.authQuery.GenToken(ctx, &definition.GenTokenRequest{
		Uuid: req.UUID,
	})
	if err != nil {
		return GenTokenResponse{}, err
	}

	return GenTokenResponse{
		TokenPair: TokenPair{
			AccessToken:  res.GetAccessToken(),
			RefreshToken: res.GetRefreshToken(),
		},
	}, nil
}

func (r *serviceAccessor) Login(ctx context.Context, req LoginRequest) (LoginResponse, error) {
	res, err := r.authenticationCommand.Login(ctx, &definition.LoginRequest{
		Email:    req.Email,
		Password: req.Password,
	})
	if err != nil {
		return LoginResponse{}, err
	}

	return LoginResponse{
		TokenPair: TokenPair{
			AccessToken:  res.GetAccessToken(),
			RefreshToken: res.GetRefreshToken(),
		},
	}, nil
}
