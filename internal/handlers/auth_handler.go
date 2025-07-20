package handlers

import (
    "jwt-auth-crud/internal/dto"
    "jwt-auth-crud/internal/services"
    "net/http"

    "github.com/gin-gonic/gin"
)

type AuthHandler interface {
    Register(c *gin.Context)
    Login(c *gin.Context)
}

type authHandler struct {
    authService services.AuthService
}

func NewAuthHandler(authService services.AuthService) AuthHandler {
    return &authHandler{
        authService: authService,
    }
}

func (h *authHandler) Register(c *gin.Context) {
    var req dto.RegisterRequest
    if err := c.ShouldBindJSON(&req); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    response, err := h.authService.Register(&req)
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    c.JSON(http.StatusCreated, gin.H{
        "message": "User registered successfully",
        "data":    response,
    })
}

func (h *authHandler) Login(c *gin.Context) {
    var req dto.LoginRequest
    if err := c.ShouldBindJSON(&req); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    response, err := h.authService.Login(&req)
    if err != nil {
        c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
        return
    }

    c.JSON(http.StatusOK, gin.H{
        "message": "Login successful",
        "data":    response,
    })
}
