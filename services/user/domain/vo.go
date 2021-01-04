package domain

import (
	"errors"
	"regexp"
	"strconv"

	"github.com/google/uuid"
	"github.com/takuya911/project-services/services/user/shared"
)

type (
	// UserUUID vo
	UserUUID string
	// UserName vo
	UserName string
	// Email vo
	Email string
	// Password vo
	Password string
	// TelephoneNumber vo
	TelephoneNumber string
	// Gender vo
	Gender int
	// UserAttributes struct
	UserAttributes struct {
		name            UserName
		email           Email
		password        Password
		telephoneNumber TelephoneNumber
		gender          Gender
	}
	// TokenPair struct
	TokenPair struct {
		AccessToken  string
		RefreshToken string
	}
)

// NewUserAttributes func
func NewUserAttributes(
	name string,
	password string,
	email string,
	telephoneNumber string,
	gender int64,
) (UserAttributes, error) {
	nm, err := parseUserName(name)
	if err != nil {
		return UserAttributes{}, err
	}
	pass, err := parsePassword(password)
	if err != nil {
		return UserAttributes{}, err
	}
	em, err := parseEmail(email)
	if err != nil {
		return UserAttributes{}, err
	}
	tl, err := parseTelephoneNumber(telephoneNumber)
	if err != nil {
		return UserAttributes{}, err
	}
	gn, err := parseGender(gender)
	if err != nil {
		return UserAttributes{}, err
	}

	return UserAttributes{
		name:            nm,
		email:           em,
		password:        pass,
		telephoneNumber: tl,
		gender:          gn,
	}, nil
}

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
func parseUserName(userName string) (UserName, error) {
	if userName == "" {
		return "", errors.New(shared.RequiredUserName)
	}
	return UserName(userName), nil
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
func parseTelephoneNumber(telephoneNumber string) (TelephoneNumber, error) {
	if telephoneNumber == "" {
		return "", errors.New(shared.RequiredTelephoneNumber)
	}
	if _, err := strconv.Atoi(telephoneNumber); err != nil {
		return "", errors.New(shared.InvalidTelephoneNumberFormat)
	}
	return TelephoneNumber(telephoneNumber), nil
}
func parseGender(gender int64) (Gender, error) {
	if gender < 1 || gender > 3 {
		return 0, errors.New(shared.InvalidGenderFormat)
	}
	return Gender(gender), nil
}
