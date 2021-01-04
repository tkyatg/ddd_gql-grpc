package usercommandservice

import (
	"errors"
	"testing"

	gomock "github.com/golang/mock/gomock"
	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
	"github.com/google/uuid"
	"github.com/takuya911/project-services/services/user/domain"
)

type usecaseTestHelper struct {
	ctrl *gomock.Controller
	repo *domain.MockUserRepository
	uc   Usecase
}

func newUsecaseTestHelper(t *testing.T) *usecaseTestHelper {
	ctrl := gomock.NewController(t)
	repo := domain.NewMockUserRepository(ctrl)
	uc := NewUsecase(repo)

	return &usecaseTestHelper{
		ctrl: ctrl,
		repo: repo,
		uc:   uc,
	}
}

func TestUsecaseCreate(t *testing.T) {
	h := newUsecaseTestHelper(t)
	defer h.ctrl.Finish()

	userUUID := uuid.New()

	req := createRequest{
		name:            "name",
		email:           "test@gmail.gom",
		password:        "password",
		telephoneNumber: "09084363174",
		gender:          1,
	}
	attr, err := domain.NewUserAttributes(req.name, req.password, req.email, req.telephoneNumber, req.gender)
	if err != nil {
		t.Fatal(err)
	}

	h.repo.EXPECT().Create(attr).Return(domain.UserUUID(userUUID.String()), nil)

	res, err := h.uc.create(req)
	if err != nil {
		t.Fatal(err)
	}

	opts := cmp.Options{
		cmpopts.IgnoreUnexported(createResponse{}),
	}
	if diff := cmp.Diff(createResponse{
		userUUID: userUUID.String(),
	}, res, opts); diff != "" {
		t.Fatal(diff)
	}
}

func TestUsecaseCreateERROR01(t *testing.T) {
	h := newUsecaseTestHelper(t)
	defer h.ctrl.Finish()

	userUUID := uuid.New()

	req := createRequest{
		name:            "name",
		email:           "test@gmail.gom",
		password:        "password",
		telephoneNumber: "09084363174",
		gender:          1,
	}
	attr, newAttrErr := domain.NewUserAttributes(req.name, req.password, req.email, req.telephoneNumber, req.gender)
	if newAttrErr != nil {
		t.Fatal(newAttrErr)
	}
	err := errors.New("error")
	h.repo.EXPECT().Create(attr).Return(domain.UserUUID(userUUID.String()), err)

	_, createErr := h.uc.create(req)
	if err != createErr {
		t.Fatal(createErr)
	}

}

func TestUsecaseUpdate(t *testing.T) {
	h := newUsecaseTestHelper(t)
	defer h.ctrl.Finish()

	userUUID := uuid.New()
	req := updateRequest{
		userUUID:        userUUID.String(),
		name:            "name",
		email:           "test@gmail.com",
		password:        "password",
		telephoneNumber: "09084363176",
		gender:          1,
	}
	attr, err := domain.NewUserAttributes(req.name, req.password, req.email, req.telephoneNumber, req.gender)
	if err != nil {
		t.Fatal(err)
	}
	id, err := domain.ParseUserUUID(req.userUUID)
	if err != nil {
		t.Fatal(err)
	}
	h.repo.EXPECT().Update(id, attr).Return(nil)
	err = h.uc.update(req)
	if err != nil {
		t.Fatal(err)
	}

}

func TestUsecaseUpdateERROR01(t *testing.T) {
	h := newUsecaseTestHelper(t)
	defer h.ctrl.Finish()

	userUUID := uuid.New()
	req := updateRequest{
		userUUID:        userUUID.String(),
		name:            "name",
		email:           "test@gmail.com",
		password:        "password",
		telephoneNumber: "09084363176",
		gender:          1,
	}
	attr, newAttrErr := domain.NewUserAttributes(req.name, req.password, req.email, req.telephoneNumber, req.gender)
	if newAttrErr != nil {
		t.Fatal(newAttrErr)
	}
	id, parseErr := domain.ParseUserUUID(req.userUUID)
	if parseErr != nil {
		t.Fatal(parseErr)
	}

	err := errors.New("error")
	h.repo.EXPECT().Update(id, attr).Return(err)
	updateErr := h.uc.update(req)
	if err != updateErr {
		t.Fatal(updateErr)
	}
}

func TestUsecaseDelete(t *testing.T) {
	h := newUsecaseTestHelper(t)
	defer h.ctrl.Finish()

	userUUID := uuid.New()

	req := deleteRequest{
		userUUID: userUUID.String(),
	}
	id, err := domain.ParseUserUUID(req.userUUID)
	if err != nil {
		t.Fatal(err)
	}
	h.repo.EXPECT().Delete(id).Return(nil)

	err = h.uc.delete(req)
	if err != nil {
		t.Fatal(err)
	}

}

func TestUsecaseDeleteERROR01(t *testing.T) {
	h := newUsecaseTestHelper(t)
	defer h.ctrl.Finish()

	userUUID := uuid.New()

	req := deleteRequest{
		userUUID: userUUID.String(),
	}
	id, parseErr := domain.ParseUserUUID(req.userUUID)
	if parseErr != nil {
		t.Fatal(parseErr)
	}
	err := errors.New("error")
	h.repo.EXPECT().Delete(id).Return(err)

	deleteErr := h.uc.delete(req)
	if err != deleteErr {
		t.Fatal(deleteErr)
	}

}
