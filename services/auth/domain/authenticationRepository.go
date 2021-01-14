package domain

type (
	authenticationRepository struct {
		da AuthenticationDataAccessor
	}
	// AuthenticationRepository interface
	AuthenticationRepository interface {
		Login(email Email) (UserUUID, Password, error)
	}
	// AuthenticationDataAccessor interface
	AuthenticationDataAccessor interface {
		login(email Email) (UserUUID, Password, error)
	}
)

// NewAuthenticationRepository func
func NewAuthenticationRepository(
	da AuthenticationDataAccessor,
) AuthenticationRepository {
	return &authenticationRepository{da}
}

func (r *authenticationRepository) Login(email Email) (UserUUID, Password, error) {
	uuid, password, err := r.da.login(email)
	if err != nil {
		return "", "", err
	}
	return uuid, password, nil
}
