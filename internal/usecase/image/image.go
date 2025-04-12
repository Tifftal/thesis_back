package image

import (
	"go.uber.org/zap"
	"thesis_back/internal/repository/image"
)

type imageUseCase struct {
	repo   image.IImageRepository
	logger *zap.Logger
}

type IImageUseCase interface {
}

func NewImageUseCase(repo *image.IImageRepository, logger *zap.Logger) IImageUseCase {
	return &imageUseCase{
		repo:   repo,
		logger: logger.Named("ImageUseCase"),
	}
}
