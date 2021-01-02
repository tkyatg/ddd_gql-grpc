package domain

type (
	userRepository struct {
		da UserDataAccessor
	}
	// UserRepository interface
	UserRepository interface {
		create(attr *UserAttributes) error
		update(id UserUUID, attr *UserAttributes) error
		delete(id UserUUID) error
	}
	// UserDataAccessor interface
	UserDataAccessor interface {
		create(attr *UserAttributes) error
		update(id UserUUID, attr *UserAttributes) error
		delete(id UserUUID) error
	}
)

// NewUserRepository func
func NewUserRepository(
	da UserDataAccessor,
) UserRepository {
	return &userRepository{da}
}

func (r *userRepository) create(attr *UserAttributes) error {
	return nil
}

func (r *userRepository) update(id UserUUID, attr *UserAttributes) error {
	return nil
}

func (r *userRepository) delete(id UserUUID) error {
	return nil
}
