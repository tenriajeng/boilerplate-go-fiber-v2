package features

import (
	"boilerplate-go-fiber-v2/config"
	"boilerplate-go-fiber-v2/internal/domain/repository"
	domainService "boilerplate-go-fiber-v2/internal/domain/service"
	"boilerplate-go-fiber-v2/internal/handler"
	repo "boilerplate-go-fiber-v2/internal/repository"
	"boilerplate-go-fiber-v2/internal/service"

	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

// AuthContainer holds auth-related dependencies
type AuthContainer struct {
	// Repositories
	UserRepo repository.UserRepository
	AuthRepo repository.AuthRepository

	// Services
	UserService domainService.UserService
	AuthService domainService.AuthService

	// Handlers
	AuthHandler *handler.AuthHandler
}

// NewAuthContainer creates auth container
func NewAuthContainer(db *gorm.DB, redis *redis.Client, cfg *config.Config) *AuthContainer {
	container := &AuthContainer{}

	// Initialize repositories
	if db != nil {
		container.UserRepo = repo.NewUserRepository(db)
		container.AuthRepo = repo.NewAuthRepository(db)
	}

	// Initialize services
	if container.UserRepo != nil {
		container.UserService = service.NewUserService(container.UserRepo)
		container.AuthService = service.NewAuthService(container.UserRepo, container.AuthRepo, container.UserService, cfg)
	}

	// Initialize handlers
	if container.AuthService != nil && container.UserService != nil {
		container.AuthHandler = handler.NewAuthHandler(container.AuthService, container.UserService)
	}

	return container
}

// GetAuthHandler returns auth handler
func (c *AuthContainer) GetAuthHandler() *handler.AuthHandler {
	return c.AuthHandler
}

// GetUserService returns user service
func (c *AuthContainer) GetUserService() domainService.UserService {
	return c.UserService
}

// GetAuthService returns auth service
func (c *AuthContainer) GetAuthService() domainService.AuthService {
	return c.AuthService
}
