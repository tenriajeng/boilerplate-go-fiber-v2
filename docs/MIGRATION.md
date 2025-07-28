# 🗄️ Database Migration Management

This document explains how to manage database migrations in the Boilerplate Go Fiber v2 project.

## 📋 Overview

We use **golang-migrate/migrate** for database migrations with PostgreSQL. The migration system provides:

-   ✅ Version control for database schema
-   ✅ Up/down migrations
-   ✅ CLI tool integration
-   ✅ Makefile commands
-   ✅ Automatic configuration loading
-   ✅ **Data wipe functionality for development/testing**

## 🛠️ Tools Available

### 1. CLI Migration Tool

```bash
# Show help
go run cmd/migrate/main.go

# Run all pending migrations
go run cmd/migrate/main.go -action=up

# Rollback all migrations
go run cmd/migrate/main.go -action=down

# Show current version
go run cmd/migrate/main.go -action=status

# Create new migration
go run cmd/migrate/main.go -action=create -name=add_users

# Force to specific version
go run cmd/migrate/main.go -action=force -version=1

# Wipe all data and recreate schema (DANGEROUS!)
go run cmd/migrate/main.go -action=wipe -confirm
```

### 2. Makefile Commands

```bash
# Show all available commands
make help

# Migration commands
make migrate-up
make migrate-down
make migrate-status
make migrate-create NAME=migration_name
make migrate-force VERSION=1
make migrate-wipe
```

## 🗑️ Wipe Data Functionality

### ⚠️ **DANGER ZONE** - Wipe All Data

The wipe functionality will **DELETE ALL DATA** in the database and recreate the schema from scratch. Use with extreme caution!

### When to Use Wipe Data:

-   🧪 **Development**: Fresh start during development
-   🧪 **Testing**: Clean database for integration tests
-   🔄 **Reset**: Complete database reset
-   🐛 **Debugging**: Eliminate data-related issues

### How to Use:

#### Method 1: CLI Tool (with confirmation)

```bash
# First run (shows warning)
go run cmd/migrate/main.go -action=wipe

# With confirmation flag
go run cmd/migrate/main.go -action=wipe -confirm
```

#### Method 2: Makefile (interactive confirmation)

```bash
make migrate-wipe
# Will prompt: "Are you sure you want to continue? (yes/no):"
```

### What Wipe Does:

1. **Drops all tables** in the database
2. **Drops all sequences** (auto-increment)
3. **Removes migration version** tracking
4. **Runs all migrations** to recreate schema
5. **Fresh database** ready for use

### Safety Features:

-   ✅ **Confirmation required** for CLI tool
-   ✅ **Interactive prompt** for Makefile
-   ✅ **Clear warnings** about data loss
-   ✅ **Database name display** before action
-   ✅ **Automatic schema recreation** after wipe

## 📁 Migration Files Structure

```
migrations/
├── 00001_create_users_table.up.sql
├── 00001_create_users_table.down.sql
├── 00002_create_auth_tables.up.sql
├── 00002_create_auth_tables.down.sql
├── 00003_add_payments.up.sql
├── 00003_add_payments.down.sql
└── 00004_add_orders.up.sql
└── 00004_add_orders.down.sql
```

**Format Penamaan:**

-   ✅ **Sequential numbering**: `00001`, `00002`, `00003`, etc.
-   ✅ **5-digit format**: Memastikan urutan yang konsisten
-   ✅ **Descriptive names**: `create_users_table`, `add_payments`, etc.
-   ✅ **Up/Down pairs**: Setiap migration memiliki file `.up.sql` dan `.down.sql`

## 🔧 Configuration

The migration tool automatically loads configuration from `.env` file:

```env
DB_HOST=localhost
DB_PORT=5432
DB_USER=postgres
DB_PASSWORD=ilham
DB_NAME=boilerplate_go_fiber_v2
DB_SSL_MODE=disable
```

## 📝 Creating Migrations

### Method 1: Using Makefile (Recommended)

```bash
make migrate-create NAME=add_new_table
```

### Method 2: Using CLI Tool

```bash
go run cmd/migrate/main.go -action=create -name=add_new_table
```

### Method 3: Direct migrate command

```bash
migrate create -ext sql -dir migrations add_new_table
```

## 🚀 Running Migrations

