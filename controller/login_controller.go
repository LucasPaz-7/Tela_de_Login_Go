package controller

import (
	
    "github.com/LucasPaz-7/Secretaria_Api_Go/usecase"
    "net/http"
    
    "github.com/gin-gonic/gin"
)

type LoginController struct {
    useCase *usecase.UserUseCase
}

func NewLoginController(useCase *usecase.UserUseCase) *LoginController {
    return &LoginController{useCase: useCase}
}

type LoginRequest struct {
    Email    string `json:"email"`
    Password string `json:"password"`
}

func (c *LoginController) Login(ctx *gin.Context) {
    var req LoginRequest
    
    if err := ctx.ShouldBindJSON(&req); err != nil {
        ctx.JSON(http.StatusBadRequest, gin.H{"error": "Dados inválidos"})
        return
    }

    user, err := c.useCase.Login(req.Email, req.Password)
    if err != nil {
        ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Credenciais incorretas"})
        return
    }

    ctx.JSON(http.StatusOK, gin.H{
        "message": "Login bem-sucedido",
        "user":    user,
    })
}