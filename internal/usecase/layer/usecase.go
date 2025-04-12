package layer

import (
	"context"
	"go.uber.org/zap"
	"thesis_back/internal/domain"
	"thesis_back/internal/repository/layer"
)

type layerUseCase struct {
	repo   layer.ILayerRepository
	logger *zap.Logger
}

type ILayerUseCase interface {
	Create(ctx context.Context, layer *domain.Layer) (*domain.Layer, error)
	Update(ctx context.Context, layer *domain.Layer) (*domain.Layer, error)
	Delete(ctx context.Context, id uint) error
}

func NewLayerUseCase(repo layer.ILayerRepository, logger *zap.Logger) ILayerUseCase {
	return &layerUseCase{
		repo:   repo,
		logger: logger.Named("LayerUseCase"),
	}
}

func (lu *layerUseCase) Create(ctx context.Context, layer *domain.Layer) (*domain.Layer, error) {
	if err := lu.repo.Create(ctx, layer); err != nil {
		return nil, err
	}

	return layer, nil
}

func (lu *layerUseCase) Update(ctx context.Context, layer *domain.Layer) (*domain.Layer, error) {
	if err := lu.repo.Update(ctx, layer); err != nil {
		return nil, err
	}

	return layer, nil
}

func (lu *layerUseCase) Delete(ctx context.Context, id uint) error {
	if err := lu.repo.Delete(ctx, id); err != nil {
		return err
	}

	return nil
}
