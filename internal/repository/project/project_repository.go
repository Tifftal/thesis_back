package project

import (
	"context"
	"gorm.io/gorm"
	"thesis_back/internal/domain"
)

type ProjectRepository struct {
	db *gorm.DB
}

type IProjectRepository interface {
	Create(ctx context.Context)
	GetByID(ctx context.Context, id string) (*domain.Project, error)
	Update(ctx context.Context, project *domain.Project) error
	Delete(ctx context.Context, id string) error
}
