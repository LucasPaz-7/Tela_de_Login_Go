package db

import (
	"fmt"
	"os"
	"errors"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func ConnectDB() (*gorm.DB, error) {
	// Obter configurações do banco de dados das variáveis de ambiente
	host := os.Getenv("DB_HOST")
	if host == "" {
		return nil, errors.New("variável de ambiente DB_HOST não definida")
	}
	
	user := os.Getenv("DB_USER")
	if user == "" {
		return nil, errors.New("variável de ambiente DB_USER não definida")
	}
	
	password := os.Getenv("DB_PASSWORD")
	if password == "" {
		return nil, errors.New("variável de ambiente DB_PASSWORD não definida")
	}
	
	dbname := os.Getenv("DB_NAME")
	if dbname == "" {
		return nil, errors.New("variável de ambiente DB_NAME não definida")
	}
	
	port := os.Getenv("DB_PORT")
	if port == "" {
		return nil, errors.New("variável de ambiente DB_PORT não definida")
	}

	// Montar string de conexão
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=require", 
		host, user, password, dbname, port)

	// Estabelecer conexão
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, fmt.Errorf("erro ao conectar ao banco de dados: %v", err)
	}

	// Testar a conexão
	sqlDB, _ := db.DB()
	err = sqlDB.Ping()
	if err != nil {
		return nil, fmt.Errorf("erro ao testar conexão com o banco: %v", err)
	}

	return db, nil
}