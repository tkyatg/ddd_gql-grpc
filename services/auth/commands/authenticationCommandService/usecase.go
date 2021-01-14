package authenticationcommandservice

import (
	"github.com/takuya911/project-services/services/auth/domain"
	"github.com/takuya911/project-services/services/auth/shared"
)

type (
	usecase struct {
		repo domain.AuthenticationRepository
	}
	loginRequest struct {
		email    string
		password string
	}
	loginResponse struct {
		loginResult  bool
		accessToken  string
		refreshToken string
	}
	// Usecase interface
	Usecase interface {
		login(req loginRequest) (loginResponse, error)
	}
)

// NewUsecase はコンストラクタです
func NewUsecase(repo domain.AuthenticationRepository) Usecase {
	return &usecase{repo}
}

func (uc *usecase) login(req loginRequest) (loginResponse, error) {
	uuid, err := uc.repo.Login(domain.Email(req.email), domain.Password(req.password))
	if err != nil {
		return loginResponse{}, err
	}
	tokenPair, err := shared.GenTokenPair(string(uuid))
	if err != nil {
		return loginResponse{}, err
	}

	return loginResponse{
		loginResult:  true,
		accessToken:  tokenPair.AccessToken,
		refreshToken: tokenPair.RefreshToken,
	}, nil
}
