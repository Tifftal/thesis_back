package image

import (
	"github.com/minio/minio-go/v7"
	"gorm.io/gorm"
)

type imageRepository struct {
	db *gorm.DB
	s3 *minio.Client
}

type IImageRepository interface {
}

func NewImageRepository(db *gorm.DB, s3 *minio.Client) IImageRepository {
	return &imageRepository{
		db: db,
		s3: s3,
	}
}
