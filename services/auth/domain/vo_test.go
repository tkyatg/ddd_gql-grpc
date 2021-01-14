package domain

import (
	"errors"
	"testing"

	"github.com/google/uuid"
	"github.com/takuya911/project-services/services/user/shared"
)

func TestVoParseUserUUID(t *testing.T) {
	t.Parallel()
	if userUUID, err := ParseUserUUID(""); userUUID != "" && err != errors.New(shared.RequiredUserUUID) {
		t.Fatal(userUUID, err)
	}
	if userUUID, err := ParseUserUUID("1111111"); userUUID != "" && err != errors.New(shared.InvalidUUIDFormat) {
		t.Fatal(userUUID, err)
	}
	if userUUID, err := ParseUserUUID(uuid.New().String()); userUUID == "" || err != nil {
		t.Fatal(userUUID, err)
	}
}

func TestVoParsePassword(t *testing.T) {
	t.Parallel()
	if password, err := ParsePassword(""); password != "" && err != errors.New(shared.RequiredPassword) {
		t.Fatal(password, err)
	}
	if password, err := ParsePassword("111111"); password != "" && err != errors.New(shared.InvalidPasswordLength) {
		t.Fatal(password, err)
	}
	if password, err := ParsePassword("gdpspvls"); password == "" || err != nil {
		t.Fatal(password, err)
	}
}

func TestVoParseEmail(t *testing.T) {
	t.Parallel()
	if email, err := ParseEmail(""); email != "" && err != errors.New(shared.RequiredEmail) {
		t.Fatal(email, err)
	}
	if email, err := ParseEmail("e@"); email != "" && err != errors.New(shared.InvalidEmailLength) {
		t.Fatal(email, err)
	}
	if email, err := ParseEmail("testgmail.com"); email != "" && err != errors.New(shared.InvalidEmailFormat) {
		t.Fatal(email, err)
	}
	if email, err := ParsePassword("test@gmail.com"); email == "" || err != nil {
		t.Fatal(email, err)
	}
}
