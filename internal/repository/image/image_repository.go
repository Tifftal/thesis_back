package image

import (
	"context"
	"errors"
	"fmt"
	"github.com/minio/minio-go/v7"
	"gorm.io/gorm"
	"io"
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
	LoadImage(ctx context.Context, id uint) ([]byte, error)
	CreateImage(ctx context.Context, image *domain.Image) (*domain.Image, error)
	Update(ctx context.Context, name string, units string, width int64, id uint) (*domain.Image, error)
	Delete(ctx context.Context, id uint) error
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

func (r *imageRepository) LoadImage(ctx context.Context, id uint) ([]byte, error) {
	var image domain.Image

	err := r.db.Where("id = ?", id).First(&image).Error
	if err != nil {
		return nil, err
	}

	obj, err := r.s3.GetObject(ctx, "thesis", image.FileName, minio.GetObjectOptions{})
	if err != nil {
		return nil, err
	}
	defer obj.Close()

	imageData, err := io.ReadAll(obj)
	if err != nil {
		return nil, err
	}

	return imageData, nil
}

func (r *imageRepository) CreateImage(ctx context.Context, image *domain.Image) (*domain.Image, error) {
	if err := r.db.WithContext(ctx).Create(image).Error; err != nil {
		return nil, err
	}

	return image, nil
}

func (r *imageRepository) Update(ctx context.Context, name, units string, width int64, id uint) (*domain.Image, error) {
	updates := make(map[string]interface{})

	if name != "" {
		updates["name"] = name
	}
	if units != "" {
		updates["units"] = units
	}
	if width != 0 {
		updates["width"] = width
	}

	if len(updates) == 0 {
		return nil, errors.New("no fields to update")
	}

	var image domain.Image
	result := r.db.Model(&domain.Image{}).
		Where("id = ?", id).
		Updates(updates).
		First(&image)

	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, domain.ErrImageNotFound
		}
		return nil, result.Error
	}

	return &image, nil
}

func (r *imageRepository) Delete(ctx context.Context, id uint) error {
	err := r.db.Where("id = ?", id).Delete(&domain.Image{}).Error
	if err != nil {
		return err
	}

	return nil
}
