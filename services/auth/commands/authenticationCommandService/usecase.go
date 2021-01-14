package authenticationcommandservice

import "github.com/takuya911/project-services/services/user/domain"

type (
	usecase struct {
		repo domain.AuthenticationRepository
	}
	loginRequest struct {
		email    string
		password string
	}
	loginResponse struct {
		result       int64
		token        string
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

	return loginResponse{
		result:       1,
		token:        "",
		refreshToken: "",
	}, nil
}
