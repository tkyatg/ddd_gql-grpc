package userqueryservice

import (
	"errors"
	"testing"

	gomock "github.com/golang/mock/gomock"
	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
	"github.com/google/uuid"
)

type usecaseTestHelper struct {
	ctrl *gomock.Controller
	da   *MockDataAccessor
	uc   Usecase
}

func newUsecaseTestHelper(t *testing.T) *usecaseTestHelper {
	ctrl := gomock.NewController(t)
	da := NewMockDataAccessor(ctrl)
	uc := NewUsecase(da)

	return &usecaseTestHelper{
		ctrl: ctrl,
		da:   da,
		uc:   uc,
	}
}

func TestUsecaseGetUserByID(t *testing.T) {
	h := newUsecaseTestHelper(t)
	defer h.ctrl.Finish()
	req := getUserByIDRequest{
		userUUID: uuid.New().String(),
	}

	h.da.EXPECT().getByID(req).Return(getUserByIDResponse{
		userUUID:        req.userUUID,
		name:            "name",
		email:           "test@gmail.com",
		password:        "password",
		telephoneNumber: "0909090909090",
		gender:          1,
	}, nil)

	res, err := h.uc.getByID(req)
	if err != nil {
		t.Fatal(err)
	}

	opts := cmp.Options{
		cmpopts.IgnoreUnexported(getUserByIDResponse{}),
	}
	if diff := cmp.Diff(getUserByIDResponse{
		userUUID:        req.userUUID,
		name:            "name",
		email:           "test@gmail.com",
		password:        "password",
		telephoneNumber: "0909090909090",
		gender:          1,
	}, res, opts); diff != "" {
		t.Fatal(diff)
	}
}

func TestUsecaseGetUserByIDERROR01(t *testing.T) {
	h := newUsecaseTestHelper(t)
	defer h.ctrl.Finish()
	req := getUserByIDRequest{
		userUUID: uuid.New().String(),
	}
	err := errors.New("error")
	h.da.EXPECT().getByID(req).Return(getUserByIDResponse{
		userUUID:        req.userUUID,
		name:            "name",
		email:           "test@gmail.com",
		password:        "password",
		telephoneNumber: "0909090909090",
		gender:          1,
	}, err)

	_, getByIDErr := h.uc.getByID(req)
	if err != getByIDErr {
		t.Fatal(getByIDErr)
	}

}
