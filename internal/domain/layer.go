package domain

import (
	"gorm.io/datatypes"
	"gorm.io/gorm"
)

type Layer struct {
	gorm.Model
	ImageID      uint           `json:"imageID" gorm:"not null"`
	Measurements datatypes.JSON `json:"measurements" gorm:"type:jsonb"` // jsonb для PostgreSQL
}
