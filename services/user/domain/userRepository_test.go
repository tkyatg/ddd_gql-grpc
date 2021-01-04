package domain

import (
	"testing"

	gomock "github.com/golang/mock/gomock"
	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
	"github.com/google/uuid"
)

func TestRepositoryCreate(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	da := NewMockUserDataAccessor(ctrl)
	repository := NewUserRepository(da)

	userUUID := uuid.New()

	da.EXPECT().create(UserAttributes{
		name:            UserName("name"),
		email:           Email("email"),
		password:        Password("password"),
		telephoneNumber: TelephoneNumber("09084363176"),
		gender:          Gender(1),
	}).Return(UserUUID(userUUID.String()), nil)

	res, err := repository.Create(UserAttributes{
		name:            UserName("name"),
		email:           Email("email"),
		password:        Password("password"),
		telephoneNumber: TelephoneNumber("09084363176"),
		gender:          Gender(1),
	})
	if err != nil {
		t.Fatal(err)
	}

	opts := cmp.Options{
		cmpopts.IgnoreUnexported(),
	}
	if diff := cmp.Diff(
		UserUUID(userUUID.String()), res, opts); diff != "" {
		t.Fatal(diff)
	}
}
