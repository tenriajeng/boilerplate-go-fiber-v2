# Boilerplate Go Fiber v2

A production-ready Go Fiber boilerplate with Clean Architecture, authentication, payment integration, and comprehensive testing.

## 🚀 Features

-   **Clean Architecture** - Separation of concerns with domain-driven design
-   **Authentication & Authorization** - JWT-based auth with role-based access control
-   **Security** - CSRF protection, security headers, rate limiting
-   **Payment Integration** - Xendit and Midtrans payment gateway adapters
-   **Database** - PostgreSQL with GORM and connection pooling
-   **Caching** - Redis single instance with connection pooling
-   **Logging** - Structured logging with Logrus
-   **Configuration** - Viper for advanced config management
-   **Validation** - Go-playground/validator for request validation
-   **Testing** - Unit, integration, and e2e tests with Testify
-   **Docker** - Containerized development and production
-   **CI/CD** - GitHub Actions workflow
-   **Monitoring** - Health checks and structured logging
-   **Migration Management** - CLI tool and Makefile commands for database migrations

## 📁 Project Structure

```
boilerplate-go-fiber-v2/
├── cmd/
│   ├── main.go                      # Entry point aplikasi
│   └── migrate/
│       └── main.go                  # Migration CLI tool
│
├── config/
│   ├── config.go                    # Load .env & config
│   ├── database.go                  # Database connection
│   └── redis.go                     # Redis connection
│
├── internal/
│   ├── domain/                      # Core business logic (Clean Architecture)
│   │   ├── entity/                  # Business entities
│   │   │   ├── user.go
│   │   │   ├── auth.go
│   │   │   ├── payment.go
│   │   │   └── order.go
│   │   ├── repository/              # Repository interfaces (contracts)
│   │   │   ├── user_repository.go
│   │   │   ├── auth_repository.go
│   │   │   ├── payment_repository.go
│   │   │   └── order_repository.go
│   │   └── service/                 # Business logic services
│   │       ├── auth_service.go
│   │       ├── user_service.go
│   │       ├── payment_service.go
│   │       └── order_service.go
│   │
│   ├── handler/                     # HTTP handlers
│   │   ├── auth_handler.go
│   │   ├── user_handler.go
│   │   ├── payment_handler.go
│   │   ├── order_handler.go
│   │   └── health_handler.go
│   │
│   ├── dto/                         # Data Transfer Objects
│   │   ├── auth/
│   │   │   ├── request.go           # Auth request DTOs
│   │   │   └── response.go          # Auth response DTOs
│   │   ├── user/
│   │   │   ├── request.go           # User request DTOs
│   │   │   └── response.go          # User response DTOs
│   │   ├── payment/
│   │   │   ├── request.go           # Payment request DTOs
│   │   │   └── response.go          # Payment response DTOs
│   │   ├── order/
│   │   │   ├── request.go           # Order request DTOs
│   │   │   └── response.go          # Order response DTOs
│   │   └── common/
│   │       ├── pagination.go        # Pagination DTOs
│   │       ├── filter.go            # Filter DTOs
│   │       └── response.go          # Common response DTOs
│   │
│   ├── route/                       # Route management with versioning
│   │   ├── routes.go                # Main route setup
│   │   ├── v1/                      # API v1 routes
│   │   │   ├── auth_routes.go
│   │   │   ├── user_routes.go
│   │   │   ├── payment_routes.go
│   │   │   ├── order_routes.go
│   │   │   └── health_routes.go
│   │   └── v2/                      # API v2 routes (future)
│   │       ├── auth_routes.go
│   │       ├── user_routes.go
│   │       ├── payment_routes.go
│   │       ├── order_routes.go
│   │       └── health_routes.go
│   │
│   ├── middleware/                  # Fiber middleware
│   │   ├── auth.go
│   │   ├── cors.go
│   │   ├── csrf.go                  # CSRF protection
│   │   ├── helmet.go                # Security headers
│   │   ├── logger.go
│   │   └── rate_limiter.go
│   │
│   ├── repository/                  # Repository implementations (infrastructure)
│   │   ├── user_repository.go
│   │   ├── auth_repository.go
│   │   ├── payment_repository.go
│   │   └── order_repository.go
│   │
│   ├── model/                       # GORM models with conversion methods
│   │   ├── user.go
│   │   ├── auth.go
│   │   ├── payment.go
│   │   └── order.go
│   │
│   └── pkg/                         # Shared packages
│       ├── utils/
│       │   ├── app.go               # App initialization
│       │   └── password.go          # Password utilities
│       ├── jwt/
│       │   └── jwt.go               # JWT utilities
│       ├── response/
│       │   └── response.go          # HTTP response helpers
│       └── validator/
│           └── validator.go         # Validation utilities
│
├── migrations/                      # Database migrations
│   ├── 00001_create_users_table.up.sql
│   ├── 00001_create_users_table.down.sql
│   ├── 00002_create_auth_tables.up.sql
│   ├── 00002_create_auth_tables.down.sql
│   └── ...
│
├── docs/                           # Documentation
│   ├── MIGRATION.md                # Migration management guide
│   └── API.md                      # API documentation
│
├── scripts/                        # Utility scripts
│   └── setup_db.sh                 # Database setup script
│
├── tests/                          # Test files
│   ├── unit/
│   ├── integration/
│   └── e2e/
│
├── .env                            # Environment variables
├── .env.example                    # Environment template
├── .gitignore                      # Git ignore rules
├── go.mod                          # Go modules
├── go.sum                          # Go modules checksum
├── Makefile                        # Build and migration commands
├── Dockerfile                      # Docker configuration
├── docker-compose.yml              # Docker Compose setup
└── README.md                       # This file
```

