package authqueryservice

import (
	"context"
	"errors"
	"testing"

	gomock "github.com/golang/mock/gomock"
	"github.com/google/go-cmp/cmp"
	"github.com/google/uuid"
	definition "github.com/takuya911/project-auth-definition"
)

type serverTestHelper struct {
	ctrl *gomock.Controller
	uc   *MockUsecase
	sv   definition.AuthQueryServiceServer
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

func TestServerGenToken(t *testing.T) {
	h := newServerTestHelper(t)
	defer h.ctrl.Finish()

	uUUID := uuid.New()

	h.uc.EXPECT().genToken(genTokenRequest{
		userUUID: uUUID.String(),
	}).Return(genTokenResponse{
		accessToken:  "accessToken",
		refreshToken: "refreshToken",
	}, nil)

	res, err := h.sv.GenToken(h.ctx, &definition.GenTokenRequest{
		Uuid: uUUID.String(),
	})
	if err != nil {
		t.Fatal(err)
	}

	opts := cmp.Options{}
	if diff := cmp.Diff(&definition.GenTokenResponse{
		AccessToken:  "accessToken",
		RefreshToken: "refreshToken",
	}, res, opts); diff != "" {
		t.Fatal(diff)
	}
}

func TestServerGenTokenError01(t *testing.T) {
	h := newServerTestHelper(t)
	defer h.ctrl.Finish()

	uUUID := uuid.New()

	err := errors.New("error")
	h.uc.EXPECT().genToken(genTokenRequest{
		userUUID: uUUID.String(),
	}).Return(genTokenResponse{}, err)

	_, actualErr := h.sv.GenToken(h.ctx, &definition.GenTokenRequest{
		Uuid: uUUID.String(),
	})
	if err != actualErr {
		t.Fatal(actualErr)
	}

}
