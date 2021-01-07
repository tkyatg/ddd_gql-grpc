package shared

type (
	// Env interface
	Env interface {
		GetGraphqlServerPort() string
		GetUserServerName() string
		GetUserServerPort() string
	}
)
