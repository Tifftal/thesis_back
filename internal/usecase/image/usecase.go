package image

import (
	"context"
	"fmt"
	"go.uber.org/zap"
	"mime/multipart"
	"thesis_back/internal/domain"
	"thesis_back/internal/repository/image"
	"time"
)

type imageUseCase struct {
	repo   image.IImageRepository
	logger *zap.Logger
}

type IImageUseCase interface {
	UploadImage(ctx context.Context, name string, units string, width int64, imageFile *multipart.FileHeader, projectID uint) (*domain.Image, error)
	LoadImage(ctx context.Context, id uint) ([]byte, error)
	Update(ctx context.Context, name string, units string, width int64, id uint) (*domain.Image, error)
	Delete(ctx context.Context, id uint) error
}

func NewImageUseCase(repo image.IImageRepository, logger *zap.Logger) IImageUseCase {
	return &imageUseCase{
		repo:   repo,
		logger: logger.Named("ImageUseCase"),
	}
}

func (u *imageUseCase) UploadImage(ctx context.Context, name, units string, width int64, imageFile *multipart.FileHeader, projectID uint) (*domain.Image, error) {
	file, err := imageFile.Open()
	if err != nil {
		return nil, domain.ErrImageNotOpens
	}
	defer file.Close()

	fileName := fmt.Sprintf("%d_%s", time.Now().UnixNano(), imageFile.Filename)

	location, err := u.repo.UploadImage(ctx, fileName, file, imageFile.Size)
	if err != nil {
		u.logger.Warn("Upload error", zap.Error(err))
		return nil, domain.ErrImageNotUploaded
	}

	image, err := u.repo.CreateImage(ctx, &domain.Image{
		FileName:  fileName,
		ProjectID: projectID,
		Name:      name,
		Units:     units,
		Width:     width,
		URL:       location,
	})
	if err != nil {
		u.logger.Warn("Upload error", zap.Error(err))
		return nil, domain.ErrImageNotUploaded
	}

	return image, nil
}

func (u *imageUseCase) LoadImage(ctx context.Context, id uint) ([]byte, error) {
	imageBytes, err := u.repo.LoadImage(ctx, id)
	if err != nil {
		u.logger.Warn("LoadImage error", zap.Error(err))
		return nil, domain.ErrImageNotLoaded
	}

	return imageBytes, nil
}

func (u *imageUseCase) Update(ctx context.Context, name, units string, width int64, id uint) (*domain.Image, error) {
	image, err := u.repo.Update(ctx, name, units, width, id)
	if err != nil {
		u.logger.Warn("Update error", zap.Error(err))
		return nil, err
	}

	return image, nil
}

func (u *imageUseCase) Delete(ctx context.Context, id uint) error {
	if err := u.repo.Delete(ctx, id); err != nil {
		return err
	}

	return nil
}
