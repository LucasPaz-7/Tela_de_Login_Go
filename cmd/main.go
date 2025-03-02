package main

import (
	"Api-Go-Secretaria/controller"
	"Api-Go-Secretaria/db"
	"Api-Go-Secretaria/repository"
	"Api-Go-Secretaria/usecase"
	"log"
	"os"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	// Conecta ao banco de dados
	dbConnection, err := db.ConnectDB()
	if err != nil {
		log.Fatalf("Erro ao conectar ao banco: %v", err)
	}

	// Inicializa as camadas para a API de Login
	// Camada de repositório
	userRepository := repository.NewUserRepository(dbConnection)

	// Camada de usecase
	userUseCase := usecase.NewUserUseCase(userRepository)

	// Camada de controller
	loginController := controller.NewLoginController(userUseCase)

	// Rotas
	router.POST("/login", loginController.Login)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	router.Run(":" + port)
}