package authenticationcommandservice

import (
	"github.com/takuya911/project-services/services/auth/domain"
	"github.com/takuya911/project-services/services/auth/shared"
)

type (
	usecase struct {
		repo  domain.AuthenticationRepository
		token shared.Token
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
func NewUsecase(repo domain.AuthenticationRepository, token shared.Token) Usecase {
	return &usecase{repo, token}
}

func (uc *usecase) login(req loginRequest) (loginResponse, error) {
	uuid, err := uc.repo.Login(domain.Email(req.email), domain.Password(req.password))
	if err != nil {
		return loginResponse{}, err
	}
	accessToken, refreshToken, err := uc.token.GenTokenPair(string(uuid))
	if err != nil {
		return loginResponse{}, err
	}

	return loginResponse{
		loginResult:  true,
		accessToken:  accessToken,
		refreshToken: refreshToken,
	}, nil
}
