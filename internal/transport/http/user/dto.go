package user

import (
	"thesis_back/internal/domain"
	"time"
)

type CreateUserDTO struct {
	Username   string `json:"username" binding:"required"`
	Password   string `json:"password" binding:"required,min=8"`
	FirstName  string `json:"firstName" binding:"required"`
	LastName   string `json:"lastName" binding:"required"`
	Patronymic string `json:"patronymic,omitempty"`
}

type LoginUserDTO struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required,min=8"`
}

type UserResponse struct {
	ID         uint      `json:"id"`
	Username   string    `json:"username"`
	FirstName  string    `json:"firstName"`
	LastName   string    `json:"lastName"`
	Patronymic string    `json:"patronymic"`
	CreatedAt  time.Time `json:"createdAt"`
	UpdatedAt  time.Time `json:"updatedAt"`
}

type AuthResponse struct {
	User         UserResponse `json:"user"`
	AccessToken  string       `json:"accessToken"`
	RefreshToken string       `json:"refreshToken"`
	ExpiresAt    time.Time    `json:"expiresAt"`
}

type ErrorResponse struct {
	Message string `json:"message"`
}

func ToUserResponse(id uint, username, firstName, lastName, patronymic string, createdAt, updatedAt time.Time) UserResponse {
	return UserResponse{
		ID:         id,
		Username:   username,
		FirstName:  firstName,
		LastName:   lastName,
		Patronymic: patronymic,
		CreatedAt:  createdAt,
		UpdatedAt:  updatedAt,
	}
}

func ToAuthResponse(tokens *domain.TokenPair, user *domain.User) AuthResponse {
	return AuthResponse{
		AccessToken:  tokens.AccessToken,
		RefreshToken: tokens.RefreshToken,
		ExpiresAt:    tokens.ExpiresAt,
		User:         ToUserResponse(user.ID, user.Username, user.FirstName, user.LastName, user.Patronymic, user.CreatedAt, user.UpdatedAt),
	}
}
