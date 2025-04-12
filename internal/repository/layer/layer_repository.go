package layer

import "gorm.io/gorm"

type layerRepository struct {
	db *gorm.DB
}

type ILayerRepository interface {
}

func NewLayerRepository(db *gorm.DB) ILayerRepository {
	return &layerRepository{
		db: db,
	}
}
