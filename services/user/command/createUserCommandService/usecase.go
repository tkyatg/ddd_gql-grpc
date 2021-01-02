package createusercommandservice

import "github.com/takuya911/project-services/services/user/domain"

type (
	usecase struct {
		repo domain.UserRepository
	}
	createRequest struct {
		name            string
		email           string
		password        string
		telephoneNumber string
		gender          int64
	}
	createResponse struct {
		userUUID  string
		tokenPair *tokenPair
	}
	tokenPair struct {
		accessToken  string
		refreshToken string
	}
	updateRequest struct {
		userUUID        string
		name            string
		email           string
		password        string
		telephoneNumber string
		gender          int64
	}
	updateResponse struct {
		userUUID  string
		tokenPair *tokenPair
	}
	deleteRequest struct {
		userUUID string
	}
	deleteResponse struct {
		userUUID string
	}
	// Usecase interface
	Usecase interface {
		create(req createRequest) (createResponse, error)
		update(req updateRequest) (updateResponse, error)
		delete(req deleteRequest) (deleteResponse, error)
	}
)

// NewUsecase はコンストラクタです
func NewUsecase(repo domain.UserRepository) Usecase {
	return &usecase{repo}
}

func (uc *usecase) create(req createRequest) (createResponse, error) {
	attr, err := domain.NewUserAttributes(req.name, req.password, req.email, req.telephoneNumber, req.gender)
	if err != nil {
		return createResponse{}, err
	}
	if err := uc.repo.create(attr); err != nil {
		return createResponse{}, err
	}
	return createResponse{}, nil
}

func (uc *usecase) update(req updateRequest) (updateResponse, error) {
	id, err := domain.ParseUserUUID(req.userUUID)
	if err != nil {
		return updateResponse{}, err
	}
	attr, err := domain.NewUserAttributes(req.name, req.password, req.email, req.telephoneNumber, req.gender)
	if err != nil {
		return updateResponse{}, err
	}
	if err := uc.repo.update(id, attr); err != nil {
		return updateResponse{}, err
	}

	return updateResponse{}, nil
}

func (uc *usecase) delete(req deleteRequest) (deleteResponse, error) {
	id, err := domain.ParseUserUUID(req.userUUID)
	if err != nil {
		return deleteResponse{}, err
	}
	if err := uc.repo.delete(id); err != nil {
		return deleteResponse{}, err
	}
	return deleteResponse{}, nil

}
