package layer

import (
	"go.uber.org/zap"
	"thesis_back/internal/repository/layer"
)

type layerUseCase struct {
	repo   layer.ILayerRepository
	logger *zap.Logger
}

type ILayerUseCase interface {
}

func NewLayerUseCase(repo *layer.ILayerRepository, logger *zap.Logger) ILayerUseCase {
	return &layerUseCase{
		repo:   repo,
		logger: logger.Named("LayerUseCase"),
	}
}
