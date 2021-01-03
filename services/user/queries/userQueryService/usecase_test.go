package userqueryservice

import (
	"errors"
	"testing"

	gomock "github.com/golang/mock/gomock"
	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
)

func TestUsecaseGetUserByID(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	usecase := NewMockUsecase(ctrl)

	usecase.EXPECT().getByID(getUserByIDRequest{
		userUUID: "id",
	}).Return(getUserByIDResponse{
		userUUID:        "id",
		name:            "name",
		email:           "email",
		password:        "password",
		telephoneNumber: "telephoneNumber",
		gender:          1,
	}, nil)

	res, err := usecase.getByID(getUserByIDRequest{
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

func TestUsecaseGetUserByIDERROR01(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	usecase := NewMockUsecase(ctrl)

	err := errors.New("error")
	usecase.EXPECT().getByID(getUserByIDRequest{
		userUUID: "id",
	}).Return(getUserByIDResponse{}, err)

	_, getByIDErr := usecase.getByID(getUserByIDRequest{
		userUUID: "id",
	})
	if err != getByIDErr {
		t.Fatal(getByIDErr)
	}

}
