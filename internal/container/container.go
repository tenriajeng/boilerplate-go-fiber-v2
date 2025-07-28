package container

import (
	"boilerplate-go-fiber-v2/config"
	"boilerplate-go-fiber-v2/internal/container/features"
	domainService "boilerplate-go-fiber-v2/internal/domain/service"
	"boilerplate-go-fiber-v2/internal/handler"

	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

// Container holds all application dependencies
type Container struct {
	// Feature containers
	Auth *features.AuthContainer
	User *features.UserContainer

	// Shared dependencies
	DB     *gorm.DB
	Redis  *redis.Client
	Config *config.Config
}

// NewContainer creates and initializes all dependencies
func NewContainer(db *gorm.DB, redis *redis.Client, cfg *config.Config) *Container {
	container := &Container{
		DB:     db,
		Redis:  redis,
		Config: cfg,
	}

	// Initialize feature containers
	container.Auth = features.NewAuthContainer(db, redis, cfg)
	container.User = features.NewUserContainer(db, redis, cfg)

	return container
}

// GetAuthHandler returns auth handler
func (c *Container) GetAuthHandler() *handler.AuthHandler {
	if c.Auth != nil {
		return c.Auth.GetAuthHandler()
	}
	return nil
}

// GetUserService returns user service
func (c *Container) GetUserService() domainService.UserService {
	if c.Auth != nil {
		return c.Auth.GetUserService()
	}
	return nil
}

// GetAuthService returns auth service
func (c *Container) GetAuthService() domainService.AuthService {
	if c.Auth != nil {
		return c.Auth.GetAuthService()
	}
	return nil
}
