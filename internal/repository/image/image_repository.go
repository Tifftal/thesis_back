package image

import (
	"context"
	"fmt"
	"github.com/minio/minio-go/v7"
	"gorm.io/gorm"
	"mime/multipart"
	"thesis_back/internal/domain"
)

type imageRepository struct {
	db      *gorm.DB
	s3      *minio.Client
	baseUrl string
}

type IImageRepository interface {
	UploadImage(ctx context.Context, fileName string, file multipart.File, size int64) (string, error)
	CreateImage(ctx context.Context, image *domain.Image) (*domain.Image, error)
}

func NewImageRepository(db *gorm.DB, s3 *minio.Client, baseUrl string) IImageRepository {
	return &imageRepository{
		db:      db,
		s3:      s3,
		baseUrl: baseUrl,
	}
}

func (r *imageRepository) UploadImage(ctx context.Context, fileName string, file multipart.File, size int64) (string, error) {
	_, err := r.s3.PutObject(
		ctx,
		"thesis",
		fileName,
		file,
		size,
		minio.PutObjectOptions{},
	)
	if err != nil {
		return "", err
	}

	return fmt.Sprintf("%s/%s", r.baseUrl, fileName), nil
}

func (r *imageRepository) CreateImage(ctx context.Context, image *domain.Image) (*domain.Image, error) {
	if err := r.db.WithContext(ctx).Create(image).Error; err != nil {
		return nil, err
	}

	return image, nil
}
