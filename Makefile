# Makefile for Boilerplate Go Fiber v2

.PHONY: help build run test clean migrate-up migrate-down migrate-status migrate-create migrate-force migrate-wipe

# Default target
help:
	@echo "🚀 Boilerplate Go Fiber v2 - Available Commands:"
	@echo ""
	@echo "📦 Build & Run:"
	@echo "  make build          # Build the application"
	@echo "  make run            # Run the application"
	@echo "  make clean          # Clean build artifacts"
	@echo ""
	@echo "🗄️  Database Migrations:"
	@echo "  make migrate-up     # Run all pending migrations"
	@echo "  make migrate-down   # Rollback all migrations"
	@echo "  make migrate-status # Show migration status"
	@echo "  make migrate-create # Create new migration (usage: make migrate-create NAME=migration_name)"
	@echo "  make migrate-force  # Force migration to version (usage: make migrate-force VERSION=1)"
	@echo "  make migrate-wipe   # Wipe all data and recreate schema (DANGEROUS!)"
	@echo ""
	@echo "🧪 Testing:"
	@echo "  make test           # Run tests"
	@echo "  make test-coverage  # Run tests with coverage"
	@echo ""
	@echo "🔧 Development:"
	@echo "  make dev            # Run in development mode with hot reload"
	@echo "  make lint           # Run linter"
	@echo "  make fmt            # Format code"
	@echo ""

# Build the application
build:
	@echo "🔨 Building application..."
	go build -o bin/app cmd/main.go
	@echo "✅ Build completed!"

# Run the application
run: build
	@echo "🚀 Starting application..."
	./bin/app

# Clean build artifacts
clean:
	@echo "🧹 Cleaning build artifacts..."
	rm -rf bin/
	@echo "✅ Clean completed!"

# Database migrations
migrate-up:
	@echo "📈 Running migrations up..."
	go run cmd/migrate/main.go -action=up

migrate-down:
	@echo "📉 Rolling back migrations..."
	go run cmd/migrate/main.go -action=down

migrate-status:
	@echo "📊 Migration status:"
	go run cmd/migrate/main.go -action=status

migrate-create:
	@if [ -z "$(NAME)" ]; then \
		echo "❌ Error: NAME parameter is required"; \
		echo "Usage: make migrate-create NAME=migration_name"; \
		exit 1; \
	fi
	@echo "📝 Creating migration: $(NAME)"
	go run cmd/migrate/main.go -action=create -name=$(NAME)

migrate-force:
	@if [ -z "$(VERSION)" ]; then \
		echo "❌ Error: VERSION parameter is required"; \
		echo "Usage: make migrate-force VERSION=1"; \
		exit 1; \
	fi
	@echo "🔧 Forcing migration to version: $(VERSION)"
	go run cmd/migrate/main.go -action=force -version=$(VERSION)

migrate-wipe:
	@echo "⚠️  WARNING: This will DELETE ALL DATA in the database!"
	@echo "📊 Database: boilerplate_go_fiber_v2"
	@echo ""
	@read -p "Are you sure you want to continue? (yes/no): " confirm; \
	if [ "$$confirm" != "yes" ]; then \
		echo "❌ Operation cancelled"; \
		exit 1; \
	fi
	@echo "🗑️  Wiping database and recreating schema..."
	go run cmd/migrate/main.go -action=wipe -confirm

# Testing
test:
	@echo "🧪 Running tests..."
	go test ./...

test-coverage:
	@echo "🧪 Running tests with coverage..."
	go test -coverprofile=coverage.out ./...
	go tool cover -html=coverage.out -o coverage.html
	@echo "📊 Coverage report generated: coverage.html"

# Development
dev:
	@echo "🔥 Starting development server with hot reload..."
	@if command -v air > /dev/null; then \
		air; \
	else \
		echo "⚠️  Air not installed. Installing air..."; \
		go install github.com/cosmtrek/air@latest; \
		air; \
	fi

# Code quality
lint:
	@echo "🔍 Running linter..."
	@if command -v golangci-lint > /dev/null; then \
		golangci-lint run; \
	else \
		echo "⚠️  golangci-lint not installed. Installing..."; \
		go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest; \
		golangci-lint run; \
	fi

fmt:
	@echo "🎨 Formatting code..."
	go fmt ./...
	go vet ./...

# Database setup
setup-db:
	@echo "🗄️  Setting up database..."
	@if [ -f "scripts/setup_db.sh" ]; then \
		chmod +x scripts/setup_db.sh; \
		./scripts/setup_db.sh; \
	else \
		echo "⚠️  setup_db.sh not found"; \
	fi

# Install dependencies
deps:
	@echo "📦 Installing dependencies..."
	go mod download
	go mod tidy

# Generate documentation
docs:
	@echo "📚 Generating documentation..."
	@if command -v swag > /dev/null; then \
		swag init -g cmd/main.go; \
	else \
		echo "⚠️  swag not installed. Installing..."; \
		go install github.com/swaggo/swag/cmd/swag@latest; \
		swag init -g cmd/main.go; \
	fi

# Docker commands
docker-build:
	@echo "🐳 Building Docker image..."
	docker build -t boilerplate-go-fiber-v2 .

docker-run:
	@echo "🐳 Running Docker container..."
	docker run -p 8080:8080 boilerplate-go-fiber-v2

docker-clean:
	@echo "🧹 Cleaning Docker images..."
	docker rmi boilerplate-go-fiber-v2 || true
