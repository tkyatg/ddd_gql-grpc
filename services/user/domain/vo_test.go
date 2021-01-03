package domain

import (
	"errors"
	"testing"

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
}
