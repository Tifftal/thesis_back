package domain

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Username   string
	Password   string
	FirstName  string
	LastName   string
	Patronymic string
}
