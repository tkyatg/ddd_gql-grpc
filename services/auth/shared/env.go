package shared

type (
	// Env interface
	Env interface {
		GetDBHost() string
		GetDBPort() string
		GetDBUser() string
		GetDBPassword() string
		GetDBName() string
		GetAuthServicePort() string
		GetTokenSubject() string
		GetRefreshTokenSubject() string
		GetJwtSignKey() string
	}
)
