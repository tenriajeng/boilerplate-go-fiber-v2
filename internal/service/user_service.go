package service

import (
	"context"
	"errors"
	"time"

	"boilerplate-go-fiber-v2/internal/domain/entity"
	"boilerplate-go-fiber-v2/internal/domain/repository"
	"boilerplate-go-fiber-v2/internal/domain/service"
	"boilerplate-go-fiber-v2/pkg/utils"
)

type userService struct {
	userRepo repository.UserRepository
}

// NewUserService creates a new user service
func NewUserService(userRepo repository.UserRepository) service.UserService {
	return &userService{
		userRepo: userRepo,
	}
}

// Register registers a new user
func (s *userService) Register(ctx context.Context, user *entity.User) error {
	// Hash password
	hashedPassword, err := utils.HashPassword(user.Password)
	if err != nil {
		return err
	}
	user.Password = hashedPassword

	// Set default values
	user.Role = "user"
	user.Status = "active"
	user.CreatedAt = time.Now()
	user.UpdatedAt = time.Now()

	// Create user
	return s.userRepo.Create(ctx, user)
}

// GetByID gets a user by ID
func (s *userService) GetByID(ctx context.Context, id uint) (*entity.User, error) {
	return s.userRepo.GetByID(ctx, id)
}

// GetByEmail gets a user by email
func (s *userService) GetByEmail(ctx context.Context, email string) (*entity.User, error) {
	return s.userRepo.GetByEmail(ctx, email)
}

// Update updates a user
func (s *userService) Update(ctx context.Context, user *entity.User) error {
	user.UpdatedAt = time.Now()
	return s.userRepo.Update(ctx, user)
}

// Delete deletes a user
func (s *userService) Delete(ctx context.Context, id uint) error {
	return s.userRepo.Delete(ctx, id)
}

// List gets users with filtering
func (s *userService) List(ctx context.Context, filter repository.UserFilter) ([]*entity.User, error) {
	return s.userRepo.List(ctx, filter)
}

// Count counts users with filtering
func (s *userService) Count(ctx context.Context, filter repository.UserFilter) (int64, error) {
	return s.userRepo.Count(ctx, filter)
}

// UpdateProfile updates user profile
func (s *userService) UpdateProfile(ctx context.Context, userID uint, updates map[string]interface{}) error {
	user, err := s.userRepo.GetByID(ctx, userID)
	if err != nil {
		return err
	}

	// Update allowed fields only
	if firstName, ok := updates["first_name"].(string); ok {
		user.FirstName = firstName
	}
	if lastName, ok := updates["last_name"].(string); ok {
		user.LastName = lastName
	}
	if phone, ok := updates["phone"].(string); ok {
		user.Phone = phone
	}
	if avatar, ok := updates["avatar"].(string); ok {
		user.Avatar = avatar
	}

	user.UpdatedAt = time.Now()
	return s.userRepo.Update(ctx, user)
}

// ChangePassword changes user password
func (s *userService) ChangePassword(ctx context.Context, userID uint, oldPassword, newPassword string) error {
	user, err := s.userRepo.GetByID(ctx, userID)
	if err != nil {
		return err
	}

	// Verify old password
	if !utils.CheckPassword(oldPassword, user.Password) {
		return errors.New("invalid old password")
	}

	// Hash new password
	hashedPassword, err := utils.HashPassword(newPassword)
	if err != nil {
		return err
	}

	user.Password = hashedPassword
	user.UpdatedAt = time.Now()
	return s.userRepo.Update(ctx, user)
}

// UpdateStatus updates user status
func (s *userService) UpdateStatus(ctx context.Context, userID uint, status string) error {
	// Validate status
	validStatuses := []string{"active", "inactive", "suspended", "banned"}
	isValid := false
	for _, validStatus := range validStatuses {
		if status == validStatus {
			isValid = true
			break
		}
	}
	if !isValid {
		return errors.New("invalid status")
	}

	return s.userRepo.UpdateStatus(ctx, userID, status)
}

// VerifyEmail verifies user email
func (s *userService) VerifyEmail(ctx context.Context, userID uint) error {
	user, err := s.userRepo.GetByID(ctx, userID)
	if err != nil {
		return err
	}

	user.MarkEmailVerified()
	return s.userRepo.Update(ctx, user)
}

// UpdateLastLogin updates user's last login time
func (s *userService) UpdateLastLogin(ctx context.Context, userID uint) error {
	return s.userRepo.UpdateLastLogin(ctx, userID)
}
