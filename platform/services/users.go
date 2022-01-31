package services

import (
	"database/sql"
	"fmt"
	"github.com/google/uuid"
	"github.com/omeid/pgerror"
	"github.com/rs/zerolog/log"
	"github.com/sultaniman/confetti/platform/entities"
	"github.com/sultaniman/confetti/platform/http"
	"github.com/sultaniman/confetti/platform/mailer"
	"github.com/sultaniman/confetti/platform/repo"
	"github.com/sultaniman/confetti/platform/schema"
	"github.com/sultaniman/confetti/util"
	"time"
)

const (
	UserConfirmationTTL = time.Hour
	PasswordResetTTL    = 15 * time.Minute
)

type UserService interface {
	Get(id uuid.UUID) (*schema.UserResponse, error)
	GetByEmail(email string) (*schema.UserResponse, error)
	Create(user *schema.NewUserRequest) (*schema.UserResponse, error)
	Update(userId uuid.UUID, user *schema.UpdateUserRequest) (*schema.UserResponse, error)
	UpdateEmail(userId uuid.UUID, user *schema.UpdateUserEmailRequest) (*schema.UserResponse, error)
	UpdatePassword(userId uuid.UUID, emailUpdate *schema.UpdateUserPasswordRequest) (*schema.UserResponse, error)
	ResetPasswordRequest(email string) (*schema.ActionCode, error)
	CreateConfirmation(userId uuid.UUID) (*schema.ActionCode, error)
	ResendConfirmation(userId uuid.UUID) error
	ConfirmUser(code string) error
	Delete(id uuid.UUID) (*schema.UserResponse, error)
	Exists(userId uuid.UUID) bool
	EmailExists(email string) bool
}

type userService struct {
	usersRepo   repo.UserRepo
	mailHandler mailer.Mailer
}

func NewUserService(usersRepo repo.UserRepo, mailHandler mailer.Mailer) UserService {
	return &userService{
		usersRepo:   usersRepo,
		mailHandler: mailHandler,
	}
}

func (s *userService) Get(userId uuid.UUID) (*schema.UserResponse, error) {
	if !s.usersRepo.Exists(userId) {
		log.Info().
			Str("user_id", userId.String()).
			Msg(fmt.Sprintf("User not found"))

		return nil, http.NotFoundError("User not found")
	}

	user, err := s.usersRepo.Get(userId)
	if err != nil {
		log.Info().
			Str("user_id", userId.String()).
			Msg(fmt.Sprintf("Unable to fetch user"))

		return nil, http.FetchError(err, "Unable to fetch user")
	}

	return s.userToResponse(user), nil
}

func (s *userService) GetByEmail(email string) (*schema.UserResponse, error) {
	if !s.usersRepo.EmailExists(email) {
		log.Info().
			Str("email", email).
			Msg(fmt.Sprintf("User not found"))

		return nil, http.NotFoundError("User not found")
	}

	user, err := s.usersRepo.GetByEmail(email)
	if err != nil {
		log.Info().
			Str("email", email).
			Msg(fmt.Sprintf("Unable to fetch user"))

		return nil, http.FetchError(err, "Unable to fetch user")
	}

	return s.userToResponse(user), nil
}

func (s *userService) Create(newUserRequest *schema.NewUserRequest) (*schema.UserResponse, error) {
	if !util.IsStrongPassword(newUserRequest.Password) {
		log.Warn().Msg("Password is weak")
		return nil, http.InsecurePasswordError()
	}

	password, err := util.HashPassword(newUserRequest.Password)
	if err != nil {
		log.Error().
			Err(err).
			Msg(fmt.Sprintf("Unable to generate password for user '%s'", newUserRequest.Email))

		return nil, http.InternalError(err)
	}

	user, err := s.usersRepo.Create(&entities.NewUser{
		FullName:    newUserRequest.FullName,
		Email:       newUserRequest.Email,
		Password:    password,
		IsActive:    true,
		IsConfirmed: false,
		Settings:    newUserRequest.Settings,
		Provider:    newUserRequest.Provider,
	})

	if err != nil {
		if e := pgerror.UniqueViolation(err); e != nil {
			log.Error().
				Err(err).
				Msg("Username or email already exists")

			return nil, http.Conflict("Username or email already exists")
		} else {
			log.Error().
				Err(err).
				Msg("Unable to create user")

			return nil, http.InternalError(err)
		}
	}

	response := s.userToResponse(user)
	return response, nil
}

