# Implementation Plan - Boilerplate Go Fiber v2

## 🎯 Project Overview

This document outlines the step-by-step implementation plan for the Go Fiber boilerplate with Clean Architecture, authentication, payment integration, and comprehensive testing.

## 📋 Implementation Phases

### Phase 1: Project Setup & Core Infrastructure

**Duration**: 2-3 days
**Priority**: High

#### 1.1 Project Initialization

-   [ ] Initialize Go module
-   [ ] Setup project structure
-   [ ] Create basic configuration files
-   [ ] Setup environment variables

#### 1.2 Core Dependencies

-   [ ] Install and configure Go Fiber v2
-   [ ] Setup GORM with PostgreSQL
-   [ ] Configure Redis client
-   [ ] Setup golang-migrate/migrate
-   [ ] Install validation libraries (go-playground/validator/v10)
-   [ ] Setup JWT library (golang-jwt/jwt/v5)
-   [ ] Configure Logrus for structured logging
-   [ ] Setup Viper for configuration management

#### 1.3 Database Setup

-   [ ] Create PostgreSQL database
-   [ ] Setup database connection with connection pooling
-   [ ] Create initial migration files
-   [ ] Setup database seeding

#### 1.4 Basic Configuration

-   [ ] Implement config package with Viper
-   [ ] Setup environment variable loading
-   [ ] Configure database connection
-   [ ] Setup Redis connection
-   [ ] Create basic health check endpoints

### Phase 2: Domain Layer Implementation

**Duration**: 3-4 days
**Priority**: High

#### 2.1 Entity Definitions

-   [ ] Create User entity (with TFA fields)
-   [ ] Create Auth entity (sessions, password resets, TFA codes)
-   [ ] Create Payment entity
-   [ ] Create Order entity
-   [ ] Add business methods to entities
-   [ ] Add TFA-related methods to User entity

#### 2.2 Repository Interfaces

-   [ ] Define UserRepository interface
-   [ ] Define AuthRepository interface
-   [ ] Define PaymentRepository interface
-   [ ] Define OrderRepository interface
-   [ ] Add method contracts for all CRUD operations

#### 2.3 Service Layer

-   [ ] Implement AuthService with business logic
-   [ ] Implement UserService with business logic
-   [ ] Implement PaymentService with business logic
-   [ ] Implement OrderService with business logic
-   [ ] Add validation and error handling
-   [ ] Implement TFA service (generate, validate codes)
-   [ ] Implement email service for password reset
-   [ ] Implement rate limiting service for auth

### Phase 3: Infrastructure Layer

**Duration**: 2-3 days
**Priority**: High

#### 3.1 Repository Implementations

-   [ ] Implement UserRepository with GORM
-   [ ] Implement AuthRepository with GORM
-   [ ] Implement PaymentRepository with GORM
-   [ ] Implement OrderRepository with GORM
-   [ ] Add proper error handling and logging

#### 3.2 External Services

-   [ ] Create PaymentGateway interface
-   [ ] Implement XenditAdapter
-   [ ] Implement MidtransAdapter
-   [ ] Add webhook handling
-   [ ] Implement payment status checking
-   [ ] Implement EmailService interface
-   [ ] Implement SMTP email adapter
-   [ ] Implement SendGrid email adapter
-   [ ] Add TFA code generation service

#### 3.3 Database Migrations

-   [ ] Create users table migration (with TFA fields)
-   [ ] Create auth_sessions table migration
-   [ ] Create password_resets table migration
-   [ ] Create tfa_codes table migration
-   [ ] Create orders table migration
-   [ ] Create payments table migration
-   [ ] Add indexes and constraints

### Phase 4: Application Layer

**Duration**: 3-4 days
**Priority**: High

#### 4.1 DTOs (Data Transfer Objects)

-   [ ] Create auth DTOs (request/response)
-   [ ] Create user DTOs (request/response)
-   [ ] Create payment DTOs (request/response)
-   [ ] Create order DTOs (request/response)
-   [ ] Add validation tags
-   [ ] Create common DTOs (pagination, filters)
-   [ ] Create TFA DTOs (enable, verify, disable)
-   [ ] Create password reset DTOs
-   [ ] Create email verification DTOs

#### 4.2 HTTP Handlers

-   [ ] Implement AuthHandler
-   [ ] Implement UserHandler
-   [ ] Implement PaymentHandler
-   [ ] Implement OrderHandler
-   [ ] Implement HealthHandler
-   [ ] Add proper error responses
-   [ ] Add TFA endpoints (enable, verify, disable)
-   [ ] Add password reset endpoints
-   [ ] Add email verification endpoints

#### 4.3 Route Management

-   [ ] Setup main route configuration
-   [ ] Create v1 route structure
-   [ ] Implement auth routes
-   [ ] Implement user routes
-   [ ] Implement payment routes
-   [ ] Implement order routes
-   [ ] Setup v2 route structure (future)

