package services

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/imanhodjaev/getout/platform/http"
	"github.com/imanhodjaev/getout/platform/repo"
	"github.com/imanhodjaev/getout/platform/schema"
	"github.com/imanhodjaev/getout/util"
	"github.com/lestrrat-go/jwx/jwt"
	"github.com/rs/zerolog/log"
	"github.com/spf13/viper"
	"time"
)

type AuthService interface {
	Login(ctx *fiber.Ctx, loginRequest *schema.LoginRequest) (*schema.TokenResponse, error)
	RefreshAuthToken(ctx *fiber.Ctx) (*schema.TokenResponse, error)
	Logout(ctx *fiber.Ctx) error
}

type authService struct {
	usersRepo  repo.UserRepo
	jwxService *JWXService
}

func NewAuthService(usersRepo repo.UserRepo, jwxService *JWXService) AuthService {
	return &authService{
		usersRepo:  usersRepo,
		jwxService: jwxService,
	}
}

func (a *authService) Login(ctx *fiber.Ctx, loginRequest *schema.LoginRequest) (*schema.TokenResponse, error) {
	if len(loginRequest.Email) <= 3 {
		return nil, http.BadRequestWithMessage("Please provide e-mail")
	}

	if len(loginRequest.Password) <= 3 {
		return nil, http.BadRequestWithMessage("Please provide password")
	}

	userID, err := uuid.Parse(ctx.Params("user_id"))
	if !a.usersRepo.Exists(userID) {
		log.Info().
			Str("user_id", userID.String()).
			Msg(fmt.Sprintf("User not found"))

		return nil, http.UnauthorizedError("E-mail or password wrong")
	}

	user, err := a.usersRepo.Get(userID)
	if err != nil {
		log.Info().
			Str("user_id", userID.String()).
			Msg(fmt.Sprintf("Unable to fetch user"))

		return nil, http.UnauthorizedError("E-mail or password wrong")
	}

	if !user.IsActive {
		return nil, http.InactiveUserError()
	}

	err = util.CheckPassword(user.Password, loginRequest.Password)
	if err != nil {
		return nil, http.UnauthorizedError("E-mail or password wrong")
	}

	// issue access_token (short-lived) and refresh_token (to update it)
	refreshTokenTTL := viper.GetDuration("refresh_token_ttl")
	accessTokenTTL := viper.GetDuration("access_token_ttl")
	now := time.Now()
	refreshToken := jwt.New()

	err = refreshToken.Set(jwt.ExpirationKey, now.Add(refreshTokenTTL))
	if err != nil {
		log.Info().
			Str("user_id", userID.String()).
			Msg(fmt.Sprintf("JWT Refresh token unable to set %s", jwt.ExpirationKey))
	}

	err = refreshToken.Set(jwt.SubjectKey, user.ID)
	if err != nil {
		log.Info().
			Str("user_id", userID.String()).
			Msg(fmt.Sprintf("JWT Refresh token unable to set %s", jwt.SubjectKey))
	}

	// for security reasons we store refresh_token as a secure cookie (which is not in oauth standard)
	// it can be changed when we will have an ability to revoke (or blacklist) them
	refreshTokenCookie, err := a.jwxService.GetRefreshTokenCookie(refreshToken)
	if err != nil {
		return nil, err
	}

	ctx.Cookie(refreshTokenCookie)

	authToken := jwt.New()
	err = authToken.Set(jwt.ExpirationKey, now.Add(accessTokenTTL))
	if err != nil {
		log.Info().
			Str("user_id", userID.String()).
			Msg(fmt.Sprintf("JWT Access token unable to set %s", jwt.ExpirationKey))
	}

	err = authToken.Set(jwt.SubjectKey, user.ID)
	if err != nil {
		log.Info().
			Str("user_id", userID.String()).
			Msg(fmt.Sprintf("JWT Access token unable to set %s", jwt.SubjectKey))
	}

	return a.jwxService.AuthTokenResponse(authToken)
}

func (a *authService) RefreshAuthToken(ctx *fiber.Ctx) (*schema.TokenResponse, error) {
	return a.jwxService.RefreshAuthToken(
		ctx.Cookies(RefreshTokenCookieName, ""),
		func(refreshToken jwt.Token) (jwt.Token, error) {
			userID := uuid.MustParse(refreshToken.Subject())
			if !a.usersRepo.Exists(userID) {
				return nil, http.NotFoundError("User not found")
			}

			accessTokenTTL := viper.GetDuration("oauth_access_token_ttl")
			authToken := jwt.New()
			err := authToken.Set(jwt.ExpirationKey, time.Now().Add(accessTokenTTL))
			if err != nil {
				log.Info().
					Str("user_id", userID.String()).
					Msg(fmt.Sprintf("JWT Access token unable to set %s", jwt.ExpirationKey))
			}

			err = authToken.Set(jwt.SubjectKey, userID)
			if err != nil {
				log.Info().
					Str("user_id", userID.String()).
					Msg(fmt.Sprintf("JWT Access token unable to set %s", jwt.SubjectKey))
			}

			return authToken, nil
		},
	)
}
func (a *authService) JWKS(ctx *fiber.Ctx) error {
	return ctx.JSON(a.jwxService.JWKS())
}

func (a *authService) Logout(ctx *fiber.Ctx) error {
	ctx.ClearCookie(RefreshTokenCookieName)
	return ctx.SendStatus(fiber.StatusNoContent)
}
