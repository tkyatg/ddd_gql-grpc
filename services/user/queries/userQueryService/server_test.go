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

func TestServerGetByID(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	usecase := NewMockUsecase(ctrl)
	server := NewServer(usecase)
	ctx := context.Background()

	userUUID := uuid.New()

	usecase.EXPECT().getByID(getUserByIDRequest{
		userUUID: userUUID.String(),
	}).Return(getUserByIDResponse{
		userUUID:        userUUID.String(),
		name:            "name",
		email:           "email",
		password:        "password",
		telephoneNumber: "090-8436-3174",
		gender:          1,
	}, nil)

	res, err := server.GetByID(ctx, &definition.GetUserRequest{
		Uuid: userUUID.String(),
	})
	if err != nil {
		t.Fatal(err)
	}

	opts := cmp.Options{}
	if diff := cmp.Diff(&definition.GetUserResponse{
		Uuid:            userUUID.String(),
		Name:            "name",
		Email:           "email",
		Password:        "password",
		TelephoneNumber: "090-8436-3174",
		Gender:          1,
	}, res, opts); diff != "" {
		t.Fatal(diff)
	}

}

func TestServerGetByIDERROR01(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	usecase := NewMockUsecase(ctrl)
	server := NewServer(usecase)
	ctx := context.Background()

	userUUID := uuid.New()
	err := errors.New("error")
	usecase.EXPECT().getByID(getUserByIDRequest{
		userUUID: userUUID.String(),
	}).Return(getUserByIDResponse{}, err)

	_, getByIDErr := server.GetByID(ctx, &definition.GetUserRequest{
		Uuid: userUUID.String(),
	})
	if err != getByIDErr {
		t.Fatal(getByIDErr)
	}

}
