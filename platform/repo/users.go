package repo

import (
	sq "github.com/Masterminds/squirrel"
	"github.com/google/uuid"
	"github.com/imanhodjaev/getout/platform/entities"
	"time"
)

//go:generate mockgen -source=users.go -destination=../mocks/users.go -package=mocks
type UserRepo interface {
	Get(id uuid.UUID) (*entities.User, error)
	Create(user *entities.NewUser) (*entities.User, error)
	Delete(id uuid.UUID) (*entities.User, error)
	Update(userId uuid.UUID, user *entities.UpdateUser) (*entities.User, error)
	Exists(userId uuid.UUID) bool
	UpdateEmail(userId uuid.UUID, newEmail string) (*entities.User, error)
	UpdatePassword(userId uuid.UUID, newPassword string) (*entities.User, error)
	GetPasswordHash(id uuid.UUID) (string, error)
}

type userRepo struct {
	Base *Repo
}

func NewUserRepo(base *Repo) UserRepo {
	return &userRepo{
		Base: base,
	}
}

func (r *userRepo) Get(id uuid.UUID) (*entities.User, error) {
	query, args, err := r.Base.
		Select("users").
		Where(sq.Eq{"id": id}).
		ToSql()

	if err != nil {
		return nil, err
	}

	var user entities.User
	return &user, r.Base.DB.Get(user, query, args...)
}

func (r *userRepo) GetPasswordHash(id uuid.UUID) (string, error) {
	query, args, err := r.Base.
		Select("users").
		Where(sq.Eq{"id": id}).
		ToSql()

	if err != nil {
		return "", err
	}

	var user entities.User
	return user.Password, r.Base.DB.Get(user, query, args...)
}

func (r *userRepo) Create(user *entities.NewUser) (*entities.User, error) {
	var userRow *entities.User
	query, args, err := r.Base.
		Insert(
			"users",
			"full_name",
			"username",
			"email",
			"password",
			"is_admin",
			"is_active",
			"settings",
			"provider",
			"created_at",
			"updated_at",
		).
		Values(
			user.FullName,
			user.FullName,
			user.Email,
			user.Password,
			user.IsAdmin,
			user.IsActive,
			user.Settings,
			user.Provider,
			time.Now().UTC(),
			time.Now().UTC(),
		).
		ToSql()

	if err != nil {
		return nil, err
	}

	return userRow, r.Base.DB.Get(userRow, query, args...)
}

func (r *userRepo) Update(userId uuid.UUID, user *entities.UpdateUser) (*entities.User, error) {
	var userRow *entities.User
	querySet := r.Base.
		Update("users", true).
		Where(sq.Eq{"id": userId})

	if user.Username != "" {
		querySet = querySet.Set("username", user.Username)
	}

	if user.FullName != "" {
		querySet = querySet.Set("full_name", user.FullName)
	}

	if user.Settings != nil {
		querySet = querySet.Set("settings", user.Settings)
	}

	query, args, err := querySet.ToSql()
	if err != nil {
		return nil, err
	}

	return userRow, r.Base.DB.Get(userRow, query, args...)
}

func (r *userRepo) Exists(userId uuid.UUID) bool {
	var userCount int
	query, args, err := r.Base.
		Count("users", sq.Eq{"id": userId}).
		ToSql()

	if err != nil {
		return false
	}

	err = r.Base.DB.Get(userCount, query, args...)
	if err != nil {
		return false
	}

	return userCount > 0
}

func (r *userRepo) UpdateEmail(userId uuid.UUID, newEmail string) (*entities.User, error) {
	var userRow *entities.User
	query, args, err := r.Base.
		Update("users", true).
		Where(sq.Eq{"id": userId}).
		Set("email", newEmail).
		ToSql()

	if err != nil {
		return nil, err
	}

	return userRow, r.Base.DB.Get(userRow, query, args...)
}

func (r *userRepo) UpdatePassword(userId uuid.UUID, newPassword string) (*entities.User, error) {
	var userRow *entities.User
	query, args, err := r.Base.
		Update("users", true).
		Where(sq.Eq{"id": userId}).
		Set("password", newPassword).
		ToSql()

	if err != nil {
		return nil, err
	}

	return userRow, r.Base.DB.Get(userRow, query, args...)
}

func (r *userRepo) Delete(id uuid.UUID) (*entities.User, error) {
	var user *entities.User
	query, args, err := r.Base.
		Delete("users", sq.Eq{"id": id}).
		ToSql()

	if err != nil {
		return nil, err
	}

	return user, r.Base.DB.Get(user, query, args...)
}