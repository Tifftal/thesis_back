package domain

import (
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
	"time"
)

type User struct {
	gorm.Model
	ID         uint      `gorm:"primaryKey;autoIncrement"`
	Username   string    `json:"username" gorm:"not null;unique"`
	Password   string    `json:"password" gorm:"not null"`
	FirstName  string    `json:"firstName" gorm:"not null"`
	LastName   string    `json:"lastName" gorm:"not null"`
	Patronymic string    `json:"patronymic,omitempty"`
	Projects   []Project `json:"projects" gorm:"foreignKey:UserID"`
}

type TokenPair struct {
	AccessToken  string    `json:"accessToken"`
	RefreshToken string    `json:"refreshToken"`
	ExpiresAt    time.Time `json:"expiresAt"`
}

func (u *User) SetPassword(password string) error {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	u.Password = string(hash)
	return nil
}

func (u *User) CheckPassword(password string) error {
	return bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(password))
}
