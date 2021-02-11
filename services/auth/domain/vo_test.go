package domain

import (
	"errors"
	"testing"

	"github.com/google/uuid"
	"github.com/takuya911/ddd_gql-grpc/services/auth/shared"
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
