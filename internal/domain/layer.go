package domain

import (
	"gorm.io/datatypes"
	"gorm.io/gorm"
)

type Layer struct {
	gorm.Model
	ID           uint           `gorm:"primaryKey;autoIncrement"`
	ImageID      uint           `json:"imageID" gorm:"not null"`
	Name         string         `json:"name" gorm:"not null"`
	Color        string         `json:"color"`
	Measurements datatypes.JSON `json:"measurements" gorm:"type:jsonb"`
}
