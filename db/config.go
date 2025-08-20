package db

import (
	"os"
)

// Config contém as configurações de conexão com o banco
type Config struct {
	Host     string
	Port     int
	User     string
	Password string
	DBName   string
	SSLMode  string
}

// NewConfig cria uma nova configuração com valores padrão ou de variáveis de ambiente
func NewConfig() *Config {
	return &Config{
		Host:     getEnv("DB_HOST", "localhost"), // Mudado para localhost por padrão
		Port:     5432,
		User:     getEnv("DB_USER", "postgres"),
		Password: getEnv("DB_PASSWORD", "postgres"),
		DBName:   getEnv("DB_NAME", "postgres"),
		SSLMode:  getEnv("DB_SSLMODE", "disable"),
	}
}

// getEnv retorna o valor da variável de ambiente ou o valor padrão
func getEnv(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}
