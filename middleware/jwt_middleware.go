package middleware

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"github.com/joho/godotenv"
)

type Claims struct {
	Username string `json:"username"`
	Role     string `json:"role"`
	jwt.RegisteredClaims
}

func getJWTSecret() []byte {
	if err := godotenv.Load(); err != nil {
		log.Println("Warning: .env file not found, using default JWT secret")
	}

	key := os.Getenv("JWT_SECRET")
	if key == "" {
		log.Println("Error: JWT_SECRET is missing in .env")
		os.Exit(1)
	}

	return []byte(key)
}

// AuthMiddleware middleware untuk validasi JWT
func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.JSON(http.StatusUnauthorized, gin.H{
				"status":  "failed",
				"message": "Unauthorized",
			})
			c.Abort()
			return
		}

		tokenString := strings.TrimPrefix(authHeader, "Bearer ")
		if tokenString == authHeader {
			c.JSON(http.StatusUnauthorized, gin.H{
				"status":  "failed",
				"message": "Bearer token required",
			})
			c.Abort()
			return
		}

		secretKey := getJWTSecret()
		if secretKey == nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"status":  "failed",
				"message": "JWT secret key not set",
			})
			c.Abort()
			return
		}

		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
			}
			return secretKey, nil
		})

		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{
				"status":  "failed",
				"message": "Invalid or expired token",
			})
			c.Abort()
			return
		}

		if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
			c.Set("username", claims["username"])
			c.Set("role", claims["role"])
		}

		c.Next()
	}
}