## 🗄️ Migration Management

We provide comprehensive migration management tools:

### Quick Start

```bash
# Show all available commands
make help

# Run migrations
make migrate-up

# Check migration status
make migrate-status

# Create new migration
make migrate-create NAME=add_new_table

# Wipe all data and recreate schema (DANGEROUS!)
make migrate-wipe
```

### CLI Tool

```bash
# Show help
go run cmd/migrate/main.go

# Run migrations
go run cmd/migrate/main.go -action=up

# Check status
go run cmd/migrate/main.go -action=status

# Wipe data (with confirmation)
go run cmd/migrate/main.go -action=wipe -confirm
```

### Wipe Data Functionality

⚠️ **DANGER ZONE** - Use with extreme caution!

```bash
# Interactive confirmation
make migrate-wipe

# Direct wipe with confirmation flag
go run cmd/migrate/main.go -action=wipe -confirm
```

**What wipe does:**

-   🗑️ Drops all tables and sequences
-   🔄 Removes migration version tracking
-   📈 Runs all migrations to recreate schema
-   ✨ Fresh database ready for use

**Safety features:**

-   ✅ Confirmation required for CLI tool
-   ✅ Interactive prompt for Makefile
-   ✅ Clear warnings about data loss
-   ✅ Database name display before action

### Documentation

For detailed migration management, see [docs/MIGRATION.md](docs/MIGRATION.md).

## 🚀 Quick Start

### Prerequisites

-   Go 1.21+
-   PostgreSQL 14+
-   Redis 6+
-   Docker & Docker Compose (optional)

### Tech Stack

-   **Framework**: Go Fiber v2
-   **ORM**: GORM with PostgreSQL
-   **Migration**: golang-migrate/migrate
-   **Cache**: Redis (single instance)
-   **Authentication**: JWT v5
-   **Security**: CSRF protection, Helmet middleware
-   **Validation**: go-playground/validator/v10
-   **Configuration**: Viper
-   **Logging**: Logrus (structured logging)
-   **Testing**: Testify
-   **Rate Limiting**: Fiber built-in limiter

### Installation

1. **Clone the repository**

```bash
git clone https://github.com/your-username/boilerplate-go-fiber-v2.git
cd boilerplate-go-fiber-v2
```

2. **Install dependencies**

```bash
go mod download
```

3. **Setup environment**

```bash
cp .env.example .env
# Edit .env with your configuration
```

4. **Run database migrations**

```bash
make migrate
```

5. **Start the application**

```bash
make run
```

### Docker Setup

```bash
# Start all services
docker-compose up -d

# Run migrations
docker-compose exec app make migrate

# View logs
docker-compose logs -f app
```

## 📚 API Documentation

### API Versioning

This project supports API versioning with `/api/v1`, `/api/v2`, etc.

### Authentication Endpoints (v1)

```http
POST /api/v1/auth/register
POST /api/v1/auth/login
POST /api/v1/auth/logout
POST /api/v1/auth/refresh
POST /api/v1/auth/forgot-password
POST /api/v1/auth/reset-password
POST /api/v1/auth/change-password
```

### User Endpoints (v1)

```http
GET  /api/v1/users/profile
PUT  /api/v1/users/profile
PUT  /api/v1/users/avatar
DELETE /api/v1/users/account
GET  /api/v1/users/admin
GET  /api/v1/users/admin/:id
PUT  /api/v1/users/admin/:id/status
```

### Payment Endpoints (v1)

```http
POST /api/v1/payments
GET  /api/v1/payments
GET  /api/v1/payments/:id
POST /api/v1/payments/webhook
```

### Order Endpoints (v1)

