package domain

import "errors"

var (
	ErrUserNotFound       = errors.New("Пользователь не найден")
	ErrInvalidCredentials = errors.New("Неверные имя пользователя или пароль")
	ErrUserExists         = errors.New("Пользователь с таким именем уже существует")
	ErrUnauthorized       = errors.New("Неавторизован")
	ErrInvalidRequestBody = errors.New("Данные неверны")
)
