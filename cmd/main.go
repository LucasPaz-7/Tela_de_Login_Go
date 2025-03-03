package main

import (
	"log"
	"os"
	"time"

	"github.com/LucasPaz-7/Secretaria_Api_Go/controller"
	"github.com/LucasPaz-7/Secretaria_Api_Go/db"
	"github.com/LucasPaz-7/Secretaria_Api_Go/repository"
	"github.com/LucasPaz-7/Secretaria_Api_Go/usecase"

	"github.com/LucasPaz-7/Secretaria_Api_Go/model"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
    router := gin.Default()

    // Conecta ao banco de dados
    dbConnection, err := db.ConnectDB()
    if err != nil {
        log.Fatalf("Erro ao conectar ao banco: %v", err)
    }

    router.Use(cors.New(cors.Config{
        AllowOrigins:     []string{"http://localhost:5173", "http://localhost:8080", "https://secretaria-frontend.vercel.app"}, 
        AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
        AllowHeaders:     []string{"Origin", "Content-Type"},
        ExposeHeaders:    []string{"Content-Length"},
        AllowCredentials: true,
        MaxAge: 12 * time.Hour,
    }))

    // Cria a tabela de usuários automaticamente
    dbConnection.AutoMigrate(&model.User{})

    // Inicializa as camadas
    userRepository := repository.NewUserRepository(dbConnection)
    userUseCase := usecase.NewUserUseCase(userRepository)
    loginController := controller.NewLoginController(userUseCase)

    // Rotas
    router.GET(("/"), func(c *gin.Context) {
        c.JSON(200, gin.H{  "message": "Secretaria API" })
    })
    router.POST("/login", loginController.Login)

    // Inicia o servidor
    port := os.Getenv("PORT")
    if port == "" {
        port = "8080"
    }
    router.Run(":" + port)
}