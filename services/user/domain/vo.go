package domain

import (
	"errors"

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
)

// NewUserAttributes func
func NewUserAttributes(
	name string,
	password string,
	email string,
	telephoneNumber string,
	gender int64,
) (*UserAttributes, error) {
	nm, err := parseUserName(name)
	if err != nil {
		return nil, err
	}
	pass, err := parsePassword(password)
	if err != nil {
		return nil, err
	}
	em, err := parseEmail(email)
	if err != nil {
		return nil, err
	}
	tl, err := parseTelephoneNumber(telephoneNumber)
	if err != nil {
		return nil, err
	}
	gn, err := parseGender(gender)
	if err != nil {
		return nil, err
	}

	return &UserAttributes{
		name:            nm,
		email:           em,
		password:        pass,
		telephoneNumber: tl,
		gender:          gn,
	}, nil
}
func parseUserUUID(id string) (UserUUID, error) {
	if id == "" {
		return "", errors.New(shared.RequiredUserUUID)
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
	return Password(password), nil
}
func parseEmail(email string) (Email, error) {
	if email == "" {
		return "", errors.New(shared.RequiredEmail)
	}
	return Email(email), nil
}
func parseTelephoneNumber(telephoneNumber string) (TelephoneNumber, error) {
	if telephoneNumber == "" {
		return "", errors.New(shared.RequiredTelephoneNumber)
	}
	return TelephoneNumber(telephoneNumber), nil
}
func parseGender(gender int64) (Gender, error) {
	if gender == 0 {
		return 0, errors.New(shared.RequiredGender)
	}
	return Gender(gender), nil
}