### Phase 5: Security & Middleware

**Duration**: 2-3 days
**Priority**: High

#### 5.1 Authentication Middleware

-   [ ] Implement JWT token validation
-   [ ] Add user context injection
-   [ ] Implement role-based authorization
-   [ ] Add token refresh logic
-   [ ] Implement rate limiting for auth endpoints
-   [ ] Add TFA (Two-Factor Authentication) support
-   [ ] Implement password reset functionality
-   [ ] Add email verification system

#### 5.2 Security Middleware

-   [ ] Implement CSRF protection
-   [ ] Add Helmet security headers
-   [ ] Setup CORS configuration
-   [ ] Implement rate limiting
-   [ ] Add request logging

#### 5.3 Validation

-   [ ] Setup custom validators
-   [ ] Implement request validation
-   [ ] Add error handling for validation
-   [ ] Create validation helpers

### Phase 6: Testing Implementation

**Duration**: 3-4 days
**Priority**: Medium

#### 6.1 Unit Tests

-   [ ] Test domain services
-   [ ] Test repository implementations
-   [ ] Test external adapters
-   [ ] Test utility functions
-   [ ] Add test coverage reporting

#### 6.2 Integration Tests

-   [ ] Test HTTP handlers
-   [ ] Test authentication flow
-   [ ] Test payment integration
-   [ ] Test database operations
-   [ ] Test external API calls

#### 6.3 Test Utilities

-   [ ] Create test database setup
-   [ ] Implement test fixtures
-   [ ] Add mock generators
-   [ ] Setup test helpers

### Phase 7: Documentation & DevOps

**Duration**: 2-3 days
**Priority**: Medium

#### 7.1 API Documentation

-   [ ] Create OpenAPI/Swagger documentation
-   [ ] Document all endpoints
-   [ ] Add request/response examples
-   [ ] Create Postman collection

#### 7.2 Docker Setup

-   [ ] Create Dockerfile
-   [ ] Create docker-compose.yml
-   [ ] Setup multi-stage builds
-   [ ] Configure production Dockerfile

#### 7.3 CI/CD Pipeline

-   [ ] Setup GitHub Actions
-   [ ] Add automated testing
-   [ ] Add code quality checks
-   [ ] Setup deployment pipeline

### Phase 8: Advanced Features

**Duration**: 2-3 days
**Priority**: Low

#### 8.1 Monitoring & Logging

-   [ ] Implement structured logging
-   [ ] Add request tracing
-   [ ] Setup error monitoring
-   [ ] Add performance metrics

#### 8.2 Caching Strategy

-   [ ] Implement Redis caching
-   [ ] Add cache invalidation
-   [ ] Setup cache middleware
-   [ ] Add cache monitoring

#### 8.3 Performance Optimization

-   [ ] Add database query optimization
-   [ ] Implement connection pooling
-   [ ] Add response compression
-   [ ] Setup load balancing

## 🛠️ Technical Implementation Details

### Database Schema

#### Users Table

```sql
CREATE TABLE users (
    id BIGSERIAL PRIMARY KEY,
    email VARCHAR(255) UNIQUE NOT NULL,
    username VARCHAR(50) UNIQUE NOT NULL,
    password_hash VARCHAR(255) NOT NULL,
    first_name VARCHAR(100) NOT NULL,
    last_name VARCHAR(100) NOT NULL,
    phone VARCHAR(20),
    avatar VARCHAR(500),
    role VARCHAR(20) DEFAULT 'user',
    status VARCHAR(20) DEFAULT 'active',
    email_verified_at TIMESTAMP,
    phone_verified_at TIMESTAMP,
    last_login_at TIMESTAMP,
    tfa_enabled BOOLEAN DEFAULT FALSE,
    tfa_secret VARCHAR(255),
    tfa_backup_codes TEXT[],
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP
);
```

#### Auth Sessions Table

```sql
CREATE TABLE auth_sessions (
    id BIGSERIAL PRIMARY KEY,
    user_id BIGINT NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    token VARCHAR(500) UNIQUE NOT NULL,
    refresh_token VARCHAR(500) UNIQUE NOT NULL,
    expires_at TIMESTAMP NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);
```

#### Password Resets Table

```sql
CREATE TABLE password_resets (
    id BIGSERIAL PRIMARY KEY,
    user_id BIGINT NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    token VARCHAR(255) UNIQUE NOT NULL,
    expires_at TIMESTAMP NOT NULL,
    used BOOLEAN DEFAULT FALSE,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);
```

#### TFA Codes Table

```sql
CREATE TABLE tfa_codes (
    id BIGSERIAL PRIMARY KEY,
    user_id BIGINT NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    code VARCHAR(10) NOT NULL,
    expires_at TIMESTAMP NOT NULL,
    used BOOLEAN DEFAULT FALSE,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);
```

### API Endpoints Structure

