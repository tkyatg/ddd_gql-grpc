package authserviceaccessor

import (
	"context"

	definition "github.com/takuya911/project-auth-definition"
)

type (
	serviceAccessor struct {
		authQuery definition.AuthQueryServiceClient
	}
	// ServiceAccessor interface
	ServiceAccessor interface {
		GenToken(ctx context.Context, req GenTokenRequest) (GenTokenResponse, error)
	}
)

// NewAuthServiceAccessor func
func NewAuthServiceAccessor(authQueryClient definition.AuthQueryServiceClient) ServiceAccessor {
	return &serviceAccessor{
		authQuery: authQueryClient,
	}
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
