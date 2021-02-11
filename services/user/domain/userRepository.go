package domain

import (
	"errors"

	"github.com/takuya911/ddd_gql-grpc/services/user/shared"
)

type (
	userRepository struct {
		da UserDataAccessor
	}
	// UserRepository interface
	UserRepository interface {
		Create(attr UserAttributes) (UserUUID, error)
		Update(id UserUUID, attr UserAttributes) error
		Delete(id UserUUID) error
	}
	// UserDataAccessor interface
	UserDataAccessor interface {
		create(attr UserAttributes) (UserUUID, error)
		update(id UserUUID, attr UserAttributes) error
		delete(id UserUUID) error
		emailAlreadyUsedCreate(email Email) (bool, error)
		emailAlreadyUsedUpdate(id UserUUID, email Email) (bool, error)
	}
)

// NewUserRepository func
func NewUserRepository(
	da UserDataAccessor,
) UserRepository {
	return &userRepository{da}
}

func (r *userRepository) Create(attr UserAttributes) (UserUUID, error) {
	used, err := r.da.emailAlreadyUsedCreate(attr.email)
	if err != nil {
		return "", err
	}
	if used {
		return "", errors.New(shared.EmailAlreadyUsed)
	}
	return r.da.create(attr)
}

func (r *userRepository) Update(id UserUUID, attr UserAttributes) error {
	used, err := r.da.emailAlreadyUsedUpdate(id, attr.email)
	if err != nil {
		return err
	}
	if used {
		return errors.New(shared.EmailAlreadyUsed)
	}

	return r.da.update(id, attr)
}

func (r *userRepository) Delete(id UserUUID) error {
	return r.da.delete(id)
}
