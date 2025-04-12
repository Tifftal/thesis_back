package middleware

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"net/http"
	"strings"
	"thesis_back/internal/domain"
	"thesis_back/internal/service"
)

func IsAuthenticated(authService *service.AuthService, logger *zap.Logger) gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")

		fmt.Println(authHeader)
		if authHeader == "" {
			logger.Warn("Authorization Header is Empty", zap.String("AuthorizationHeader", c.Request.Header.Get("Authorization")))
			c.AbortWithStatusJSON(http.StatusUnauthorized, domain.ErrUnauthorized.Error())
			return
		}

		headerParts := strings.Split(authHeader, " ")
		fmt.Println(headerParts, len(headerParts), headerParts[0])
		if len(headerParts) != 2 || headerParts[0] != "Bearer" {
			logger.Warn("Authorization Header is not Bearer")
			c.AbortWithStatusJSON(http.StatusUnauthorized, domain.ErrUnauthorized.Error())
			return
		}

		tokenString := headerParts[1]

		userID, err := authService.ValidateToken(tokenString)
		if err != nil {
			logger.Warn("Authorization Token Validation Error", zap.Error(err))
			c.AbortWithStatusJSON(http.StatusUnauthorized, domain.ErrUnauthorized.Error())
			return
		}
		fmt.Println(userID)
		c.Set("userID", userID)
		c.Next()
	}
}
