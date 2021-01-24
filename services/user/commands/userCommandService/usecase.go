package usercommandservice

import (
	"github.com/takuya911/project-services/services/user/domain"
	"github.com/takuya911/project-services/services/user/shared"
)

type (
	usecase struct {
		repo domain.UserRepository
		hash shared.Hash
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
func NewUsecase(repo domain.UserRepository, hash shared.Hash) Usecase {
	return &usecase{repo, hash}
}

func (uc *usecase) create(req createRequest) (createResponse, error) {
	hashedPassword, err := uc.hash.GenEncryptedPassword(req.password)
	if err != nil {
		return createResponse{}, err
	}

	attr, err := domain.NewUserAttributes(req.name, hashedPassword, req.email, req.telephoneNumber, req.gender)
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
	hashedPassword, err := uc.hash.GenEncryptedPassword(req.password)
	if err != nil {
		return err
	}
	attr, err := domain.NewUserAttributes(req.name, hashedPassword, req.email, req.telephoneNumber, req.gender)
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
