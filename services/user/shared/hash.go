package shared

type (
	// Hash interface
	Hash interface {
		GenEncryptedPassword(password string) (string, error)
		CompareHashAndPass(dbPassword string, formPassword string) error
	}
)
