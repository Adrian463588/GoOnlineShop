package middleware

import (
	"ecommerce-backend/pkg/utils"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			// Try to get from custom header "token" as per postman collection sometimes
			authHeader = c.GetHeader("token")
		}

		if authHeader == "" {
			utils.APIResponse(c, http.StatusUnauthorized, false, "Unauthorized", nil, []string{"No token found"})
			c.Abort()
			return
		}

		// Handle Bearer prefix if present
		tokenString := strings.Replace(authHeader, "Bearer ", "", 1)

		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			return utils.SecretKey, nil
		})

		if err != nil || !token.Valid {
			utils.APIResponse(c, http.StatusUnauthorized, false, "Unauthorized", nil, []string{"Invalid token"})
			c.Abort()
			return
		}

		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok {
			utils.APIResponse(c, http.StatusUnauthorized, false, "Unauthorized", nil, []string{"Invalid token claims"})
			c.Abort()
			return
		}

		// Set context
		c.Set("user_id", uint(claims["user_id"].(float64)))
		c.Set("is_admin", claims["is_admin"].(bool))

		c.Next()
	}
}

func AdminOnly() gin.HandlerFunc {
	return func(c *gin.Context) {
		isAdmin, exists := c.Get("is_admin")
		if !exists || !isAdmin.(bool) {
			utils.APIResponse(c, http.StatusForbidden, false, "Forbidden", nil, []string{"Admin access required"})
			c.Abort()
			return
		}
		c.Next()
	}
}