package domain

import "github.com/jinzhu/gorm"

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

func (d userDataAccessor) create(attr *UserAttributes) error {
	return nil
}
func (d userDataAccessor) update(id UserUUID, attr *UserAttributes) error {
	return nil
}
func (d userDataAccessor) delete(id UserUUID) error {
	return nil
}
