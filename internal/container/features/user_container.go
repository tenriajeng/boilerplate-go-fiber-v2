package features

import (
	"boilerplate-go-fiber-v2/config"
	"boilerplate-go-fiber-v2/internal/domain/repository"
	domainService "boilerplate-go-fiber-v2/internal/domain/service"
	repo "boilerplate-go-fiber-v2/internal/repository"
	"boilerplate-go-fiber-v2/internal/service"

	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

// UserContainer holds user-related dependencies
type UserContainer struct {
	// Repositories
	UserRepo repository.UserRepository

	// Services
	UserService domainService.UserService

	// Handlers
	// UserHandler *handler.UserHandler // TODO: Implement when needed
}

// NewUserContainer creates user container
func NewUserContainer(db *gorm.DB, redis *redis.Client, cfg *config.Config) *UserContainer {
	container := &UserContainer{}

	// Initialize repositories
	if db != nil {
		container.UserRepo = repo.NewUserRepository(db)
	}

	// Initialize services
	if container.UserRepo != nil {
		container.UserService = service.NewUserService(container.UserRepo)
	}

	// Initialize handlers (when UserHandler is implemented)
	// if container.UserService != nil {
	//     container.UserHandler = handler.NewUserHandler(container.UserService)
	// }

	return container
}

// GetUserService returns user service
func (c *UserContainer) GetUserService() domainService.UserService {
	return c.UserService
}

// GetUserHandler returns user handler (when implemented)
// func (c *UserContainer) GetUserHandler() *handler.UserHandler {
//     return c.UserHandler
// }