func (s *userService) Update(userId uuid.UUID, userUpdate *schema.UpdateUserRequest) (*schema.UserResponse, error) {
	if !s.usersRepo.Exists(userId) {
		log.Info().
			Str("user_id", userId.String()).
			Msg(fmt.Sprintf("User not found"))

		return nil, http.NotFoundError("User not found")
	}

	user, err := s.usersRepo.Update(userId, &entities.UpdateUser{
		FullName: userUpdate.FullName,
		Settings: userUpdate.Settings,
	})

	if err != nil {
		if e := pgerror.UniqueViolation(err); e != nil {
			log.Error().
				Err(err).
				Msg("Username already exists")

			return nil, http.Conflict("Username already exists")
		} else {
			log.Error().
				Err(err).
				Str("user_id", userId.String()).
				Msg(fmt.Sprintf("Unable to update user"))

			return nil, http.InternalError(err)
		}
	}

	response := s.userToResponse(user)
	return response, nil
}

func (s *userService) UpdateEmail(userId uuid.UUID, emailUpdate *schema.UpdateUserEmailRequest) (*schema.UserResponse, error) {
	if !s.usersRepo.Exists(userId) {
		log.Info().
			Str("user_id", userId.String()).
			Msg(fmt.Sprintf("User not found"))

		return nil, http.NotFoundError("User not found")
	}

	user, err := s.usersRepo.Get(userId)
	if err != nil {
		if e := pgerror.UniqueViolation(err); e != nil {
			log.Error().
				Err(err).
				Msg("Email already exists")

			return nil, http.Conflict("Email already exists")
		} else {
			log.Error().
				Err(err).
				Str("user_id", userId.String()).
				Msg(fmt.Sprintf("Unable to get password hash for user"))

			return nil, err
		}
	}

	err = util.CheckPassword(user.Password, emailUpdate.Password)
	if err != nil {
		return nil, http.InvalidPasswordError()
	}

	updatedUser, err := s.usersRepo.UpdateEmail(userId, emailUpdate.Email)
	if err != nil {
		log.Error().
			Err(err).
			Str("user_id", userId.String()).
			Msg(fmt.Sprintf("Unable to update user"))

		return nil, http.InternalError(err)
	}

	response := s.userToResponse(updatedUser)
	return response, nil
}

func (s *userService) UpdatePassword(userId uuid.UUID, passwordUpdate *schema.UpdateUserPasswordRequest) (*schema.UserResponse, error) {
	if !s.usersRepo.Exists(userId) {
		return nil, http.NotFoundError("User not found")
	}

	if !util.IsStrongPassword(passwordUpdate.NewPassword) {
		log.Warn().Msg("New password is weak")
		return nil, http.InsecurePasswordError()
	}

	user, err := s.usersRepo.Get(userId)
	if err != nil {
		log.Error().
			Err(err).
			Str("user_id", userId.String()).
			Msg(fmt.Sprintf("Unable to get password hash for user"))

		return nil, http.InternalError(err)
	}

	err = util.CheckPassword(user.Password, passwordUpdate.OldPassword)
	if err != nil {
		return nil, http.InvalidPasswordError()
	}

	password, err := util.HashPassword(passwordUpdate.NewPassword)
	if err != nil {
		log.Error().
			Err(err).
			Msg(fmt.Sprintf("Unable to generate password for user '%s'", userId.String()))

		return nil, http.InternalError(err)
	}

	updatedUser, err := s.usersRepo.UpdatePassword(userId, password)
	if err != nil {
		log.Error().
			Err(err).
			Str("user_id", userId.String()).
			Msg(fmt.Sprintf("Unable to update password for user"))

		return nil, http.InternalError(err)
	}

	return s.userToResponse(updatedUser), nil
}

