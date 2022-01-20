package services

import (
	"fmt"
	"github.com/google/uuid"
	"github.com/imanhodjaev/confetti/platform/entities"
	"github.com/imanhodjaev/confetti/platform/http"
	"github.com/imanhodjaev/confetti/platform/repo"
	"github.com/imanhodjaev/confetti/platform/schema"
	"github.com/imanhodjaev/confetti/util"
	"github.com/omeid/pgerror"
	"github.com/rs/zerolog/log"
)

type UserService interface {
	Get(id uuid.UUID) (*schema.UserResponse, error)
	GetByEmail(email string) (*schema.UserResponse, error)
	Create(user *schema.NewUserRequest) (*schema.UserResponse, error)
	Update(userId uuid.UUID, user *schema.UpdateUserRequest) (*schema.UserResponse, error)
	UpdateEmail(userId uuid.UUID, user *schema.UpdateUserEmailRequest) (*schema.UserResponse, error)
	UpdatePassword(userId uuid.UUID, emailUpdate *schema.UpdateUserPasswordRequest) (*schema.UserResponse, error)
	Delete(id uuid.UUID) (*schema.UserResponse, error)
	Exists(userId uuid.UUID) bool
	EmailExists(email string) bool
}

type userService struct {
	usersRepo repo.UserRepo
}

func NewUserService(usersRepo repo.UserRepo) UserService {
	return &userService{
		usersRepo: usersRepo,
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
		FullName: newUserRequest.FullName,
		Email:    newUserRequest.Email,
		Password: password,
		IsActive: true,
		Settings: newUserRequest.Settings,
		Provider: newUserRequest.Provider,
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
