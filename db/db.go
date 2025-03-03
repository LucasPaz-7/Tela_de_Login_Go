package db

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func ConnectDB() (*gorm.DB, error) {
	// Configuração fixa para desenvolvimento local
	dsn := "host=localhost user=postgres password=1234 dbname=postgres port=5434 "

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