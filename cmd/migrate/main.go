package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"strconv"
	"strings"

	"boilerplate-go-fiber-v2/config"
)

func main() {
	// Parse command line flags
	var (
		action  = flag.String("action", "", "Migration action: up, down, create, version, force, wipe")
		steps   = flag.Int("steps", 0, "Number of steps for up/down (0 = all)")
		version = flag.Int("version", 0, "Version for force command")
		name    = flag.String("name", "", "Name for create command")
		confirm = flag.Bool("confirm", false, "Confirm wipe data (required for wipe action)")
	)
	flag.Parse()

	// Load configuration
	cfg := config.Load()

	// Debug: Print configuration
	fmt.Printf("🔧 Configuration:\n")
	fmt.Printf("  DB_HOST: %s\n", cfg.Database.Host)
	fmt.Printf("  DB_PORT: %s\n", cfg.Database.Port)
	fmt.Printf("  DB_USER: %s\n", cfg.Database.User)
	fmt.Printf("  DB_PASSWORD: %s\n", cfg.Database.Password)
	fmt.Printf("  DB_NAME: %s\n", cfg.Database.Name)
	fmt.Printf("  DB_SSL_MODE: %s\n", cfg.Database.SSLMode)

	// Build database URL
	dbURL := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=%s",
		cfg.Database.User,
		cfg.Database.Password,
		cfg.Database.Host,
		cfg.Database.Port,
		cfg.Database.Name,
		cfg.Database.SSLMode,
	)

	// Get migrations directory
	migrationsDir, err := filepath.Abs("migrations")
	if err != nil {
		log.Fatal("Failed to get migrations directory:", err)
	}

	// Execute action
	switch *action {
	case "up":
		args := []string{"-path", migrationsDir, "-database", dbURL, "up"}
		if *steps > 0 {
			args = append(args, "-step", fmt.Sprintf("%d", *steps))
		}
		runMigrateCommand(args...)

	case "down":
		args := []string{"-path", migrationsDir, "-database", dbURL, "down"}
		if *steps > 0 {
			args = append(args, "-step", fmt.Sprintf("%d", *steps))
		}
		runMigrateCommand(args...)

	case "version":
		runMigrateCommand("-path", migrationsDir, "-database", dbURL, "version")

	case "force":
		if *version < 0 {
			log.Fatal("Version must be >= 0")
		}
		runMigrateCommand("-path", migrationsDir, "-database", dbURL, "force", fmt.Sprintf("%d", *version))

	case "create":
		if *name == "" {
			log.Fatal("Name is required for create command")
		}
		createSequentialMigration(migrationsDir, *name)

	case "status":
		runMigrateCommand("-path", migrationsDir, "-database", dbURL, "version")

	case "wipe":
		if !*confirm {
			fmt.Println("⚠️  WARNING: This will DELETE ALL DATA in the database!")
			fmt.Printf("📊 Database: %s\n", cfg.Database.Name)
			fmt.Printf("🔗 Database URL: %s\n", dbURL)
			fmt.Println("")
			fmt.Println("To confirm, run with -confirm flag:")
			fmt.Printf("  go run cmd/migrate/main.go -action=wipe -confirm\n")
			fmt.Println("")
			fmt.Println("Or use Makefile:")
			fmt.Printf("  make migrate-wipe\n")
			os.Exit(1)
		}

		fmt.Println("🗑️  Wiping all data from database...")
		wipeDatabase(cfg)
		fmt.Println("✅ Database wiped successfully!")
		fmt.Println("📈 Running migrations to recreate schema...")
		runMigrateCommand("-path", migrationsDir, "-database", dbURL, "up")

	default:
		fmt.Println("🚀 Migration Management Tool")
		fmt.Println("")
		fmt.Println("Usage:")
		fmt.Println("  go run cmd/migrate/main.go -action=up          # Run all pending migrations")
		fmt.Println("  go run cmd/migrate/main.go -action=down        # Rollback all migrations")
		fmt.Println("  go run cmd/migrate/main.go -action=up -steps=1 # Run 1 migration")
		fmt.Println("  go run cmd/migrate/main.go -action=down -steps=1 # Rollback 1 migration")
		fmt.Println("  go run cmd/migrate/main.go -action=version     # Show current version")
		fmt.Println("  go run cmd/migrate/main.go -action=force -version=1 # Force to version")
		fmt.Println("  go run cmd/migrate/main.go -action=create -name=add_users # Create new migration")
		fmt.Println("  go run cmd/migrate/main.go -action=status     # Show migration status")
		fmt.Println("  go run cmd/migrate/main.go -action=wipe -confirm # Wipe all data and recreate schema")
		fmt.Println("")
		fmt.Printf("📊 Database: %s\n", cfg.Database.Name)
		fmt.Printf("🔗 Database URL: %s\n", dbURL)
	}
}

