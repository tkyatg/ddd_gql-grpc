package userqueryservice

import (
	"testing"

	gomock "github.com/golang/mock/gomock"
	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
)

func TestUsecaseGetUserByID(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	usecase := NewMockUsecase(ctrl)

	usecase.EXPECT().getUserByID(getUserByIDRequest{
		userUUID: "id",
	}).Return(getUserByIDResponse{
		userUUID:        "id",
		name:            "name",
		email:           "email",
		password:        "password",
		telephoneNumber: "telephoneNumber",
		gender:          1,
	}, nil)

	res, err := usecase.getUserByID(getUserByIDRequest{
		userUUID: "id",
	})
	if err != nil {
		t.Fatal(err)
	}

	opts := cmp.Options{
		cmpopts.IgnoreUnexported(getUserByIDResponse{}),
	}
	if diff := cmp.Diff(getUserByIDResponse{
		userUUID:        "id",
		name:            "name",
		email:           "email",
		password:        "password",
		telephoneNumber: "telephoneNumber",
		gender:          1,
	}, res, opts); diff != "" {
		t.Fatal(diff)
	}
}