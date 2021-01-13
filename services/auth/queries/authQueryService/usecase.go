package authqueryservice

import (
	"github.com/takuya911/project-services/services/auth/shared"
)

type (
	usecase struct {
	}
	genTokenRequest struct {
		userUUID string
	}
	genTokenResponse struct {
		token        string
		refreshToken string
	}
	// Usecase interface
	Usecase interface {
		genToken(req genTokenRequest) (genTokenResponse, error)
	}
)

// NewUsecase function
func NewUsecase() Usecase {
	return &usecase{}
}

func (uc *usecase) genToken(req genTokenRequest) (genTokenResponse, error) {
	tokenPair, err := shared.GenTokenPair(req.userUUID)
	if err != nil {
		return genTokenResponse{}, err
	}
	return genTokenResponse{
		token:        tokenPair.Token,
		refreshToken: tokenPair.RefreshToken,
	}, nil
}
