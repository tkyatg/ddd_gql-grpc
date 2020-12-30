package userqueryservice

import (
	"context"

	"github.com/jinzhu/gorm"
)

type dataAccessor struct {
	Conn *gorm.DB
}

// NewDataAccessor function
func NewDataAccessor(conn *gorm.DB) DataAccessor {
	return &dataAccessor{conn}
}

func (d *dataAccessor) getUserByID(ctx context.Context, req getUserByIDRequest) (getUserByIDResponse, error) {
	return getUserByIDResponse{}, nil
}
