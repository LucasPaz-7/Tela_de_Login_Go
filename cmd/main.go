package main

import (
	"github.com/LucasPaz-7/Secretaria_Api_Go/controller"
	"github.com/LucasPaz-7/Secretaria_Api_Go/db"
	"github.com/LucasPaz-7/Secretaria_Api_Go/repository"
	"github.com/LucasPaz-7/Secretaria_Api_Go/usecase"
	"log"
	"os"

	"github.com/LucasPaz-7/Secretaria_Api_Go/model"
	"github.com/gin-gonic/gin"
)

func main() {
    router := gin.Default()

    // Conecta ao banco de dados
    dbConnection, err := db.ConnectDB()
    if err != nil {
        log.Fatalf("Erro ao conectar ao banco: %v", err)
    }

    // Cria a tabela de usuários automaticamente
    dbConnection.AutoMigrate(&model.User{})

    // Inicializa as camadas
    userRepository := repository.NewUserRepository(dbConnection)
    userUseCase := usecase.NewUserUseCase(userRepository)
    loginController := controller.NewLoginController(userUseCase)

    // Rotas
    router.POST("/login", loginController.Login)

    // Inicia o servidor
    port := os.Getenv("PORT")
    if port == "" {
        port = "8080"
    }
    router.Run(":" + port)
}