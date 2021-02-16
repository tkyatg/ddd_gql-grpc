package authqueryservice

import (
	"testing"

	gomock "github.com/golang/mock/gomock"
	"github.com/takuya911/ddd_gql-grpc/services/auth/adapter/env"
	"github.com/takuya911/ddd_gql-grpc/services/auth/adapter/jwt"
	"github.com/takuya911/ddd_gql-grpc/services/auth/shared"
)

type usecaseTestHelper struct {
	ctrl  *gomock.Controller
	uc    Usecase
	token shared.Token
}

func newUsecaseTestHelper(t *testing.T) *usecaseTestHelper {
	ctrl := gomock.NewController(t)
	env := env.NewEnv()
	token := jwt.NewToken(env)
	uc := NewUsecase(token)

	return &usecaseTestHelper{
		ctrl:  ctrl,
		uc:    uc,
		token: token,
	}
}

// func TestUsecaseGenToken(t *testing.T) {
// 	h := newUsecaseTestHelper(t)
// 	defer h.ctrl.Finish()

// 	uUUID := uuid.New()

// 	req := genTokenRequest{
// 		userUUID: uUUID.String(),
// 	}

// 	res, err := h.uc.genToken(req)
// 	if err != nil {
// 		t.Fatal(err)
// 	}

// 	opts := cmp.Options{
// 		cmpopts.IgnoreUnexported(genTokenResponse{}),
// 	}
// 	if diff := cmp.Diff(genTokenResponse{
// 		"accessToken",
// 		"refreshToken",
// 	}, res, opts); diff != "" {
// 		t.Fatal(diff)
// 	}
// }
