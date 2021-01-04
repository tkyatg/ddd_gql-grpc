package userqueryservice

import (
	"context"
	"errors"
	"testing"

	gomock "github.com/golang/mock/gomock"
	"github.com/google/go-cmp/cmp"
	"github.com/google/uuid"
	definition "github.com/takuya911/project-user-definition"
)

type serverTestHelper struct {
	ctrl *gomock.Controller
	uc   *MockUsecase
	sv   definition.UserQueryServiceServer
	ctx  context.Context
}

func newServerTestHelper(t *testing.T) *serverTestHelper {
	ctrl := gomock.NewController(t)
	uc := NewMockUsecase(ctrl)
	sv := NewServer(uc)
	ctx := context.Background()

	return &serverTestHelper{
		ctrl: ctrl,
		uc:   uc,
		sv:   sv,
		ctx:  ctx,
	}
}

func TestServerGetByID(t *testing.T) {
	h := newServerTestHelper(t)
	defer h.ctrl.Finish()

	req := getUserByIDRequest{
		userUUID: uuid.New().String(),
	}
	h.uc.EXPECT().getByID(req).Return(getUserByIDResponse{
		userUUID:        req.userUUID,
		name:            "name",
		email:           "test@gmail.com",
		password:        "password",
		telephoneNumber: "0909090909090",
		gender:          1,
	}, nil)

	res, err := h.sv.GetByID(h.ctx, &definition.GetUserRequest{
		Uuid: req.userUUID,
	})
	if err != nil {
		t.Fatal(err)
	}

	opts := cmp.Options{}
	if diff := cmp.Diff(&definition.GetUserResponse{
		Uuid:            req.userUUID,
		Name:            "name",
		Email:           "test@gmail.com",
		Password:        "password",
		TelephoneNumber: "0909090909090",
		Gender:          1,
	}, res, opts); diff != "" {
		t.Fatal(diff)
	}

}

func TestServerGetByIDERROR01(t *testing.T) {
	h := newServerTestHelper(t)
	defer h.ctrl.Finish()

	req := getUserByIDRequest{
		userUUID: uuid.New().String(),
	}
	err := errors.New("error")
	h.uc.EXPECT().getByID(req).Return(getUserByIDResponse{}, err)
	_, getByIDErr := h.sv.GetByID(h.ctx, &definition.GetUserRequest{
		Uuid: req.userUUID,
	})
	if err != getByIDErr {
		t.Fatal(getByIDErr)
	}

}
