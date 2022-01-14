package schema

type LoginRequest struct {
	Email    string
	Password string
}

type TokenResponse struct {
	AccessToken  string
	TokenType    string
	ExpiresIn    int
	RefreshToken string
}
