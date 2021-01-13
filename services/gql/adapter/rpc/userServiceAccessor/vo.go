package userserviceaccessor

// CreateUserRequest struct
type CreateUserRequest struct {
	Name            string
	Email           string
	Password        string
	TelephoneNumber string
	Gender          int64
}

// UpdateUserRequest struct
type UpdateUserRequest struct {
	UUID            string
	Name            string
	Email           string
	Password        string
	TelephoneNumber string
	Gender          int64
}

// DeleteUserRequest struct
type DeleteUserRequest struct {
	UUID string
}
