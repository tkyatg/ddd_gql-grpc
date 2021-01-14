package domain

import (
	"github.com/jinzhu/gorm"
)

type (
	userDataAccessor struct {
		db *gorm.DB
	}
)

// NewUserDataAccessor fun
func NewUserDataAccessor(
	db *gorm.DB,
) UserDataAccessor {
	return &userDataAccessor{db}
}

func (d userDataAccessor) login(attr string) error {
	return nil
}
