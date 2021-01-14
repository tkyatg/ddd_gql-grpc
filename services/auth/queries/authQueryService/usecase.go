package authqueryservice

import (
	"github.com/takuya911/project-services/services/auth/shared"
)

type (
	usecase struct {
		token shared.Token
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
func NewUsecase(token shared.Token) Usecase {
	return &usecase{token}
}

func (uc *usecase) genToken(req genTokenRequest) (genTokenResponse, error) {
	accessToken, refreshToken, err := uc.token.GenTokenPair(req.userUUID)
	if err != nil {
		return genTokenResponse{}, err
	}
	return genTokenResponse{
		token:        accessToken,
		refreshToken: refreshToken,
	}, nil
}
