package shared

type (
	// Env interface
	Env interface {
		GetDBHost() string
		GetDBPort() string
		GetDBUser() string
		GetDBPassword() string
		GetDBName() string
		GetUserServicePort() string
	}
)