func (s *userService) ConfirmUser(code string) error {
	actionCode, err := s.usersRepo.GetActionCode(&entities.ActionCodeCheck{
		Type: entities.UserConfirmations,
		Code: code,
	})

	if err != nil {
		return http.NotFoundError("Confirmation code not found")
	}

	if actionCode.CreatedAt.Add(UserConfirmationTTL).Before(time.Now().UTC()) {
		return http.Conflict("Confirmation code has already expired")
	}

	_, err = s.usersRepo.ConfirmUser(actionCode.UserId)
	if err != nil {
		log.Error().
			Err(err).
			Str("user_id", actionCode.UserId.String()).
			Msg(fmt.Sprintf("Unable to confirm user"))

		return http.InternalErrorWithMessage("Something went wrong during conifmation")
	}

	return nil
}

func (s *userService) Delete(userId uuid.UUID) (*schema.UserResponse, error) {
	if !s.usersRepo.Exists(userId) {
		return nil, http.NotFoundError("User not found")
	}

	user, err := s.usersRepo.Delete(userId)
	if err != nil {
		log.Error().
			Err(err).
			Str("user_id", userId.String()).
			Msg(fmt.Sprintf("Unable to delete user"))

		return nil, http.InternalError(err)
	}

	response := s.userToResponse(user)
	return response, nil
}

func (s *userService) ResetPasswordRequest(email string) (*schema.ActionCode, error) {
	if !s.usersRepo.EmailExists(email) {
		return nil, http.NotFoundError("User not found")
	}

	passwordReset, err := s.usersRepo.CreateActionCode(&entities.ActionCodeRequest{
		Type:  string(entities.PasswordResets),
		Email: email,
	})

	if err != nil {
		log.Error().
			Err(err).
			Str("email", email).
			Msg("Unable to reset password")

		return nil, http.InternalError(err)
	}

	return &schema.ActionCode{
		ID:        passwordReset.ID,
		UserId:    passwordReset.UserId,
		Code:      passwordReset.Code,
		CreatedAt: passwordReset.CreatedAt,
	}, nil
}

func (s *userService) CreateConfirmation(userId uuid.UUID) (*schema.ActionCode, error) {
	user, err := s.usersRepo.Get(userId)
	if err != nil {
		return nil, s.handleError(err)
	}

	if user.IsConfirmed {
		return nil, http.Conflict("User has already confirmed account")
	}

	actionCode, err := s.usersRepo.CreateActionCode(&entities.ActionCodeRequest{
		Type:  "user_confirmations",
		Email: user.Email,
	})

	if err != nil {
		log.Error().
			Err(err).
			Str("email", user.Email).
			Msg("Unable to create user confirmation")

		return nil, s.handleError(err)
	}

	return &schema.ActionCode{
		ID:        actionCode.ID,
		UserId:    actionCode.UserId,
		Code:      actionCode.Code,
		CreatedAt: actionCode.CreatedAt,
	}, nil
}

func (s *userService) ResendConfirmation(userId uuid.UUID) error {
	user, err := s.Get(userId)
	if err != nil {
		return err
	}

	confirmation, err := s.CreateConfirmation(user.ID)
	if err != nil {
		return err
	}

	err = s.mailHandler.SendConfirmationCode(user.Email, confirmation.Code)
	if err != nil {
		return err
	}

	return nil
}

func (s *userService) Exists(userId uuid.UUID) bool {
	return s.usersRepo.Exists(userId)
}

func (s *userService) EmailExists(email string) bool {
	return s.usersRepo.EmailExists(email)
}

func (s *userService) userToResponse(user *entities.User) *schema.UserResponse {
	return &schema.UserResponse{
		ID:          user.ID,
		FullName:    user.FullName,
		Email:       user.Email,
		IsAdmin:     user.IsAdmin,
		IsActive:    user.IsActive,
		IsConfirmed: user.IsConfirmed,
		Settings:    user.Settings,
		Provider:    user.Provider,
		Password:    user.Password,
		CreatedAt:   user.CreatedAt,
		UpdatedAt:   user.UpdatedAt,
	}
}

func (s *userService) handleError(err error) error {
	if err == sql.ErrNoRows {
		return http.NotFoundError("Card not found")
	} else {
		return http.InternalError(err)
	}
}
