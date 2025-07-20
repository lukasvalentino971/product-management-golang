package config

import "os"

type Config struct {
    Port        string
    DatabaseURL string
    JWTSecret   string
}

func Load() *Config {
    return &Config{
        Port:        getEnv("PORT", "8080"),
        DatabaseURL: getEnv("DATABASE_URL", "root:root@tcp(localhost:3306)/product_testing?charset=utf8mb4&parseTime=True&loc=Local"),
        JWTSecret:   getEnv("JWT_SECRET", "your-secret-key"),
    }
}

func getEnv(key, defaultValue string) string {
    if value := os.Getenv(key); value != "" {
        return value
    }
    return defaultValue
}
