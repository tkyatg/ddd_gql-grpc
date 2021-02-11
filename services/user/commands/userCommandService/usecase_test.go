package usercommandservice

import (
	"errors"
	"testing"

	gomock "github.com/golang/mock/gomock"
	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
	"github.com/google/uuid"
	"github.com/takuya911/ddd_gql-grpc/services/user/adapter/hash"
	"github.com/takuya911/ddd_gql-grpc/services/user/domain"
	"github.com/takuya911/ddd_gql-grpc/services/user/shared"
)

type usecaseTestHelper struct {
	ctrl *gomock.Controller
	repo *domain.MockUserRepository
	uc   Usecase
	hash shared.Hash
}

func newUsecaseTestHelper(t *testing.T) *usecaseTestHelper {
	ctrl := gomock.NewController(t)
	repo := domain.NewMockUserRepository(ctrl)
	hash := hash.NewHash()
	uc := NewUsecase(repo, hash)

	return &usecaseTestHelper{
		ctrl: ctrl,
		repo: repo,
		uc:   uc,
		hash: hash,
	}
}

func TestUsecaseCreate(t *testing.T) {
	h := newUsecaseTestHelper(t)
	defer h.ctrl.Finish()

	uUUID := uuid.New()

	req := createRequest{
		name:            "name",
		email:           "test@gmail.gom",
		password:        "password",
		telephoneNumber: "09084363174",
		gender:          1,
	}

	h.repo.EXPECT().Create(gomock.Any()).Return(domain.UserUUID(uUUID.String()), nil)

	res, err := h.uc.create(req)
	if err != nil {
		t.Fatal(err)
	}

	opts := cmp.Options{
		cmpopts.IgnoreUnexported(createResponse{}),
	}
	if diff := cmp.Diff(createResponse{
		userUUID: uUUID.String(),
	}, res, opts); diff != "" {
		t.Fatal(diff)
	}
}

func TestUsecaseCreateERROR01(t *testing.T) {
	h := newUsecaseTestHelper(t)
	defer h.ctrl.Finish()

	uUUID := uuid.New()

	req := createRequest{
		name:            "name",
		email:           "test@gmail.gom",
		password:        "password",
		telephoneNumber: "09084363174",
		gender:          1,
	}
	err := errors.New("error")
	h.repo.EXPECT().Create(gomock.Any()).Return(domain.UserUUID(uUUID.String()), err)

	_, createErr := h.uc.create(req)
	if err != createErr {
		t.Fatal(createErr)
	}

}

func TestUsecaseUpdate(t *testing.T) {
	h := newUsecaseTestHelper(t)
	defer h.ctrl.Finish()

	req := updateRequest{
		userUUID:        uuid.New().String(),
		name:            "name",
		email:           "test@gmail.com",
		password:        "password",
		telephoneNumber: "09084363176",
		gender:          1,
	}
	id, err := domain.ParseUserUUID(req.userUUID)
	if err != nil {
		t.Fatal(err)
	}
	h.repo.EXPECT().Update(id, gomock.Any()).Return(nil)
	err = h.uc.update(req)
	if err != nil {
		t.Fatal(err)
	}

}

func TestUsecaseUpdateERROR01(t *testing.T) {
	h := newUsecaseTestHelper(t)
	defer h.ctrl.Finish()

	req := updateRequest{
		userUUID:        uuid.New().String(),
		name:            "name",
		email:           "test@gmail.com",
		password:        "password",
		telephoneNumber: "09084363176",
		gender:          1,
	}
	id, parseErr := domain.ParseUserUUID(req.userUUID)
	if parseErr != nil {
		t.Fatal(parseErr)
	}

	err := errors.New("error")
	h.repo.EXPECT().Update(id, gomock.Any()).Return(err)
	updateErr := h.uc.update(req)
	if err != updateErr {
		t.Fatal(updateErr)
	}
}

func TestUsecaseDelete(t *testing.T) {
	h := newUsecaseTestHelper(t)
	defer h.ctrl.Finish()

	req := deleteRequest{
		userUUID: uuid.New().String(),
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

	req := deleteRequest{
		userUUID: uuid.New().String(),
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
