package config

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/spf13/viper"
)

type Config struct {
	Server   ServerConfig
	Database DatabaseConfig
	Redis    RedisConfig
	JWT      JWTConfig
	Email    EmailConfig
	TFA      TFAConfig
	Payment  PaymentConfig
}

type ServerConfig struct {
	Port string
	Host string
	Env  string
}

type DatabaseConfig struct {
	Host     string
	Port     string
	User     string
	Password string
	Name     string
	SSLMode  string
}

type RedisConfig struct {
	Host     string
	Port     string
	Password string
	DB       int
}

type JWTConfig struct {
	Secret string
	Expiry time.Duration
}

type EmailConfig struct {
	SMTPHost     string
	SMTPPort     int
	SMTPUsername string
	SMTPPassword string
	SendGridKey  string
}

type TFAConfig struct {
	Issuer    string
	Algorithm string
	Digits    int
	Period    int
}

type PaymentConfig struct {
	XenditAPIKey      string
	XenditBaseURL     string
	MidtransServerKey string
	MidtransClientKey string
	MidtransBaseURL   string
}

var AppConfig *Config

func Load() *Config {
	// Load .env file
	viper.SetConfigName(".env")
	viper.SetConfigType("env")
	viper.AddConfigPath(".")
	viper.AutomaticEnv()

	// Read config file
	if err := viper.ReadInConfig(); err != nil {
		log.Printf("Warning: .env file not found, using environment variables")
	}

	config := &Config{
		Server: ServerConfig{
			Port: getViperEnv("PORT", "8080"),
			Host: getViperEnv("HOST", "localhost"),
			Env:  getViperEnv("ENV", "development"),
		},
		Database: DatabaseConfig{
			Host:     getViperEnv("DB_HOST", "localhost"),
			Port:     getViperEnv("DB_PORT", "5432"),
			User:     getViperEnv("DB_USER", "postgres"),
			Password: getViperEnv("DB_PASSWORD", "password"),
			Name:     getViperEnv("DB_NAME", "boilerplate"),
			SSLMode:  getViperEnv("DB_SSL_MODE", "disable"),
		},
		Redis: RedisConfig{
			Host:     getViperEnv("REDIS_HOST", "localhost"),
			Port:     getViperEnv("REDIS_PORT", "6379"),
			Password: getViperEnv("REDIS_PASSWORD", ""),
			DB:       getViperEnvAsInt("REDIS_DB", 0),
		},
		JWT: JWTConfig{
			Secret: getViperEnv("JWT_SECRET", "your-super-secret-jwt-key"),
			Expiry: getViperEnvAsDuration("JWT_EXPIRY", 24*time.Hour),
		},
		Email: EmailConfig{
			SMTPHost:     getViperEnv("SMTP_HOST", "smtp.gmail.com"),
			SMTPPort:     getViperEnvAsInt("SMTP_PORT", 587),
			SMTPUsername: getViperEnv("SMTP_USERNAME", ""),
			SMTPPassword: getViperEnv("SMTP_PASSWORD", ""),
			SendGridKey:  getViperEnv("SENDGRID_API_KEY", ""),
		},
		TFA: TFAConfig{
			Issuer:    getViperEnv("TFA_ISSUER", "YourApp"),
			Algorithm: getViperEnv("TFA_ALGORITHM", "SHA1"),
			Digits:    getViperEnvAsInt("TFA_DIGITS", 6),
			Period:    getViperEnvAsInt("TFA_PERIOD", 30),
		},
		Payment: PaymentConfig{
			XenditAPIKey:      getViperEnv("XENDIT_API_KEY", ""),
			XenditBaseURL:     getViperEnv("XENDIT_BASE_URL", "https://api.xendit.co"),
			MidtransServerKey: getViperEnv("MIDTRANS_SERVER_KEY", ""),
			MidtransClientKey: getViperEnv("MIDTRANS_CLIENT_KEY", ""),
			MidtransBaseURL:   getViperEnv("MIDTRANS_BASE_URL", "https://api.midtrans.com"),
		},
	}

	AppConfig = config
	return config
}

func getViperEnv(key, defaultValue string) string {
	if value := viper.GetString(key); value != "" {
		return value
	}
	return defaultValue
}

func getViperEnvAsInt(key string, defaultValue int) int {
	if value := viper.GetInt(key); value != 0 {
		return value
	}
	return defaultValue
}

func getViperEnvAsDuration(key string, defaultValue time.Duration) time.Duration {
	if value := viper.GetDuration(key); value != 0 {
		return value
	}
	return defaultValue
}

func getEnv(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}

func (c *Config) GetDatabaseDSN() string {
	return fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		c.Database.Host,
		c.Database.Port,
		c.Database.User,
		c.Database.Password,
		c.Database.Name,
		c.Database.SSLMode,
	)
}

func (c *Config) GetRedisAddr() string {
	return fmt.Sprintf("%s:%s", c.Redis.Host, c.Redis.Port)
}
