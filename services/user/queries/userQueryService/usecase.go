package userqueryservice

import (
	"time"

	"github.com/tkyatg/ddd_gql-grpc/services/user/shared"
)

type (
	usecase struct {
		da   DataAccessor
		hash shared.Hash
	}
	getByIDRequest struct {
		userUUID string
	}
	getByIDResponse struct {
		userUUID        string
		name            string
		email           string
		password        string
		telephoneNumber string
		gender          int64
		createdAt       time.Time
		updatedAt       time.Time
	}
	getByEmailAndPasswordRequest struct {
		email    string
		password string
	}
	getByEmailAndPasswordResponse struct {
		userUUID        string
		name            string
		email           string
		password        string
		telephoneNumber string
		gender          int64
		createdAt       time.Time
		updatedAt       time.Time
	}
	// Usecase interface
	Usecase interface {
		getByID(req getByIDRequest) (getByIDResponse, error)
		getByEmailAndPassword(req getByEmailAndPasswordRequest) (getByEmailAndPasswordResponse, error)
	}
	// DataAccessor interface
	DataAccessor interface {
		getByID(req getByIDRequest) (getByIDResponse, error)
		getByEmail(req getByEmailAndPasswordRequest) (getByEmailAndPasswordResponse, error)
	}
)

// NewUsecase function
func NewUsecase(da DataAccessor, hash shared.Hash) Usecase {
	return &usecase{da, hash}
}

func (uc *usecase) getByID(req getByIDRequest) (getByIDResponse, error) {
	return uc.da.getByID(req)
}

func (uc *usecase) getByEmailAndPassword(req getByEmailAndPasswordRequest) (getByEmailAndPasswordResponse, error) {
	res, err := uc.da.getByEmail(req)
	if err != nil {
		return getByEmailAndPasswordResponse{}, err
	}
	if err := uc.hash.CompareHashAndPass(res.password, req.password); err != nil {
		return getByEmailAndPasswordResponse{}, err
	}

	return res, nil
}
