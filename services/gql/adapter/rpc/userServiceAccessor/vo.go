package userserviceaccessor

import "google.golang.org/protobuf/types/known/timestamppb"

type (
	// GetByIDRequest struct
	GetByIDRequest struct {
		UUID string
	}

	// GetByIDResponse struct
	GetByIDResponse struct {
		UUID            string
		Name            string
		Email           string
		Password        string
		TelephoneNumber string
		Gender          int64
		CreatedAt       *timestamppb.Timestamp
		UpdatedAt       *timestamppb.Timestamp
	}

	// GetByEmailAndPasswordRequest struct
	GetByEmailAndPasswordRequest struct {
		Email    string
		Password string
	}

	// GetByEmailAndPasswordResponse struct
	GetByEmailAndPasswordResponse struct {
		UUID            string
		Name            string
		Email           string
		Password        string
		TelephoneNumber string
		Gender          int64
		CreatedAt       *timestamppb.Timestamp
		UpdatedAt       *timestamppb.Timestamp
	}

	// CreateUserRequest struct
	CreateUserRequest struct {
		Name            string
		Email           string
		Password        string
		TelephoneNumber string
		Gender          int64
	}

	// CreateUserResponse struct
	CreateUserResponse struct {
		UUID string
	}

	// UpdateUserRequest struct
	UpdateUserRequest struct {
		UUID            string
		Name            string
		Email           string
		Password        string
		TelephoneNumber string
		Gender          int64
	}

	// UpdateUserResponse struct
	UpdateUserResponse struct {
		UUID string
	}

	// DeleteUserRequest struct
	DeleteUserRequest struct {
		UUID string
	}

	// DeleteUserResponse struct
	DeleteUserResponse struct {
		UUID string
	}
)
