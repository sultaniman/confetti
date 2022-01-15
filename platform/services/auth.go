package services

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/imanhodjaev/getout/platform/http"
	"github.com/imanhodjaev/getout/platform/schema"
	"github.com/imanhodjaev/getout/util"
	"github.com/lestrrat-go/jwx/jwt"
	"github.com/rs/zerolog/log"
	"time"
)

const AccessTokenDuration = time.Hour             // 1 hour
const RefreshTokenDuration = 24 * time.Hour * 180 // 180 days

type AuthService interface {
	AccessTokenAuthFlow(ctx *fiber.Ctx, loginRequest *schema.LoginRequest) (*schema.TokenResponse, error)
	RefreshAuthToken(ctx *fiber.Ctx) (*schema.TokenResponse, error)
	Register(registerPayload *schema.RegisterRequest) error
	Logout(ctx *fiber.Ctx) error
}

type authService struct {
	usersService UserService
	jwxService   *JWXService
}

func NewAuthService(usersService UserService, jwxService *JWXService) AuthService {
	return &authService{
		usersService: usersService,
		jwxService:   jwxService,
	}
}

func (a *authService) AccessTokenAuthFlow(ctx *fiber.Ctx, loginRequest *schema.LoginRequest) (*schema.TokenResponse, error) {
	if len(loginRequest.Email) <= 3 {
		return nil, http.BadRequestWithMessage("Please provide e-mail")
	}

	if len(loginRequest.Password) <= 3 {
		return nil, http.BadRequestWithMessage("Please provide password")
	}

	user, err := a.usersService.GetByEmail(loginRequest.Email)
	if err != nil {
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
	now := time.Now()
	refreshToken := jwt.New()
	err = refreshToken.Set(jwt.ExpirationKey, now.Add(RefreshTokenDuration))
	if err != nil {
		log.Info().
			Str("user_id", user.ID.String()).
			Msg(fmt.Sprintf("JWT Refresh token unable to set %s", jwt.ExpirationKey))
	}

	err = refreshToken.Set(jwt.SubjectKey, user.ID)
	if err != nil {
		log.Info().
			Str("user_id", user.ID.String()).
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
	err = authToken.Set(jwt.ExpirationKey, now.Add(AccessTokenDuration))
	if err != nil {
		log.Info().
			Str("user_id", user.ID.String()).
			Msg(fmt.Sprintf("JWT Access token unable to set %s", jwt.ExpirationKey))
	}

	err = authToken.Set(jwt.SubjectKey, user.ID)
	if err != nil {
		log.Info().
			Str("user_id", user.ID.String()).
			Msg(fmt.Sprintf("JWT Access token unable to set %s", jwt.SubjectKey))
	}

	return a.jwxService.AuthTokenResponse(authToken)
}

func (a *authService) RefreshAuthToken(ctx *fiber.Ctx) (*schema.TokenResponse, error) {
	return a.jwxService.RefreshAuthToken(
		ctx.Cookies(RefreshTokenCookieName, ""),
		func(refreshToken jwt.Token) (jwt.Token, error) {
			userID := uuid.MustParse(refreshToken.Subject())
			if !a.usersService.Exists(userID) {
				return nil, http.NotFoundError("User not found")
			}

			authToken := jwt.New()
			err := authToken.Set(jwt.ExpirationKey, time.Now().Add(AccessTokenDuration))
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

func (a *authService) Register(registerPayload *schema.RegisterRequest) error {
	if a.usersService.EmailExists(registerPayload.Email) {
		return http.Conflict("E-mail is taken by someone else")
	}
	_, err := a.usersService.Create(&schema.NewUserRequest{
		FullName: "",
		Email:    registerPayload.Email,
		Password: registerPayload.Password,
		Provider: "auth",
		Settings: []byte("{}"),
	})

	if err != nil {
		return err
	}

	return nil
}

func (a *authService) Logout(ctx *fiber.Ctx) error {
	ctx.ClearCookie(RefreshTokenCookieName)
	return ctx.SendStatus(fiber.StatusNoContent)
}
