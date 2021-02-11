package domain

import (
	"errors"

	"github.com/google/uuid"
	"github.com/takuya911/ddd_gql-grpc/services/auth/shared"
)

type (
	// UserUUID vo
	UserUUID string
)

// ParseUserUUID func
func ParseUserUUID(id string) (UserUUID, error) {
	return parseUserUUID(id)
}

func parseUserUUID(id string) (UserUUID, error) {
	if id == "" {
		return "", errors.New(shared.RequiredUserUUID)
	}
	if _, err := uuid.Parse(id); err != nil {
		return "", errors.New(shared.InvalidUUIDFormat)
	}
	return UserUUID(id), nil
}
