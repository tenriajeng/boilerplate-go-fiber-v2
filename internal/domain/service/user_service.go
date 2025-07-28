package service

import (
	"context"

	"boilerplate-go-fiber-v2/internal/domain/entity"
	"boilerplate-go-fiber-v2/internal/domain/repository"
)

type UserService interface {
	Register(ctx context.Context, user *entity.User) error
	GetByID(ctx context.Context, id uint) (*entity.User, error)
	GetByEmail(ctx context.Context, email string) (*entity.User, error)
	Update(ctx context.Context, user *entity.User) error
	Delete(ctx context.Context, id uint) error
	List(ctx context.Context, filter repository.UserFilter) ([]*entity.User, error)
	Count(ctx context.Context, filter repository.UserFilter) (int64, error)
	UpdateProfile(ctx context.Context, userID uint, updates map[string]interface{}) error
	ChangePassword(ctx context.Context, userID uint, oldPassword, newPassword string) error
	UpdateStatus(ctx context.Context, userID uint, status string) error
	VerifyEmail(ctx context.Context, userID uint) error
	UpdateLastLogin(ctx context.Context, userID uint) error
}
