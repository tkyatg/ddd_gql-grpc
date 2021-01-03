package usercommandservice

import (
	"errors"
	"testing"

	gomock "github.com/golang/mock/gomock"
	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
	"github.com/google/uuid"
)

func TestUsecaseCreate(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	usecase := NewMockUsecase(ctrl)

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

	res, err := usecase.create(createRequest{
		name:            "name",
		email:           "email",
		password:        "password",
		telephoneNumber: "090-8436-3176",
		gender:          1,
	})
	if err != nil {
		t.Fatal(err)
	}

	opts := cmp.Options{
		cmpopts.IgnoreUnexported(createResponse{}),
	}
	if diff := cmp.Diff(createResponse{
		userUUID: userUUID.String(),
	}, res, opts); diff != "" {
		t.Fatal(diff)
	}
}

func TestUsecaseCreateERROR01(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	usecase := NewMockUsecase(ctrl)

	err := errors.New("error")

	usecase.EXPECT().create(createRequest{
		name:            "name",
		email:           "email",
		password:        "password",
		telephoneNumber: "090-8436-3176",
		gender:          1,
	}).Return(createResponse{}, err)

	_, createErr := usecase.create(createRequest{
		name:            "name",
		email:           "email",
		password:        "password",
		telephoneNumber: "090-8436-3176",
		gender:          1,
	})
	if err != createErr {
		t.Fatal(createErr)
	}

}

func TestUsecaseUpdate(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	usecase := NewMockUsecase(ctrl)

	userUUID := uuid.New()

	usecase.EXPECT().update(updateRequest{
		userUUID:        userUUID.String(),
		name:            "name",
		email:           "email",
		password:        "password",
		telephoneNumber: "090-8436-3176",
		gender:          1,
	}).Return(nil)

	err := usecase.update(updateRequest{
		userUUID:        userUUID.String(),
		name:            "name",
		email:           "email",
		password:        "password",
		telephoneNumber: "090-8436-3176",
		gender:          1,
	})
	if err != nil {
		t.Fatal(err)
	}

}

func TestUsecaseUpdateERROR01(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	usecase := NewMockUsecase(ctrl)

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

	updateErr := usecase.update(updateRequest{
		userUUID:        userUUID.String(),
		name:            "name",
		email:           "email",
		password:        "password",
		telephoneNumber: "090-8436-3176",
		gender:          1,
	})
	if err != updateErr {
		t.Fatal(updateErr)
	}
}

func TestUsecaseDelete(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	usecase := NewMockUsecase(ctrl)

	userUUID := uuid.New()

	usecase.EXPECT().delete(deleteRequest{
		userUUID: userUUID.String(),
	}).Return(nil)

	err := usecase.delete(deleteRequest{
		userUUID: userUUID.String(),
	})
	if err != nil {
		t.Fatal(err)
	}

}

func TestUsecaseDeleteERROR01(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	usecase := NewMockUsecase(ctrl)

	userUUID := uuid.New()
	err := errors.New("error")
	usecase.EXPECT().delete(deleteRequest{
		userUUID: userUUID.String(),
	}).Return(err)

	deleteErr := usecase.delete(deleteRequest{
		userUUID: userUUID.String(),
	})
	if err != deleteErr {
		t.Fatal(deleteErr)
	}
}
