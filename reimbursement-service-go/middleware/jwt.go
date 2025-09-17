package middleware

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

var JwtSecret = []byte("kunci-rahasia-yang-sangat-panjang-dan-sulit-ditebak")

func JWTAuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {

		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Authorization header missing"})
			c.Abort()
			return
		}

		parts := strings.Split(authHeader, " ")
		if len(parts) != 2 || parts[0] != "Bearer" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid Authorization header"})
			c.Abort()
			return
		}

		tokenString := parts[1]

		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			return JwtSecret, nil
		})
		if err != nil || !token.Valid {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
			c.Abort()
			return
		}

		claims := token.Claims.(jwt.MapClaims)

		// Simpan email
		c.Set("user_email", claims["sub"])

		// Ambil roles dari claims
		if roles, ok := claims["roles"].([]interface{}); ok && len(roles) > 0 {
			// Ambil role pertama sebagai default
			c.Set("user_role", roles[0].(string))
		} else {
			// fallback jika roles tidak ada
			c.Set("user_role", "")
		}

		fmt.Println("Claims:", claims)
		fmt.Println("=== DEBUG END ===")

		c.Next()
	}
}
