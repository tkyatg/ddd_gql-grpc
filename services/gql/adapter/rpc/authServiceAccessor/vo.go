package authserviceaccessor

type (
	// GenTokenRequest struct
	GenTokenRequest struct {
		UUID string
	}
	// GenTokenResponse struct
	GenTokenResponse struct {
		TokenPair TokenPair
	}
	// LoginRequest struct
	LoginRequest struct {
		Email    string
		Password string
	}
	// LoginResponse struct
	LoginResponse struct {
		TokenPair TokenPair
	}
	// TokenPair struct
	TokenPair struct {
		AccessToken  string
		RefreshToken string
	}
)
