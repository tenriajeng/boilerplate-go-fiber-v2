package repository

import (
	"boilerplate-go-fiber-v2/internal/domain/entity"
	"boilerplate-go-fiber-v2/internal/domain/repository"
	"boilerplate-go-fiber-v2/internal/model"
	"context"
	"strings"

	"gorm.io/gorm"
)

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) repository.UserRepository {
	return &userRepository{db: db}
}

func (r *userRepository) Create(ctx context.Context, user *entity.User) error {
	userModel := &model.UserModel{}
	userModel.FromEntity(user)

	if err := r.db.WithContext(ctx).Create(userModel).Error; err != nil {
		return err
	}

	// Update the original entity with the generated ID
	user.ID = userModel.ID
	return nil
}

func (r *userRepository) GetByID(ctx context.Context, id uint) (*entity.User, error) {
	var userModel model.UserModel
	if err := r.db.WithContext(ctx).First(&userModel, id).Error; err != nil {
		return nil, err
	}
	return userModel.ToEntity(), nil
}

func (r *userRepository) GetByEmail(ctx context.Context, email string) (*entity.User, error) {
	var userModel model.UserModel
	if err := r.db.WithContext(ctx).Where("email = ?", email).First(&userModel).Error; err != nil {
		return nil, err
	}
	return userModel.ToEntity(), nil
}

func (r *userRepository) GetByUsername(ctx context.Context, username string) (*entity.User, error) {
	var userModel model.UserModel
	if err := r.db.WithContext(ctx).Where("username = ?", username).First(&userModel).Error; err != nil {
		return nil, err
	}
	return userModel.ToEntity(), nil
}

func (r *userRepository) Update(ctx context.Context, user *entity.User) error {
	userModel := &model.UserModel{}
	userModel.FromEntity(user)
	return r.db.WithContext(ctx).Save(userModel).Error
}

func (r *userRepository) Delete(ctx context.Context, id uint) error {
	return r.db.WithContext(ctx).Delete(&model.UserModel{}, id).Error
}

func (r *userRepository) List(ctx context.Context, filter repository.UserFilter) ([]*entity.User, error) {
	var userModels []model.UserModel
	query := r.db.WithContext(ctx)

	// Apply filters
	if filter.Search != "" {
		searchTerm := "%" + strings.ToLower(filter.Search) + "%"
		query = query.Where("LOWER(first_name) LIKE ? OR LOWER(last_name) LIKE ? OR LOWER(email) LIKE ? OR LOWER(username) LIKE ?",
			searchTerm, searchTerm, searchTerm, searchTerm)
	}

	if filter.Role != "" {
		query = query.Where("role = ?", filter.Role)
	}

	if filter.Status != "" {
		query = query.Where("status = ?", filter.Status)
	}

	// Apply sorting
	sortBy := filter.SortBy
	if sortBy == "" {
		sortBy = "created_at"
	}
	if filter.SortDesc {
		sortBy += " DESC"
	}
	query = query.Order(sortBy)

	// Apply pagination
	if filter.Page > 0 && filter.Limit > 0 {
		offset := (filter.Page - 1) * filter.Limit
		query = query.Offset(offset).Limit(filter.Limit)
	}

	if err := query.Find(&userModels).Error; err != nil {
		return nil, err
	}

	users := make([]*entity.User, len(userModels))
	for i, userModel := range userModels {
		users[i] = userModel.ToEntity()
	}
	return users, nil
}

func (r *userRepository) Count(ctx context.Context, filter repository.UserFilter) (int64, error) {
	var count int64
	query := r.db.WithContext(ctx).Model(&model.UserModel{})

	// Apply filters
	if filter.Search != "" {
		searchTerm := "%" + strings.ToLower(filter.Search) + "%"
		query = query.Where("LOWER(first_name) LIKE ? OR LOWER(last_name) LIKE ? OR LOWER(email) LIKE ? OR LOWER(username) LIKE ?",
			searchTerm, searchTerm, searchTerm, searchTerm)
	}

	if filter.Role != "" {
		query = query.Where("role = ?", filter.Role)
	}

	if filter.Status != "" {
		query = query.Where("status = ?", filter.Status)
	}

	err := query.Count(&count).Error
	return count, err
}

func (r *userRepository) UpdateLastLogin(ctx context.Context, userID uint) error {
	return r.db.WithContext(ctx).Model(&model.UserModel{}).Where("id = ?", userID).Update("last_login_at", gorm.Expr("NOW()")).Error
}

func (r *userRepository) UpdateStatus(ctx context.Context, userID uint, status string) error {
	return r.db.WithContext(ctx).Model(&model.UserModel{}).Where("id = ?", userID).Update("status", status).Error
}

func (r *userRepository) UpdateTFA(ctx context.Context, userID uint, enabled bool, secret string, backupCodes []string) error {
	return r.db.WithContext(ctx).Model(&model.UserModel{}).Where("id = ?", userID).Updates(map[string]interface{}{
		"tfa_enabled":      enabled,
		"tfa_secret":       secret,
		"tfa_backup_codes": backupCodes,
	}).Error
}
