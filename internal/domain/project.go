package domain

import "gorm.io/gorm"

type Project struct {
	gorm.Model
	ID     uint    `gorm:"primaryKey;autoIncrement"`
	Name   string  `json:"name"`
	UserID uint    `json:"userID"`
	Images []Image `gorm:"foreignKey:ProjectID" json:"images"`
}
