package services

import (
	"crypto/rsa"
	"github.com/gofiber/fiber/v2"
	"github.com/imanhodjaev/getout/platform/http"
	"github.com/imanhodjaev/getout/platform/schema"
	"github.com/imanhodjaev/getout/platform/shared"
	"github.com/lestrrat-go/jwx/jwa"
	"github.com/lestrrat-go/jwx/jwk"
	"github.com/lestrrat-go/jwx/jwt"
	"github.com/spf13/viper"
	"time"
)

const RefreshTokenCookieName = "refresh_token"
const DefaultJWA = jwa.RS256   // alg
const DefaultKeyID = "default" // kid

type JWXService struct {
	privateJWK jwk.Key
	jwks       *jwk.Set
}

func NewJWXService(privateKey *rsa.PrivateKey) (*JWXService, error) {
	privateJWK, err := jwk.New(privateKey)
	if err != nil {
		return nil, err
	}

	if err = privateJWK.Set(jwk.AlgorithmKey, DefaultJWA); err != nil {
		return nil, err
	}

	if err = privateJWK.Set(jwk.KeyIDKey, DefaultKeyID); err != nil {
		return nil, err
	}

	publicJWK, err := jwk.New(privateKey.Public())
	if err != nil {
		return nil, err
	}

	if err = publicJWK.Set(jwk.AlgorithmKey, DefaultJWA); err != nil {
		return nil, err
	}

	if err = publicJWK.Set(jwk.KeyIDKey, DefaultKeyID); err != nil {
		return nil, err
	}

	jwks := jwk.NewSet()
	err = jwks.Set("keys", []jwk.Key{publicJWK})
	if err != nil {
		return nil, err
	}

	return &JWXService{
		privateJWK: privateJWK,
		jwks:       &jwks,
	}, nil
}

func (s *JWXService) JWKS() *jwk.Set {
	return s.jwks
}

func (s *JWXService) GetRefreshTokenCookie(token jwt.Token) (*fiber.Cookie, error) {
	alg := jwa.SignatureAlgorithm(s.privateJWK.Algorithm())
	signed, err := jwt.Sign(token, alg, s.privateJWK)
	if err != nil {
		return nil, jwxError("unable to sign refresh token")
	}

	return &fiber.Cookie{
		Name:     RefreshTokenCookieName,
		Value:    string(signed),
		Expires:  token.Expiration(),
		Secure:   true,
		HTTPOnly: true,
		SameSite: fiber.CookieSameSiteStrictMode,
	}, nil
}

type RefreshTokenFunc func(jwt.Token) (jwt.Token, error)

func (s *JWXService) RefreshAuthToken(refreshTokenCookie string, refreshFunc RefreshTokenFunc) (*schema.TokenResponse, error) {
	if refreshTokenCookie == "" {
		return nil, jwxError("refresh token is not set")
	}

	refreshToken, err := jwt.Parse([]byte(refreshTokenCookie), jwt.WithKeySet(*s.jwks))
	if err != nil {
		return nil, jwxError("failed to verify refresh token")
	}

	now := time.Now()
	if refreshToken.Expiration().Before(now) {
		return nil, jwxError("refresh token has expired")
	}

	token, err := refreshFunc(refreshToken)
	if err != nil {
		return nil, err
	}

	return s.AuthTokenResponse(token)
}

func (s *JWXService) AuthTokenResponse(accessToken jwt.Token) (*schema.TokenResponse, error) {
	signed, err := jwt.Sign(accessToken, DefaultJWA, s.privateJWK)
	if err != nil {
		return nil, http.InternalError(err)
	}

	accessTokenTTL := viper.GetDuration("access_token_ttl")
	return &schema.TokenResponse{
		AccessToken:  string(signed),
		TokenType:    "Bearer",
		ExpiresIn:    int(accessTokenTTL.Seconds()),
		RefreshToken: "HttpOnly",
	}, nil
}

func jwxError(msg string) *shared.ServiceError {
	return &shared.ServiceError{
		Response:             msg,
		StatusCode:           fiber.StatusUnauthorized,
		ErrorCode:            shared.Unauthorized,
		UseResponseAsMessage: shared.Bool(true),
	}
}
