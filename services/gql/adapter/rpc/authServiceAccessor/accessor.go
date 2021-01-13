package authserviceaccessor

import (
	"context"

	definition "github.com/takuya911/project-auth-definition"
	"github.com/takuya911/project-services/services/gql/graph/model"
)

type (
	serviceAccessor struct {
		authQueryClient definition.AuthQueryServiceClient
	}
	// ServiceAccessor interface
	ServiceAccessor interface {
		GenToken(ctx context.Context, req GenTokenRequest) (*model.TokenPair, error)
	}
)

// NewAuthServiceAccessor func
func NewAuthServiceAccessor(authQueryClient definition.AuthQueryServiceClient) ServiceAccessor {
	return &serviceAccessor{authQueryClient}
}

func (r *serviceAccessor) GenToken(ctx context.Context, req GenTokenRequest) (*model.TokenPair, error) {
	res, err := r.authQueryClient.GenToken(ctx, &definition.GenTokenRequest{
		Uuid: req.UUID,
	})
	if err != nil {
		return nil, err
	}

	return &model.TokenPair{
		Token:        res.GetToken(),
		RefreshToken: res.GetRefreshToken(),
	}, nil
}
