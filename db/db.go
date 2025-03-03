package db

import (
	"fmt"
	"os"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func ConnectDB() (*gorm.DB, error) {
	// Obter configurações do banco de dados das variáveis de ambiente
	host := os.Getenv("DB_HOST")
	if host == "" {
		host = "localhost"
	}
	
	user := os.Getenv("DB_USER")
	if user == "" {
		user = "postgres"
	}
	
	password := os.Getenv("DB_PASSWORD")
	if password == "" {
		password = "1234"
	}
	
	dbname := os.Getenv("DB_NAME")
	if dbname == "" {
		dbname = "postgres"
	}
	
	port := os.Getenv("DB_PORT")
	if port == "" {
		port = "5434"
	}

	// Montar string de conexão
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s", 
		host, user, password, dbname, port)

	// Estabelecer conexão
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	// Testar a conexão
	sqlDB, _ := db.DB()
	err = sqlDB.Ping()
	if err != nil {
		return nil, err
	}

	return db, nil
}