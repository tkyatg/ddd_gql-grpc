package domain

import (
	"errors"
	"regexp"

	"github.com/google/uuid"
	"github.com/takuya911/project-services/services/auth/shared"
)

type (
	// UserUUID vo
	UserUUID string
	// Email vo
	Email string
	// Password vo
	Password string
	// LoginResult vo
	LoginResult bool
)

// ParseUserUUID func
func ParseUserUUID(id string) (UserUUID, error) {
	return parseUserUUID(id)
}

// ParsePassword func
func ParsePassword(password string) (Password, error) {
	return parsePassword(password)
}

// ParseEmail func
func ParseEmail(email string) (Email, error) {
	return parseEmail(email)
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

func parsePassword(password string) (Password, error) {
	if password == "" {
		return "", errors.New(shared.RequiredPassword)
	}
	if len(password) < 8 {
		return "", errors.New(shared.InvalidPasswordLength)
	}
	return Password(password), nil
}

func parseEmail(email string) (Email, error) {
	if email == "" {
		return "", errors.New(shared.RequiredEmail)
	}
	if len(email) < 3 && len(email) > 254 {
		return "", errors.New(shared.InvalidEmailLength)
	}
	if m, _ := regexp.MatchString(`^([\w\.\_]{2,10})@(\w{1,}).([a-z]{2,4})$`, email); !m {
		return "", errors.New(shared.InvalidEmailFormat)
	}
	return Email(email), nil
}
