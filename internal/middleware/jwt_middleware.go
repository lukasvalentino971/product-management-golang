package middleware

import (
    "jwt-auth-crud/internal/utils"
    "net/http"
    "strings"

    "github.com/gin-gonic/gin"
)

type JWTMiddleware interface {
    ValidateToken() gin.HandlerFunc
}

type jwtMiddleware struct {
    jwtSecret string
}

func NewJWTMiddleware(jwtSecret string) JWTMiddleware {
    return &jwtMiddleware{
        jwtSecret: jwtSecret,
    }
}

func (m *jwtMiddleware) ValidateToken() gin.HandlerFunc {
    return func(c *gin.Context) {
        authHeader := c.GetHeader("Authorization")
        if authHeader == "" {
            c.JSON(http.StatusUnauthorized, gin.H{"error": "Authorization header required"})
            c.Abort()
            return
        }

        // Check Bearer prefix
        tokenParts := strings.Split(authHeader, " ")
        if len(tokenParts) != 2 || tokenParts[0] != "Bearer" {
            c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid authorization format"})
            c.Abort()
            return
        }

        // Validate token
        claims, err := utils.ValidateJWT(tokenParts[1], m.jwtSecret)
        if err != nil {
            c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
            c.Abort()
            return
        }

        // Set user info in context
        c.Set("user_id", claims.UserID)
        c.Set("user_email", claims.Email)
        c.Set("user_role", claims.Role)

        c.Next()
    }
}
