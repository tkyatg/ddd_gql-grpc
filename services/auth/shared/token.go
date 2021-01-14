package shared

type (
	// Token interface
	Token interface {
		GenTokenPair(userUUID string) (string, string, error)
	}
)
