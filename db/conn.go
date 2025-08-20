package db

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

// DatabaseInterface define o contrato para operações de banco
type DatabaseInterface interface {
	Ping() error
	Close() error
	Query(query string, args ...interface{}) (*sql.Rows, error)
	QueryRow(query string, args ...interface{}) *sql.Row
	Prepare(query string) (*sql.Stmt, error)
}

// ConnectDB conecta ao banco de dados usando a configuração fornecida
func ConnectDB(config *Config) (*sql.DB, error) {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=%s",
		config.Host, config.Port, config.User, config.Password, config.DBName, config.SSLMode)

	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	return db, nil
}

// ConnectDBWithDefaults conecta ao banco usando configurações padrão (mantém compatibilidade)
func ConnectDBWithDefaults() (*sql.DB, error) {
	config := NewConfig()
	return ConnectDB(config)
}
