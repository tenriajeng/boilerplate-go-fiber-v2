#!/bin/bash

# Database setup script
echo "Setting up database..."

# Check if PostgreSQL is running
if ! pg_isready -h localhost -p 5432 > /dev/null 2>&1; then
    echo "PostgreSQL is not running. Please start PostgreSQL first."
    exit 1
fi

# Create database if it doesn't exist
echo "Creating database..."
createdb -h localhost -U postgres boilerplate_go_fiber_v2 2>/dev/null || echo "Database already exists"

# Run migrations
echo "Running migrations..."
migrate -path migrations -database "postgres://postgres:postgres@localhost:5432/boilerplate_go_fiber_v2?sslmode=disable" up

echo "Database setup complete!" 