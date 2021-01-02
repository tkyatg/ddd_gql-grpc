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
		gender          string
	}
	createResponse struct {
	}
	updateRequest struct {
		userUUID        string
		name            string
		email           string
		password        string
		telephoneNumber string
		gender          string
	}
	updateResponse struct {
	}
	deleteRequest struct {
		userUUID string
	}
	deleteResponse struct {
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
	attr, err := domain.NewUserAttributes(req.nam, req.spotCategory, req.spotAuthenticators, req.reservationRule)
	if err != nil {
		return err
	}
	return uc.repo.create(attr)
}

func (uc *usecase) update(req updateRequest) (updateResponse, error) {
	id, err := domain.ParseUserUUID(req.userUUID)
	if err != nil {
		return err
	}
	attr, err := domain.NewUserAttributes(req.nam, req.spotCategory, req.spotAuthenticators, req.reservationRule)
	if err != nil {
		return err
	}

	return uc.repo.update(id, attr)
}

func (uc *usecase) delete(req deleteRequest) (deleteResponse, error) {
	id, err := domain.ParseUserUUID(req.userUUID)
	if err != nil {
		return err
	}
	return uc.repo.delete(id)
}