### Run All Pending Migrations

```bash
make migrate-up
# or
go run cmd/migrate/main.go -action=up
```

### Run Specific Number of Migrations

```bash
go run cmd/migrate/main.go -action=up -steps=1
```

### Rollback All Migrations

```bash
make migrate-down
# or
go run cmd/migrate/main.go -action=down
```

### Rollback Specific Number of Migrations

```bash
go run cmd/migrate/main.go -action=down -steps=1
```

## 📊 Checking Migration Status

### Show Current Version

```bash
make migrate-status
# or
go run cmd/migrate/main.go -action=status
```

### Show Migration Info

```bash
migrate -path migrations -database "postgres://user:pass@localhost:5432/dbname?sslmode=disable" version
```

## 🔧 Troubleshooting

### Force Migration to Specific Version

If migrations get into a dirty state:

```bash
make migrate-force VERSION=2
# or
go run cmd/migrate/main.go -action=force -version=2
```

### Reset Database

If you need to start fresh:

```bash
# Option 1: Use wipe functionality
make migrate-wipe

# Option 2: Manual reset
dropdb boilerplate_go_fiber_v2
createdb boilerplate_go_fiber_v2
make migrate-up
```

### Fix Dirty State

If migration is marked as dirty:

```bash
# Force to current version
make migrate-force VERSION=2

# Then run migrations
make migrate-up
```

## 📋 Migration Best Practices

### 1. Naming Conventions

-   Use descriptive names: `add_users`, `create_orders`, `add_payment_gateway`
-   Use snake_case for table and column names
-   Include table name in migration name

### 2. Migration Structure

```sql
-- Up migration
CREATE TABLE users (
    id SERIAL PRIMARY KEY,
    email VARCHAR(255) UNIQUE NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- Down migration
DROP TABLE users;
```

### 3. Data Migrations

For data migrations, use separate migrations:

```sql
-- 001_add_users.up.sql
CREATE TABLE users (id SERIAL PRIMARY KEY, email VARCHAR(255));

-- 002_populate_users.up.sql
INSERT INTO users (email) VALUES ('admin@example.com');
```

### 4. Rollback Safety

Always ensure down migrations are safe:

```sql
-- Safe down migration
DROP TABLE IF EXISTS users;
```

## 🔍 Debugging

### Check Configuration

```bash
go run cmd/migrate/main.go
# Shows current configuration
```

### Check Database Connection

```bash
psql -h localhost -U postgres -d boilerplate_go_fiber_v2
```

### View Migration Files

```bash
ls -la migrations/
```

## 📚 Additional Resources

-   [golang-migrate Documentation](https://github.com/golang-migrate/migrate)
-   [PostgreSQL Documentation](https://www.postgresql.org/docs/)
-   [GORM Documentation](https://gorm.io/docs/)

## 🆘 Common Issues

### 1. Password Authentication Failed

**Problem**: `pq: password authentication failed for user "postgres"`

**Solution**: Check `.env` file and ensure correct password:

```env
DB_PASSWORD=your_actual_password
```

### 2. Database Does Not Exist

**Problem**: `pq: database "boilerplate_go_fiber_v2" does not exist`

**Solution**: Create database:

```bash
createdb boilerplate_go_fiber_v2
```

### 3. Migration Tool Not Found

**Problem**: `migrate: command not found`

**Solution**: Install migrate tool:

```bash
go install -tags 'postgres' github.com/golang-migrate/migrate/v4/cmd/migrate@latest
```

### 4. Dirty Database State

**Problem**: `Dirty database version X. Fix and force version.`

**Solution**: Force to current version:

```bash
make migrate-force VERSION=X
```

### 5. Wipe Data Permission Denied

**Problem**: `permission denied` when running wipe

**Solution**: Ensure PostgreSQL user has proper permissions:

```sql
-- Grant all privileges to user
GRANT ALL PRIVILEGES ON DATABASE boilerplate_go_fiber_v2 TO postgres;
```

## 🚨 Wipe Data Safety Checklist

Before using wipe functionality:

-   ✅ **Backup important data** (if any)
-   ✅ **Confirm database name** is correct
-   ✅ **Verify environment** (dev/test only)
-   ✅ **Check team members** are aware
-   ✅ **Have migration files** ready
-   ✅ **Test on non-production** first
