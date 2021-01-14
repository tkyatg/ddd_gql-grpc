package authenticationcommandservice

import (
	"github.com/takuya911/project-services/services/auth/domain"
	"github.com/takuya911/project-services/services/auth/shared"
)

type (
	usecase struct {
		repo  domain.AuthenticationRepository
		token shared.Token
		hash  shared.Hash
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
func NewUsecase(repo domain.AuthenticationRepository, token shared.Token, hash shared.Hash) Usecase {
	return &usecase{repo, token, hash}
}

func (uc *usecase) login(req loginRequest) (loginResponse, error) {
	email, err := domain.ParseEmail(req.email)
	if err != nil {
		return loginResponse{}, err
	}

	uuid, password, err := uc.repo.Login(email)
	if err != nil {
		return loginResponse{}, err
	}
	if err := uc.hash.CompareHashAndPass(string(password), req.password); err != nil {
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
