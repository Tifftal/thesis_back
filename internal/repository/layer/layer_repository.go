package layer

import (
	"context"
	"errors"
	"fmt"
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
	updates := make(map[string]interface{})

	fmt.Println(layer.Color)

	if layer.Name != "" {
		updates["name"] = layer.Name
	}
	if layer.Measurements != nil {
		updates["measurements"] = layer.Measurements
	}
	if layer.Color != "" {
		updates["color"] = layer.Color
	}

	result := r.db.Model(&domain.Layer{}).
		Where("id = ?", layer.ID).
		Updates(updates).
		First(&layer)

	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return domain.ErrLayerNotFound
		}

		return result.Error
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
