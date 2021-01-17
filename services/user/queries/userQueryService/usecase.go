package userqueryservice

import "time"

type (
	usecase struct {
		da DataAccessor
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
		getByEmailAndPassword(req getByEmailAndPasswordRequest) (getByEmailAndPasswordResponse, error)
	}
)

// NewUsecase function
func NewUsecase(da DataAccessor) Usecase {
	return &usecase{da}
}

func (uc *usecase) getByID(req getByIDRequest) (getByIDResponse, error) {
	return uc.da.getByID(req)
}

func (uc *usecase) getByEmailAndPassword(req getByEmailAndPasswordRequest) (getByEmailAndPasswordResponse, error) {
	return uc.da.getByEmailAndPassword(req)
}
