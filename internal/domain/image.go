package domain

import "gorm.io/gorm"

type Image struct {
	gorm.Model
	Name      string  `json:"name" gorm:"not null"`
	FileName  string  `json:"fileName" gorm:"not null"`
	ProjectID uint    `json:"projectID" gorm:"not null"`
	Layers    []Layer `json:"layer" gorm:"foreignKey:ImageID"`
	URL       string  `json:"url" gorm:"not null"`
	Width     int64   `json:"width" gorm:"not null"`
	Units     string  `json:"units" gorm:"not null"`
}
