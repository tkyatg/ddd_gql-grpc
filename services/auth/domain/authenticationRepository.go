package domain

type (
	authenticationRepository struct {
		da AuthenticationDataAccessor
	}
	// AuthenticationRepository interface
	AuthenticationRepository interface {
		Login(email Email, pasword Password) (UserUUID, error)
	}
	// AuthenticationDataAccessor interface
	AuthenticationDataAccessor interface {
		login(email Email, pasword Password) (UserUUID, error)
	}
)

// NewAuthenticationRepository func
func NewAuthenticationRepository(
	da AuthenticationDataAccessor,
) AuthenticationRepository {
	return &authenticationRepository{da}
}

func (r *authenticationRepository) Login(email Email, password Password) (UserUUID, error) {
	uuid, err := r.da.login(email, password)
	if err != nil {
		return "", err
	}
	return uuid, nil
}
