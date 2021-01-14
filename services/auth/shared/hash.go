package shared

type (
	// Hash interface
	Hash interface {
		CompareHashAndPass(dbPassword string, formPassword string) error
	}
)
