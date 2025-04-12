package layer

import (
	"context"
	"gorm.io/gorm"
	"thesis_back/internal/domain"
)

type layerRepository struct {
	db *gorm.DB
}

type ILayerRepository interface {
	Create(ctx context.Context, layer *domain.Layer) error
	Update(ctx context.Context, layer *domain.Layer) error
	Delete(ctx context.Context, id uint) error
}

func NewLayerRepository(db *gorm.DB) ILayerRepository {
	return &layerRepository{
		db: db,
	}
}

func (r *layerRepository) Create(ctx context.Context, layer *domain.Layer) error {
	err := r.db.Create(layer).Error
	if err != nil {
		return err
	}

	return nil
}

func (r *layerRepository) Update(ctx context.Context, layer *domain.Layer) error {
	err := r.db.Model(&domain.Layer{}).Where("id = ?", layer.ID).Update("name", layer.Name).Update("measurements", layer.Measurements).Error
	if err != nil {
		return err
	}

	return nil
}

func (r *layerRepository) Delete(ctx context.Context, id uint) error {
	err := r.db.Where("id = ?", id).Delete(&domain.Layer{}).Error
	if err != nil {
		return err
	}

	return nil
}
