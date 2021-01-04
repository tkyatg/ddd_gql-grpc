package usercommandservice

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
	sv   definition.UserCommandServiceServer
	ctx  context.Context
}

func newServerTestHelper(t *testing.T) *serverTestHelper {
	ctrl := gomock.NewController(t)
	usecase := NewMockUsecase(ctrl)
	server := NewServer(usecase)
	ctx := context.Background()

	return &serverTestHelper{
		ctrl: ctrl,
		uc:   usecase,
		sv:   server,
		ctx:  ctx,
	}
}

func TestServerCreate(t *testing.T) {
	h := newServerTestHelper(t)
	defer h.ctrl.Finish()

	uUUID := uuid.New()

	h.uc.EXPECT().create(createRequest{
		name:            "name",
		email:           "test@gmail.com",
		password:        "password",
		telephoneNumber: "09084363175",
		gender:          1,
	}).Return(createResponse{
		userUUID: uUUID.String(),
	}, nil)

	res, err := h.sv.Create(h.ctx, &definition.CreateRequest{
		Name:            "name",
		Email:           "test@gmail.com",
		Password:        "password",
		TelephoneNumber: "09084363175",
		Gender:          1,
	})
	if err != nil {
		t.Fatal(err)
	}

	opts := cmp.Options{}
	if diff := cmp.Diff(&definition.CreateResponse{
		Uuid: uUUID.String(),
	}, res, opts); diff != "" {
		t.Fatal(diff)
	}

}

func TestServerCreateERROR01(t *testing.T) {
	h := newServerTestHelper(t)
	defer h.ctrl.Finish()

	err := errors.New("err")
	h.uc.EXPECT().create(createRequest{
		name:            "name",
		email:           "test@gmail.com",
		password:        "password",
		telephoneNumber: "09084363172",
		gender:          1,
	}).Return(createResponse{}, err)

	_, createErr := h.sv.Create(h.ctx, &definition.CreateRequest{
		Name:            "name",
		Email:           "test@gmail.com",
		Password:        "password",
		TelephoneNumber: "09084363172",
		Gender:          1,
	})
	if err != createErr {
		t.Fatal(createErr)
	}

}

func TestServerUpdate(t *testing.T) {
	h := newServerTestHelper(t)
	defer h.ctrl.Finish()

	req := updateRequest{
		userUUID:        uuid.New().String(),
		name:            "name",
		email:           "test@gmail.com",
		password:        "password",
		telephoneNumber: "09084363172",
		gender:          1,
	}
	h.uc.EXPECT().update(req).Return(nil)

	res, err := h.sv.Update(h.ctx, &definition.UpdateRequest{
		Uuid:            req.userUUID,
		Name:            "name",
		Email:           "test@gmail.com",
		Password:        "password",
		TelephoneNumber: "09084363172",
		Gender:          1,
	})
	if err != nil {
		t.Fatal(err)
	}

	opts := cmp.Options{}
	if diff := cmp.Diff(&definition.UpdateResponse{
		Uuid: req.userUUID,
	}, res, opts); diff != "" {
		t.Fatal(diff)
	}

}

func TestServerUpdateERROR01(t *testing.T) {
	h := newServerTestHelper(t)
	defer h.ctrl.Finish()

	err := errors.New("error")
	req := updateRequest{
		userUUID:        uuid.New().String(),
		name:            "name",
		email:           "test@gmail.com",
		password:        "password",
		telephoneNumber: "09084363172",
		gender:          1,
	}

	h.uc.EXPECT().update(req).Return(err)

	_, updateErr := h.sv.Update(h.ctx, &definition.UpdateRequest{
		Uuid:            req.userUUID,
		Name:            "name",
		Email:           "test@gmail.com",
		Password:        "password",
		TelephoneNumber: "09084363172",
		Gender:          1,
	})
	if err != updateErr {
		t.Fatal(updateErr)
	}

}

func TestServerDelete(t *testing.T) {
	h := newServerTestHelper(t)
	defer h.ctrl.Finish()

	req := deleteRequest{
		userUUID: uuid.New().String(),
	}
	h.uc.EXPECT().delete(req).Return(nil)

	res, err := h.sv.Delete(h.ctx, &definition.DeleteRequest{
		Uuid: req.userUUID,
	})
	if err != nil {
		t.Fatal(err)
	}

	opts := cmp.Options{}
	if diff := cmp.Diff(&definition.DeleteResponse{
		Uuid: req.userUUID,
	}, res, opts); diff != "" {
		t.Fatal(diff)
	}

}

func TestServerDeleteERROR01(t *testing.T) {
	h := newServerTestHelper(t)
	defer h.ctrl.Finish()

	err := errors.New("error")
	req := deleteRequest{
		userUUID: uuid.New().String(),
	}

	h.uc.EXPECT().delete(req).Return(err)

	_, deleteErr := h.sv.Delete(h.ctx, &definition.DeleteRequest{
		Uuid: req.userUUID,
	})
	if err != deleteErr {
		t.Fatal(deleteErr)
	}

}
