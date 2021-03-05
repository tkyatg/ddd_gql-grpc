package userqueryservice

import (
	"context"
	"errors"
	"testing"
	"time"

	gomock "github.com/golang/mock/gomock"
	"github.com/google/go-cmp/cmp"
	"github.com/google/uuid"
	definition "github.com/tkyatg/user-definition"
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

	req := getByIDRequest{
		userUUID: uuid.New().String(),
	}
	createdAt := time.Now()
	updatedAt := time.Now()
	h.uc.EXPECT().getByID(req).Return(getByIDResponse{
		userUUID:        req.userUUID,
		name:            "name",
		email:           "test@gmail.com",
		password:        "password",
		telephoneNumber: "0909090909090",
		gender:          1,
		createdAt:       createdAt,
		updatedAt:       updatedAt,
	}, nil)

	res, err := h.sv.GetByID(h.ctx, &definition.GetByIDRequest{
		Uuid: req.userUUID,
	})
	if err != nil {
		t.Fatal(err)
	}

	opts := cmp.Options{}
	if diff := cmp.Diff(&definition.GetByIDResponse{
		Uuid:            req.userUUID,
		Name:            "name",
		Email:           "test@gmail.com",
		Password:        "password",
		TelephoneNumber: "0909090909090",
		Gender:          1,
	}, &definition.GetByIDResponse{
		Uuid:            res.Uuid,
		Name:            res.Name,
		Email:           res.Email,
		Password:        res.Password,
		TelephoneNumber: res.TelephoneNumber,
		Gender:          res.Gender,
	}, opts); diff != "" {
		t.Fatal(diff)
	}

}

func TestServerGetByIDERROR01(t *testing.T) {
	h := newServerTestHelper(t)
	defer h.ctrl.Finish()

	req := getByIDRequest{
		userUUID: uuid.New().String(),
	}
	err := errors.New("error")
	h.uc.EXPECT().getByID(req).Return(getByIDResponse{}, err)
	_, getByIDErr := h.sv.GetByID(h.ctx, &definition.GetByIDRequest{
		Uuid: req.userUUID,
	})
	if err != getByIDErr {
		t.Fatal(getByIDErr)
	}

}
