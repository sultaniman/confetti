package schema

type LoginRequest struct {
	Email    string
	Password string
}

type RegisterRequest struct {
	Email    string
	Password string
}

type ResetPasswordRequest struct {
	Email string
}

type NewPasswordRequest struct {
	Password string
}

type TokenResponse struct {
	AccessToken  string
	TokenType    string
	ExpiresIn    int
	RefreshToken string
}
