package service

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
	"thesis_back/internal/domain"
)

type JWTConfig struct {
	SecretKey     string
	AccessExpiry  time.Duration
	RefreshExpiry time.Duration
}

type AuthService struct {
	config *JWTConfig
}

func NewAuthService(cfg *JWTConfig) *AuthService {
	return &AuthService{config: cfg}
}

func (a *AuthService) GenerateTokens(user *domain.User) (*domain.TokenPair, error) {
	// Access token
	accessClaims := jwt.MapClaims{
		"userID": user.ID,
		"exp":    time.Now().Add(a.config.AccessExpiry).Unix(),
	}

	accessToken := jwt.NewWithClaims(jwt.SigningMethodHS256, accessClaims)
	accessString, err := accessToken.SignedString([]byte(a.config.SecretKey))
	if err != nil {
		return nil, err
	}

	// Refresh token
	refreshClaims := jwt.MapClaims{
		"userID": user.ID,
		"exp":    time.Now().Add(a.config.RefreshExpiry).Unix(),
	}

	refreshToken := jwt.NewWithClaims(jwt.SigningMethodHS256, refreshClaims)
	refreshString, err := refreshToken.SignedString([]byte(a.config.SecretKey))
	if err != nil {
		return nil, err
	}

	return &domain.TokenPair{
		AccessToken:  accessString,
		RefreshToken: refreshString,
		ExpiresAt:    time.Now().Add(a.config.AccessExpiry),
	}, nil
}

func (a *AuthService) ValidateToken(tokenString string) (uint, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte(a.config.SecretKey), nil
	})

	if err != nil || !token.Valid {
		return 0, domain.ErrUnauthorized
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return 0, domain.ErrUnauthorized
	}

	userIDFloat, ok := claims["userID"].(float64)
	if !ok {
		return 0, domain.ErrUnauthorized
	}

	return uint(userIDFloat), nil
}
