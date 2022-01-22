package handlers

import (
	"crypto/rsa"
	sq "github.com/Masterminds/squirrel"
	"github.com/imanhodjaev/confetti/platform/mailer"
	"github.com/imanhodjaev/confetti/platform/repo"
	"github.com/imanhodjaev/confetti/platform/services"
	"github.com/jmoiron/sqlx"
)

type Handler struct {
	BaseRepo    *repo.Repo
	UserRepo    repo.UserRepo
	UserService services.UserService
	CardService services.CardService
	AuthService services.AuthService
	JWXService  *services.JWXService
	Params      *ParamHandler
}

func NewHandler(db *sqlx.DB, key *rsa.PrivateKey) (*Handler, error) {
	psql := sq.
		StatementBuilder.
		PlaceholderFormat(sq.Dollar).
		RunWith(db)

	baseRepo := &repo.Repo{
		DB: db,
		Q:  &psql,
	}

	mailerHandler := mailer.GetMailer()
	userRepo := repo.NewUserRepo(baseRepo)
	cardRepo := repo.NewCardRepo(baseRepo)
	userService := services.NewUserService(userRepo)
	cardService := services.NewCardService(userRepo, cardRepo, key)
	jwxService, err := services.NewJWXService(key)
	if err != nil {
		return nil, err
	}

	return &Handler{
		BaseRepo:    baseRepo,
		UserRepo:    userRepo,
		UserService: userService,
		CardService: cardService,
		AuthService: services.NewAuthService(userService, jwxService, mailerHandler),
		JWXService:  jwxService,
		Params: &ParamHandler{
			UserService: userService,
			CardService: cardService,
		},
	}, nil
}
