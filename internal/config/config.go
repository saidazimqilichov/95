package config

import "os"

type Config struct {
    Port        string
    DatabaseURL string
}

func New() *Config {
    return &Config{
        Port:        getEnv("PORT", "8080"),
        DatabaseURL: getEnv("DATABASE_URL", "postgres://postgres:7777@localhost/repo?sslmode=disable"),
    }
}

func getEnv(key, defaultValue string) string {
    if value, exists := os.LookupEnv(key); exists {
        return value
    }
    return defaultValue
}
