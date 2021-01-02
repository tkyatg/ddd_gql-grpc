package domain

type (
	userRepository struct {
		da UserDataAccessor
	}
	// UserRepository interface
	UserRepository interface {
		Create(attr *UserAttributes) error
		Update(id UserUUID, attr *UserAttributes) error
		Delete(id UserUUID) error
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

func (r *userRepository) Create(attr *UserAttributes) error {
	return r.da.create(attr)
}

func (r *userRepository) Update(id UserUUID, attr *UserAttributes) error {
	return r.da.update(id, attr)
}

func (r *userRepository) Delete(id UserUUID) error {
	return r.da.delete(id)
}
