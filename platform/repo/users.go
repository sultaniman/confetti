package repo

import (
	"errors"
	sq "github.com/Masterminds/squirrel"
	"github.com/google/uuid"
	"github.com/sultaniman/confetti/platform/entities"
	"time"
)

//go:generate mockgen -source=users.go -destination=../mocks/users.go -package=mocks
type UserRepo interface {
	Get(id uuid.UUID) (*entities.User, error)
	GetByEmail(email string) (*entities.User, error)
	Create(user *entities.NewUser) (*entities.User, error)
	Delete(id uuid.UUID) (*entities.User, error)
	Update(userId uuid.UUID, user *entities.UpdateUser) (*entities.User, error)
	Exists(userId uuid.UUID) bool
	EmailExists(email string) bool
	UpdateEmail(userId uuid.UUID, newEmail string) (*entities.User, error)
	UpdatePassword(userId uuid.UUID, newPassword string) (*entities.User, error)
	CreateActionCode(actionCodeRequest *entities.ActionCodeRequest) (*entities.ActionCode, error)
	CheckActionCode(actionCodeCheck *entities.ActionCodeCheck) error
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

	user := new(entities.User)
	return user, r.Base.DB.Get(user, query, args...)
}

func (r *userRepo) GetByEmail(email string) (*entities.User, error) {
	query, args, err := r.Base.
		Select("users").
		Where(sq.Eq{"email": email}).
		ToSql()

	if err != nil {
		return nil, err
	}

	user := new(entities.User)
	return user, r.Base.DB.Get(user, query, args...)
}

func (r *userRepo) Create(user *entities.NewUser) (*entities.User, error) {
	query, args, err := r.Base.
		Insert(
			"users",
			"full_name",
			"email",
			"password",
			"is_admin",
			"is_active",
			"is_confirmed",
			"settings",
			"provider",
			"created_at",
			"updated_at",
		).
		Values(
			user.FullName,
			user.Email,
			user.Password,
			user.IsAdmin,
			user.IsActive,
			user.IsConfirmed,
			user.Settings,
			user.Provider,
			time.Now().UTC(),
			time.Now().UTC(),
		).
		ToSql()

	if err != nil {
		return nil, err
	}

	userRow := new(entities.User)
	return userRow, r.Base.DB.Get(userRow, query, args...)
}

func (r *userRepo) Update(userId uuid.UUID, user *entities.UpdateUser) (*entities.User, error) {
	querySet := r.Base.
		Update("users", true).
		Where(sq.Eq{"id": userId})

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

	userRow := new(entities.User)
	return userRow, r.Base.DB.Get(userRow, query, args...)
}

func (r *userRepo) Exists(userId uuid.UUID) bool {
	query, args, err := r.Base.
		Count("users", sq.Eq{"id": userId}).
		ToSql()

	if err != nil {
		return false
	}

	userCount := 0
	err = r.Base.DB.Get(&userCount, query, args...)
	if err != nil {
		return false
	}

	return userCount > 0
}

func (r *userRepo) EmailExists(email string) bool {
	query, args, err := r.Base.
		Count("users", sq.Eq{"email": email}).
		ToSql()

	if err != nil {
		return false
	}

	userCount := 0
	err = r.Base.DB.Get(&userCount, query, args...)
	if err != nil {
		return false
	}

	return userCount > 0
}

func (r *userRepo) UpdateEmail(userId uuid.UUID, newEmail string) (*entities.User, error) {
	query, args, err := r.Base.
		Update("users", true).
		Where(sq.Eq{"id": userId}).
		Set("email", newEmail).
		ToSql()

	if err != nil {
		return nil, err
	}

	userRow := new(entities.User)
	return userRow, r.Base.DB.Get(userRow, query, args...)
}

func (r *userRepo) UpdatePassword(userId uuid.UUID, newPassword string) (*entities.User, error) {
	query, args, err := r.Base.
		Update("users", true).
		Where(sq.Eq{"id": userId}).
		Set("password", newPassword).
		ToSql()

	if err != nil {
		return nil, err
	}

	userRow := new(entities.User)
	return userRow, r.Base.DB.Get(userRow, query, args...)
}

func (r *userRepo) CreateActionCode(actionCodeRequest *entities.ActionCodeRequest) (*entities.ActionCode, error) {
	user, err := r.GetByEmail(actionCodeRequest.Email)
	if err != nil {
		return nil, err
	}

	query, args, err := r.Base.
		Insert(
			actionCodeRequest.Type,
			"user_id",
			"code",
			"created_at",
		).
		Values(
			user.ID,
			uuid.New().String(),
			time.Now().UTC(),
		).
		ToSql()

	if err != nil {
		return nil, err
	}

	actionCode := new(entities.ActionCode)
	return actionCode, r.Base.DB.Get(actionCode, query, args...)
}

func (r *userRepo) CheckActionCode(actionCodeCheck *entities.ActionCodeCheck) error {
	query, args, err := r.Base.
		Select(actionCodeCheck.Type).
		Where(sq.Eq{"code": actionCodeCheck.Code}).
		ToSql()

	if err != nil {
		return err
	}

	actionCode := new(entities.ActionCode)
	err = r.Base.DB.Get(actionCode, query, args...)
	if err != nil {
		return err
	}

	if actionCode.CreatedAt.Add(actionCodeCheck.TTL).Before(time.Now().UTC()) {
		return errors.New("code has expired")
	}

	return nil
}

func (r *userRepo) Delete(id uuid.UUID) (*entities.User, error) {
	query, args, err := r.Base.
		Delete("users", sq.Eq{"id": id}).
		ToSql()

	if err != nil {
		return nil, err
	}

	user := new(entities.User)
	return user, r.Base.DB.Get(user, query, args...)
}
