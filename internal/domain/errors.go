package domain

import "errors"

var (
	ErrUserNotFound       = errors.New("Пользователь не найден")
	ErrInvalidCredentials = errors.New("Неверные имя пользователя или пароль")
	ErrUserExists         = errors.New("Пользователь с таким именем уже существует")
	ErrUnauthorized       = errors.New("Неавторизован")
	ErrInvalidRequestBody = errors.New("Данные неверны")
	ErrProjectNotFound    = errors.New("Проект не найден")
	ErrImageNotUploaded   = errors.New("Не удалось загрузить изображение в S3")
	ErrImageNotFound      = errors.New("Изображение не найдено")
	ErrImageNotOpens      = errors.New("Ошибка при открытии изображения")
	ErrLayerNotFound      = errors.New("Слоай не найден")
)
