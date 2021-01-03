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

func TestServerCreate(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	usecase := NewMockUsecase(ctrl)
	server := NewServer(usecase)
	ctx := context.Background()

	userUUID := uuid.New()

	usecase.EXPECT().create(createRequest{
		name:            "name",
		email:           "email",
		password:        "password",
		telephoneNumber: "090-8436-3176",
		gender:          1,
	}).Return(createResponse{
		userUUID: userUUID.String(),
	}, nil)

	res, err := server.Create(ctx, &definition.CreateRequest{
		Name:            "name",
		Email:           "email",
		Password:        "password",
		TelephoneNumber: "090-8436-3176",
		Gender:          1,
	})
	if err != nil {
		t.Fatal(err)
	}

	opts := cmp.Options{}
	if diff := cmp.Diff(&definition.CreateResponse{
		Uuid: userUUID.String(),
	}, res, opts); diff != "" {
		t.Fatal(diff)
	}

}

func TestServerCreateERROR01(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	usecase := NewMockUsecase(ctrl)
	server := NewServer(usecase)
	ctx := context.Background()

	err := errors.New("err")
	usecase.EXPECT().create(createRequest{
		name:            "name",
		email:           "email",
		password:        "password",
		telephoneNumber: "090-8436-3176",
		gender:          1,
	}).Return(createResponse{}, err)

	_, createErr := server.Create(ctx, &definition.CreateRequest{
		Name:            "name",
		Email:           "email",
		Password:        "password",
		TelephoneNumber: "090-8436-3176",
		Gender:          1,
	})
	if err != createErr {
		t.Fatal(createErr)
	}

}

func TestServerUpdate(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	usecase := NewMockUsecase(ctrl)
	server := NewServer(usecase)
	ctx := context.Background()

	userUUID := uuid.New()

	usecase.EXPECT().update(updateRequest{
		userUUID:        userUUID.String(),
		name:            "name",
		email:           "email",
		password:        "password",
		telephoneNumber: "090-8436-3176",
		gender:          1,
	}).Return(nil)

	res, err := server.Update(ctx, &definition.UpdateRequest{
		Uuid:            userUUID.String(),
		Name:            "name",
		Email:           "email",
		Password:        "password",
		TelephoneNumber: "090-8436-3176",
		Gender:          1,
	})
	if err != nil {
		t.Fatal(err)
	}

	opts := cmp.Options{}
	if diff := cmp.Diff(&definition.UpdateResponse{
		Uuid: userUUID.String(),
	}, res, opts); diff != "" {
		t.Fatal(diff)
	}

}

func TestServerUpdateERROR01(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	usecase := NewMockUsecase(ctrl)
	server := NewServer(usecase)
	ctx := context.Background()

	userUUID := uuid.New()
	err := errors.New("error")

	usecase.EXPECT().update(updateRequest{
		userUUID:        userUUID.String(),
		name:            "name",
		email:           "email",
		password:        "password",
		telephoneNumber: "090-8436-3176",
		gender:          1,
	}).Return(err)

	_, updateErr := server.Update(ctx, &definition.UpdateRequest{
		Uuid:            userUUID.String(),
		Name:            "name",
		Email:           "email",
		Password:        "password",
		TelephoneNumber: "090-8436-3176",
		Gender:          1,
	})
	if err != updateErr {
		t.Fatal(updateErr)
	}

}

func TestServerDelete(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	usecase := NewMockUsecase(ctrl)
	server := NewServer(usecase)
	ctx := context.Background()

	userUUID := uuid.New()

	usecase.EXPECT().delete(deleteRequest{
		userUUID: userUUID.String(),
	}).Return(nil)

	res, err := server.Delete(ctx, &definition.DeleteRequest{
		Uuid: userUUID.String(),
	})
	if err != nil {
		t.Fatal(err)
	}

	opts := cmp.Options{}
	if diff := cmp.Diff(&definition.DeleteResponse{
		Uuid: userUUID.String(),
	}, res, opts); diff != "" {
		t.Fatal(diff)
	}

}

func TestServerDeleteERROR01(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	usecase := NewMockUsecase(ctrl)
	server := NewServer(usecase)
	ctx := context.Background()

	userUUID := uuid.New()
	err := errors.New("error")

	usecase.EXPECT().delete(deleteRequest{
		userUUID: userUUID.String(),
	}).Return(err)

	_, deleteErr := server.Delete(ctx, &definition.DeleteRequest{
		Uuid: userUUID.String(),
	})
	if err != deleteErr {
		t.Fatal(deleteErr)
	}

}
