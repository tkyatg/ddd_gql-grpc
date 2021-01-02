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
		userUUID string
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
	}
	deleteRequest struct {
		userUUID string
	}
	deleteResponse struct {
	}
	// Usecase interface
	Usecase interface {
		create(req createRequest) (createResponse, error)
		update(req updateRequest) error
		delete(req deleteRequest) error
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
	userUUID, err := uc.repo.Create(attr)
	if err != nil {
		return createResponse{}, err
	}

	return createResponse{
		userUUID: string(userUUID),
	}, nil
}

func (uc *usecase) update(req updateRequest) error {
	id, err := domain.ParseUserUUID(req.userUUID)
	if err != nil {
		return err
	}
	attr, err := domain.NewUserAttributes(req.name, req.password, req.email, req.telephoneNumber, req.gender)
	if err != nil {
		return err
	}
	if err := uc.repo.Update(id, attr); err != nil {
		return err
	}

	return nil
}

func (uc *usecase) delete(req deleteRequest) error {
	id, err := domain.ParseUserUUID(req.userUUID)
	if err != nil {
		return err
	}
	if err := uc.repo.Delete(id); err != nil {
		return err
	}
	return nil
}
