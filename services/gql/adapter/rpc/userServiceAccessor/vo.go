package userserviceaccessor

// CreateUserRequest struct
type CreateUserRequest struct {
	Name            string
	Email           string
	Password        string
	TelephoneNumber string
	Gender          int64
}