func runMigrateCommand(args ...string) {
	cmd := exec.Command("migrate", args...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	fmt.Printf("🔄 Running: migrate %v\n", args)

	if err := cmd.Run(); err != nil {
		log.Fatal("Migration command failed:", err)
	}
}

func createSequentialMigration(migrationsDir, name string) {
	// Get next migration number
	nextNumber := getNextMigrationNumber(migrationsDir)

	// Create migration files with sequential numbering
	upFile := fmt.Sprintf("%s/%05d_%s.up.sql", migrationsDir, nextNumber, name)
	downFile := fmt.Sprintf("%s/%05d_%s.down.sql", migrationsDir, nextNumber, name)

	// Create up migration file
	upContent := fmt.Sprintf(`-- Migration %05d: %s
-- Up migration
-- TODO: Add your migration SQL here

`, nextNumber, name)

	// Create down migration file
	downContent := fmt.Sprintf(`-- Migration %05d: %s
-- Down migration
-- TODO: Add your rollback SQL here

`, nextNumber, name)

	// Write files
	if err := os.WriteFile(upFile, []byte(upContent), 0644); err != nil {
		log.Fatal("Failed to create up migration file:", err)
	}

	if err := os.WriteFile(downFile, []byte(downContent), 0644); err != nil {
		log.Fatal("Failed to create down migration file:", err)
	}

	fmt.Printf("✅ Created migration files:\n")
	fmt.Printf("  📄 %s\n", upFile)
	fmt.Printf("  📄 %s\n", downFile)
	fmt.Printf("  🔢 Migration number: %05d\n", nextNumber)
}

func getNextMigrationNumber(migrationsDir string) int {
	files, err := os.ReadDir(migrationsDir)
	if err != nil {
		log.Fatal("Failed to read migrations directory:", err)
	}

	maxNumber := 0

	for _, file := range files {
		if file.IsDir() {
			continue
		}

		filename := file.Name()
		if !strings.HasSuffix(filename, ".up.sql") && !strings.HasSuffix(filename, ".down.sql") {
			continue
		}

		// Extract number from filename
		parts := strings.Split(filename, "_")
		if len(parts) < 2 {
			continue
		}

		// Try to parse the number part
		numberStr := parts[0]
		if number, err := strconv.Atoi(numberStr); err == nil {
			if number > maxNumber {
				maxNumber = number
			}
		}
	}

	return maxNumber + 1
}

func wipeDatabase(cfg *config.Config) {
	// Connect to database and drop all tables
	dbURL := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=%s",
		cfg.Database.User,
		cfg.Database.Password,
		cfg.Database.Host,
		cfg.Database.Port,
		cfg.Database.Name,
		cfg.Database.SSLMode,
	)

	// SQL to drop all tables and reset migration version
	sql := `
		DO $$ 
		DECLARE 
			r RECORD;
		BEGIN
			-- Drop all tables
			FOR r IN (SELECT tablename FROM pg_tables WHERE schemaname = current_schema()) LOOP
				EXECUTE 'DROP TABLE IF EXISTS ' || quote_ident(r.tablename) || ' CASCADE';
			END LOOP;
			
			-- Drop all sequences
			FOR r IN (SELECT sequencename FROM pg_sequences WHERE schemaname = current_schema()) LOOP
				EXECUTE 'DROP SEQUENCE IF EXISTS ' || quote_ident(r.sequencename) || ' CASCADE';
			END LOOP;
			
			-- Drop schema_migrations table if exists
			DROP TABLE IF EXISTS schema_migrations;
		END $$;
	`

	// Execute SQL using psql
	cmd := exec.Command("psql", dbURL, "-c", sql)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	if err := cmd.Run(); err != nil {
		log.Fatal("Failed to wipe database:", err)
	}
}
