package userqueryservice

import (
	"github.com/golang/protobuf/ptypes/timestamp"
)

type (
	usecase struct {
		da DataAccessor
	}
	getUserByIDRequest struct {
		userUUID string
	}
	getUserByIDResponse struct {
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
		getUserByID(req getUserByIDRequest) (getUserByIDResponse, error)
	}
	// DataAccessor interface
	DataAccessor interface {
		getUserByID(req getUserByIDRequest) (getUserByIDResponse, error)
	}
)

// NewUsecase function
func NewUsecase(da DataAccessor) Usecase {
	return &usecase{da}
}

func (uc *usecase) getUserByID(req getUserByIDRequest) (getUserByIDResponse, error) {
	return uc.da.getUserByID(req)
}
