package domain

type (
	userRepository struct {
		da UserDataAccessor
	}
	// UserRepository interface
	UserRepository interface {
		Login(attr string) error
	}
	// UserDataAccessor interface
	UserDataAccessor interface {
		login(attr string) error
	}
)

// NewAuthenticationRepository func
func NewAuthenticationRepository(
	da UserDataAccessor,
) UserRepository {
	return &userRepository{da}
}

func (r *userRepository) Login(attr string) error {
	return r.da.login(attr)
}
