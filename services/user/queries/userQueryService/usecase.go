package userqueryservice

import (
	"github.com/golang/protobuf/ptypes/timestamp"
)

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
		createdAt       *timestamp.Timestamp
		updatedAt       *timestamp.Timestamp
	}
	// Usecase interface
	Usecase interface {
		getByID(req getByIDRequest) (getByIDResponse, error)
	}
	// DataAccessor interface
	DataAccessor interface {
		getByID(req getByIDRequest) (getByIDResponse, error)
	}
)

// NewUsecase function
func NewUsecase(da DataAccessor) Usecase {
	return &usecase{da}
}

func (uc *usecase) getByID(req getByIDRequest) (getByIDResponse, error) {
	return uc.da.getByID(req)
}