#### Authentication (v1)

```http
POST /api/v1/auth/register
POST /api/v1/auth/login
POST /api/v1/auth/logout
POST /api/v1/auth/refresh
POST /api/v1/auth/forgot-password
POST /api/v1/auth/reset-password
POST /api/v1/auth/change-password
POST /api/v1/auth/verify-email
POST /api/v1/auth/resend-verification
POST /api/v1/auth/enable-2fa
POST /api/v1/auth/disable-2fa
POST /api/v1/auth/verify-2fa
POST /api/v1/auth/setup-2fa
```

#### Users (v1)

```http
GET  /api/v1/users/profile
PUT  /api/v1/users/profile
PUT  /api/v1/users/avatar
DELETE /api/v1/users/account
GET  /api/v1/users/admin
GET  /api/v1/users/admin/:id
PUT  /api/v1/users/admin/:id/status
```

#### Payments (v1)

```http
POST /api/v1/payments
GET  /api/v1/payments
GET  /api/v1/payments/:id
POST /api/v1/payments/webhook
```

#### Orders (v1)

```http
POST /api/v1/orders
GET  /api/v1/orders
GET  /api/v1/orders/:id
PUT  /api/v1/orders/:id/status
```

### Environment Configuration

#### Required Environment Variables

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

# Email Configuration
SMTP_HOST=smtp.gmail.com
SMTP_PORT=587
SMTP_USERNAME=your-email@gmail.com
SMTP_PASSWORD=your-app-password
SENDGRID_API_KEY=your-sendgrid-api-key

# TFA Configuration
TFA_ISSUER=YourApp
TFA_ALGORITHM=SHA1
TFA_DIGITS=6
TFA_PERIOD=30

# Payment Gateway Configuration
XENDIT_API_KEY=your-xendit-api-key
XENDIT_BASE_URL=https://api.xendit.co
MIDTRANS_SERVER_KEY=your-midtrans-server-key
MIDTRANS_CLIENT_KEY=your-midtrans-client-key
MIDTRANS_BASE_URL=https://api.midtrans.com
```

## 🚀 Development Workflow

### Daily Development Process

1. **Morning Setup**

    - Pull latest changes
    - Run tests
    - Check current phase progress

2. **Development Tasks**

    - Implement features according to phase
    - Write tests for new features
    - Update documentation

3. **End of Day**
    - Commit changes
    - Push to repository
    - Update progress tracking

### Code Quality Standards

-   [ ] Follow Go coding standards
-   [ ] Write comprehensive tests
-   [ ] Add proper error handling
-   [ ] Include logging for debugging
-   [ ] Document complex functions
-   [ ] Use meaningful variable names

### Testing Strategy

-   [ ] Unit tests for business logic
-   [ ] Integration tests for APIs
-   [ ] End-to-end tests for critical flows
-   [ ] Performance tests for bottlenecks
-   [ ] Security tests for vulnerabilities

## 📊 Success Metrics

### Phase Completion Criteria

-   [ ] All features implemented
-   [ ] Tests passing with >80% coverage
-   [ ] Documentation complete
-   [ ] Security audit passed
-   [ ] Performance benchmarks met

### Quality Gates

-   [ ] Code review completed
-   [ ] Tests passing
-   [ ] No security vulnerabilities
-   [ ] Performance requirements met
-   [ ] Documentation updated

## 🎯 Timeline Summary

| Phase   | Duration | Priority | Dependencies |
| ------- | -------- | -------- | ------------ |
| Phase 1 | 2-3 days | High     | None         |
| Phase 2 | 3-4 days | High     | Phase 1      |
| Phase 3 | 2-3 days | High     | Phase 2      |
| Phase 4 | 3-4 days | High     | Phase 3      |
| Phase 5 | 2-3 days | High     | Phase 4      |
| Phase 6 | 3-4 days | Medium   | Phase 5      |
| Phase 7 | 2-3 days | Medium   | Phase 6      |
| Phase 8 | 2-3 days | Low      | Phase 7      |

**Total Estimated Duration**: 19-27 days

## 🔄 Risk Management

### Potential Risks

1. **Technical Risks**

    - Payment gateway integration complexity
    - Database performance issues
    - Security vulnerabilities

2. **Timeline Risks**
    - Scope creep
    - Resource constraints
    - External dependencies

### Mitigation Strategies

1. **Technical Mitigation**

    - Early prototyping
    - Comprehensive testing
    - Security reviews

2. **Timeline Mitigation**
    - Clear scope definition
    - Regular progress reviews
    - Buffer time allocation

## 📝 Notes

-   This plan is flexible and can be adjusted based on requirements
-   Each phase should be completed before moving to the next
-   Regular reviews and adjustments are recommended
-   Focus on quality over speed
-   Maintain documentation throughout development

---

**Last Updated**: [Current Date]
**Version**: 1.0
**Status**: Planning Phase
