package handlers

import (
	"crypto/rsa"
	sq "github.com/Masterminds/squirrel"
	"github.com/imanhodjaev/getout/platform/repo"
	"github.com/imanhodjaev/getout/platform/services"
	"github.com/jmoiron/sqlx"
)

type Handler struct {
	BaseRepo      *repo.Repo
	UserRepo      repo.UserRepo
	UserService   services.UserService
	AuthService   services.AuthService
	JWXService    *services.JWXService
	Params        *ParamHandler
	SigningSecret string
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

	userRepo := repo.NewUserRepo(baseRepo)
	userService := services.NewUserService(userRepo)
	jwxService, err := services.NewJWXService(key)
	if err != nil {
		return nil, err
	}

	return &Handler{
		BaseRepo:    baseRepo,
		UserRepo:    userRepo,
		UserService: userService,
		AuthService: services.NewAuthService(userService, jwxService),
		JWXService:  jwxService,
		Params: &ParamHandler{
			UserService: userService,
		},
	}, nil
}
