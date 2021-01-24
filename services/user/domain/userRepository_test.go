package domain

import (
	"context"
	"errors"
	"testing"

	"github.com/takuya911/project-services/services/user/shared"

	gomock "github.com/golang/mock/gomock"
	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
	"github.com/google/uuid"
)

type userRepositoryTestHelper struct {
	ctrl *gomock.Controller
	da   *MockUserDataAccessor
	repo UserRepository
	ctx  context.Context
}

func newUserRepositoryTestHelper(t *testing.T) *userRepositoryTestHelper {
	ctrl := gomock.NewController(t)
	da := NewMockUserDataAccessor(ctrl)
	repo := NewUserRepository(da)
	ctx := context.Background()

	return &userRepositoryTestHelper{
		ctrl: ctrl,
		da:   da,
		repo: repo,
		ctx:  ctx,
	}
}

func TestRepositoryCreate(t *testing.T) {
	h := newUserRepositoryTestHelper(t)
	defer h.ctrl.Finish()

	uUUID := uuid.New()
	h.da.EXPECT().emailAlreadyUsedCreate(Email("test@gmail.com")).Return(false, nil)

	h.da.EXPECT().create(UserAttributes{
		name:            UserName("name"),
		email:           Email("test@gmail.com"),
		password:        Password("password"),
		telephoneNumber: TelephoneNumber("09084363172"),
		gender:          Gender(1),
	}).Return(UserUUID(uUUID.String()), nil)

	res, err := h.repo.Create(UserAttributes{
		name:            UserName("name"),
		email:           Email("test@gmail.com"),
		password:        Password("password"),
		telephoneNumber: TelephoneNumber("09084363172"),
		gender:          Gender(1),
	})
	if err != nil {
		t.Fatal(err)
	}

	opts := cmp.Options{
		cmpopts.IgnoreUnexported(),
	}
	if diff := cmp.Diff(
		UserUUID(uUUID.String()), res, opts); diff != "" {
		t.Fatal(diff)
	}
}

func TestRepositoryCreateEROOR01(t *testing.T) {
	h := newUserRepositoryTestHelper(t)
	defer h.ctrl.Finish()

	uUUID := uuid.New()

	err := errors.New("error")

	h.da.EXPECT().emailAlreadyUsedCreate(Email("test@gmail.com")).Return(false, nil)

	h.da.EXPECT().create(UserAttributes{
		name:            UserName("name"),
		email:           Email("test@gmail.com"),
		password:        Password("password"),
		telephoneNumber: TelephoneNumber("09084363172"),
		gender:          Gender(1),
	}).Return(UserUUID(uUUID.String()), err)

	_, createErr := h.repo.Create(UserAttributes{
		name:            UserName("name"),
		email:           Email("test@gmail.com"),
		password:        Password("password"),
		telephoneNumber: TelephoneNumber("09084363172"),
		gender:          Gender(1),
	})
	if err != createErr {
		t.Fatal(createErr)
	}
}

func TestRepositoryCreateEROOR02(t *testing.T) {
	h := newUserRepositoryTestHelper(t)
	defer h.ctrl.Finish()

	err := errors.New("error")

	h.da.EXPECT().emailAlreadyUsedCreate(Email("test@gmail.com")).Return(false, err)

	_, createErr := h.repo.Create(UserAttributes{
		name:            UserName("name"),
		email:           Email("test@gmail.com"),
		password:        Password("password"),
		telephoneNumber: TelephoneNumber("09084363172"),
		gender:          Gender(1),
	})
	if err != createErr {
		t.Fatal(createErr)
	}
}

func TestRepositoryCreateEROOR03(t *testing.T) {
	h := newUserRepositoryTestHelper(t)
	defer h.ctrl.Finish()

	h.da.EXPECT().emailAlreadyUsedCreate(Email("test@gmail.com")).Return(true, nil)
	err := errors.New(shared.EmailAlreadyUsed)

	_, createErr := h.repo.Create(UserAttributes{
		name:            UserName("name"),
		email:           Email("test@gmail.com"),
		password:        Password("password"),
		telephoneNumber: TelephoneNumber("09084363172"),
		gender:          Gender(1),
	})

	if err.Error() != createErr.Error() {
		t.Fatal(createErr)
	}
}

