package domain

import "gorm.io/gorm"

type Image struct {
	gorm.Model
	Name      string  `json:"name" gorm:"not null"`
	ProjectID uint    `json:"projectID" gorm:"not null"`
	Layers    []Layer `json:"layer" gorm:"foreignKey:ImageID"`
}
