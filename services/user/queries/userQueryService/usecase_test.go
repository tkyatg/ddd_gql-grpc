package userqueryservice

import (
	"errors"
	"testing"
	"time"

	gomock "github.com/golang/mock/gomock"
	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
	"github.com/google/uuid"
	"github.com/tkyatg/ddd_gql-grpc/services/user/adapter/hash"
)

type usecaseTestHelper struct {
	ctrl *gomock.Controller
	da   *MockDataAccessor
	uc   Usecase
}

func newUsecaseTestHelper(t *testing.T) *usecaseTestHelper {
	ctrl := gomock.NewController(t)
	da := NewMockDataAccessor(ctrl)
	hash := hash.NewHash()
	uc := NewUsecase(da, hash)

	return &usecaseTestHelper{
		ctrl: ctrl,
		da:   da,
		uc:   uc,
	}
}

func TestUsecaseGetUserByID(t *testing.T) {
	h := newUsecaseTestHelper(t)
	defer h.ctrl.Finish()
	req := getByIDRequest{
		userUUID: uuid.New().String(),
	}
	createdAt := time.Now()
	updatedAt := time.Now()

	h.da.EXPECT().getByID(req).Return(getByIDResponse{
		userUUID:        req.userUUID,
		name:            "name",
		email:           "test@gmail.com",
		password:        "password",
		telephoneNumber: "0909090909090",
		gender:          1,
		createdAt:       createdAt,
		updatedAt:       updatedAt,
	}, nil)

	res, err := h.uc.getByID(req)
	if err != nil {
		t.Fatal(err)
	}

	opts := cmp.Options{
		cmpopts.IgnoreUnexported(getByIDResponse{}),
	}
	if diff := cmp.Diff(getByIDResponse{
		userUUID:        req.userUUID,
		name:            "name",
		email:           "test@gmail.com",
		password:        "password",
		telephoneNumber: "0909090909090",
		gender:          1,
		createdAt:       createdAt,
		updatedAt:       updatedAt,
	}, res, opts); diff != "" {
		t.Fatal(diff)
	}
}

func TestUsecaseGetUserByIDERROR01(t *testing.T) {
	h := newUsecaseTestHelper(t)
	defer h.ctrl.Finish()
	req := getByIDRequest{
		userUUID: uuid.New().String(),
	}
	createdAt := time.Now()
	updatedAt := time.Now()

	err := errors.New("error")
	h.da.EXPECT().getByID(req).Return(getByIDResponse{
		userUUID:        req.userUUID,
		name:            "name",
		email:           "test@gmail.com",
		password:        "password",
		telephoneNumber: "0909090909090",
		gender:          1,
		createdAt:       createdAt,
		updatedAt:       updatedAt,
	}, err)

	_, getByIDErr := h.uc.getByID(req)
	if err != getByIDErr {
		t.Fatal(getByIDErr)
	}

}