func TestRepositoryUpdate(t *testing.T) {
	h := newUserRepositoryTestHelper(t)
	defer h.ctrl.Finish()

	uUUID := uuid.New()

	h.da.EXPECT().emailAlreadyUsedUpdate(UserUUID(uUUID.String()), Email("test@gmail.com")).Return(false, nil)

	h.da.EXPECT().update(UserUUID(uUUID.String()), UserAttributes{
		name:            UserName("name"),
		email:           Email("test@gmail.com"),
		password:        Password("password"),
		telephoneNumber: TelephoneNumber("09084363172"),
		gender:          Gender(1),
	}).Return(nil)

	err := h.repo.Update(UserUUID(uUUID.String()), UserAttributes{
		name:            UserName("name"),
		email:           Email("test@gmail.com"),
		password:        Password("password"),
		telephoneNumber: TelephoneNumber("09084363172"),
		gender:          Gender(1),
	})
	if err != nil {
		t.Fatal(err)
	}
}

func TestRepositoryUpdateEROOR01(t *testing.T) {
	h := newUserRepositoryTestHelper(t)
	defer h.ctrl.Finish()

	uUUID := uuid.New()

	h.da.EXPECT().emailAlreadyUsedUpdate(UserUUID(uUUID.String()), Email("test@gmail.com")).Return(false, nil)
	err := errors.New("error")
	h.da.EXPECT().update(UserUUID(uUUID.String()), UserAttributes{
		name:            UserName("name"),
		email:           Email("test@gmail.com"),
		password:        Password("password"),
		telephoneNumber: TelephoneNumber("09084363172"),
		gender:          Gender(1),
	}).Return(err)

	updateErr := h.repo.Update(UserUUID(uUUID.String()), UserAttributes{
		name:            UserName("name"),
		email:           Email("test@gmail.com"),
		password:        Password("password"),
		telephoneNumber: TelephoneNumber("09084363172"),
		gender:          Gender(1),
	})
	if err != updateErr {
		t.Fatal(updateErr)
	}
}

func TestRepositoryUpdateEROOR02(t *testing.T) {
	h := newUserRepositoryTestHelper(t)
	defer h.ctrl.Finish()

	uUUID := uuid.New()

	err := errors.New("error")
	h.da.EXPECT().emailAlreadyUsedUpdate(UserUUID(uUUID.String()), Email("test@gmail.com")).Return(false, err)

	updateErr := h.repo.Update(UserUUID(uUUID.String()), UserAttributes{
		name:            UserName("name"),
		email:           Email("test@gmail.com"),
		password:        Password("password"),
		telephoneNumber: TelephoneNumber("09084363172"),
		gender:          Gender(1),
	})
	if err != updateErr {
		t.Fatal(updateErr)
	}

}

func TestRepositoryUpdateEROOR03(t *testing.T) {
	h := newUserRepositoryTestHelper(t)
	defer h.ctrl.Finish()

	uUUID := uuid.New()

	err := errors.New(shared.EmailAlreadyUsed)

	h.da.EXPECT().emailAlreadyUsedUpdate(UserUUID(uUUID.String()), Email("test@gmail.com")).Return(true, nil)

	updateErr := h.repo.Update(UserUUID(uUUID.String()), UserAttributes{
		name:            UserName("name"),
		email:           Email("test@gmail.com"),
		password:        Password("password"),
		telephoneNumber: TelephoneNumber("09084363172"),
		gender:          Gender(1),
	})
	if err.Error() != updateErr.Error() {
		t.Fatal(updateErr)
	}
}

func TestRepositoryDelete(t *testing.T) {
	h := newUserRepositoryTestHelper(t)
	defer h.ctrl.Finish()

	uUUID := uuid.New()

	h.da.EXPECT().delete(UserUUID(uUUID.String())).Return(nil)

	err := h.repo.Delete(UserUUID(uUUID.String()))
	if err != nil {
		t.Fatal(err)
	}
}

func TestRepositoryDeleteEROOR01(t *testing.T) {
	h := newUserRepositoryTestHelper(t)
	defer h.ctrl.Finish()

	uUUID := uuid.New()

	err := errors.New("error")
	h.da.EXPECT().delete(UserUUID(uUUID.String())).Return(err)

	deleteErr := h.repo.Delete(UserUUID(uUUID.String()))
	if err != deleteErr {
		t.Fatal(deleteErr)
	}

}