```http
POST /api/v1/orders
GET  /api/v1/orders
GET  /api/v1/orders/:id
PUT  /api/v1/orders/:id/status
```

### Health Check Endpoints

```http
GET /health
GET /api/v1/health
GET /api/v1/health/db
GET /api/v1/health/redis
```

### Future v2 Endpoints

```http
# Enhanced Authentication (v2)
POST /api/v2/auth/register
POST /api/v2/auth/login
POST /api/v2/auth/login/2fa
POST /api/v2/auth/enable-2fa
POST /api/v2/auth/disable-2fa

# Enhanced User Management (v2)
GET  /api/v2/users/profile
PUT  /api/v2/users/profile
# ... additional v2 features
```

## 🧪 Testing

### Run all tests

```bash
make test
```

### Run specific test

```bash
go test ./internal/domain/service/...
```

### Run integration tests

```bash
make test-integration
```

### Run with coverage

```bash
make test-coverage
```

## 🛠️ Development

### Available Make Commands

```bash
make run              # Run the application
make build            # Build the application
make test             # Run tests
make test-coverage    # Run tests with coverage
make migrate          # Run database migrations
make migrate-down     # Rollback migrations
make seed             # Seed database with sample data
make lint             # Run linter
make format           # Format code
make clean            # Clean build artifacts
```

### Code Structure Guidelines

1. **Domain Layer**: Pure business logic, no framework dependencies
2. **Infrastructure Layer**: Framework-specific implementations
3. **DTOs**: Separate request/response structures from entities
4. **Testing**: Unit tests for business logic, integration tests for APIs

### Adding New Features

1. **Define Entity** in `internal/domain/entity/`
2. **Create Repository Interface** in `internal/domain/repository/`
3. **Implement Repository** in `internal/repository/`
4. **Create Service** in `internal/domain/service/`
5. **Define DTOs** in `internal/dto/`
6. **Create Handler** in `internal/handler/`
7. **Create Route** in `internal/route/v1/` (or v2 for new version)
8. **Register Route** in `internal/route/routes.go`

## 🔧 Configuration

### Environment Variables

```env
# Server Configuration
PORT=8080
HOST=localhost
ENV=development

# Database Configuration
DB_HOST=localhost
DB_PORT=5432
DB_USER=postgres
DB_PASSWORD=password
DB_NAME=boilerplate
DB_SSL_MODE=disable

# Redis Configuration
REDIS_HOST=localhost
REDIS_PORT=6379
REDIS_PASSWORD=
REDIS_DB=0

# JWT Configuration
JWT_SECRET=your-super-secret-jwt-key
JWT_EXPIRY=24h

# Logging Configuration
LOG_LEVEL=info

# Security Configuration
CSRF_SECRET=your-csrf-secret-key

# Payment Gateway Configuration
XENDIT_API_KEY=your-xendit-api-key
XENDIT_BASE_URL=https://api.xendit.co
MIDTRANS_SERVER_KEY=your-midtrans-server-key
MIDTRANS_CLIENT_KEY=your-midtrans-client-key
MIDTRANS_BASE_URL=https://api.midtrans.com
```

## 🚀 Deployment

### Docker Deployment

```bash
# Build production image
docker build -f Dockerfile.prod -t boilerplate:latest .

# Run with docker-compose
docker-compose -f docker-compose.prod.yml up -d
```

### Kubernetes Deployment

```bash
# Apply Kubernetes manifests
kubectl apply -f k8s/
```

## 📊 Monitoring

### Health Check

```http
GET /health
```

### Metrics

```http
GET /metrics
```

## 🤝 Contributing

1. Fork the repository
2. Create your feature branch (`git checkout -b feature/amazing-feature`)
3. Commit your changes (`git commit -m 'Add some amazing feature'`)
4. Push to the branch (`git push origin feature/amazing-feature`)
5. Open a Pull Request

### Development Guidelines

-   Follow Go coding standards
-   Write tests for new features
-   Update documentation
-   Use conventional commits

## 📄 License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

## 🙏 Acknowledgments

-   [Fiber](https://gofiber.io/) - Fast HTTP framework
-   [GORM](https://gorm.io/) - ORM library
-   [Logrus](https://github.com/sirupsen/logrus) - Structured logging
-   [Viper](https://github.com/spf13/viper) - Configuration management
-   [Testify](https://github.com/stretchr/testify) - Testing framework
-   [Clean Architecture](https://blog.cleancoder.com/uncle-bob/2012/08/13/the-clean-architecture.html) - Architecture pattern

## 📞 Support

If you have any questions or need help, please open an issue or contact us at support@example.com.

---

**Made with ❤️ by [Your Name]**
