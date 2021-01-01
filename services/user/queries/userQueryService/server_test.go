package userqueryservice

import (
	"context"
	"testing"

	gomock "github.com/golang/mock/gomock"
	"github.com/google/go-cmp/cmp"
	definition "github.com/takuya911/project-user-definition"
)

func TestServerGetUserByID(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	usecase := NewMockUsecase(ctrl)
	server := NewServer(usecase)
	ctx := context.Background()

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

	res, err := server.GetUserByID(ctx, &definition.GetUserRequest{
		Uuid: "id",
	})
	if err != nil {
		t.Fatal(err)
	}

	opts := cmp.Options{}
	if diff := cmp.Diff(&definition.GetUserResponse{
		Uuid:            "id",
		Name:            "name",
		Email:           "email",
		Password:        "password",
		TelephoneNumber: "telephoneNumber",
		Gender:          1,
	}, res, opts); diff != "" {
		t.Fatal(diff)
	}

}
