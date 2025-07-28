package main

import (
	"log"

	"boilerplate-go-fiber-v2/internal/middleware"
	"boilerplate-go-fiber-v2/internal/route"
	"boilerplate-go-fiber-v2/pkg/utils"
)

func main() {
	// Initialize application
	app, cfg, err := utils.InitializeApp()
	if err != nil {
		log.Fatal("Failed to initialize app:", err)
	}

	// Setup middleware
	middleware.SetupMiddleware(app)

	// Initialize database and Redis
	db := utils.InitializeDatabase(cfg)
	redis := utils.InitializeRedis(cfg)

	// Setup routes
	route.SetupRoutes(app, db, redis, cfg)

	// Get port
	port := utils.GetPort()

	log.Printf("Server starting on port %s", port)
	if err := app.Listen(":" + port); err != nil {
		log.Fatal("Failed to start server:", err)
	}
}
