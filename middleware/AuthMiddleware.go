package middleware

import (
	"AgringBackend/utils"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Получаем заголовок Authorization
		tokenString := c.GetHeader("Authorization")
		if tokenString == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Authorization token not provided"})
			c.Abort() // Прерываем выполнение запроса
			return
		}

		// Убираем возможный префикс "Bearer "
		tokenString = strings.TrimPrefix(tokenString, "Bearer ")

		// Проверяем и декодируем токен
		claims, err := utils.ValidateToken(tokenString)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
			c.Abort()
			return
		}

		// Передаём данные пользователя в контекст запроса
		c.Set("username", claims.Username)

		// Продолжаем обработку запроса
		c.Next()
	}
}
