package project

import (
	"context"
	"gorm.io/gorm"
	"thesis_back/internal/domain"
)

type projectRepository struct {
	db *gorm.DB
}

type IProjectRepository interface {
	Create(ctx context.Context)
	Get(ctx context.Context) ([]*domain.Project, error)
	GetByID(ctx context.Context, id string) (*domain.Project, error)
	Update(ctx context.Context, project *domain.Project) error
	Delete(ctx context.Context, id string) error
}

func NewProjectRepository(db *gorm.DB) IProjectRepository {
	return &projectRepository{
		db: db,
	}
}

func (p projectRepository) Create(ctx context.Context) {
	//TODO implement me
	panic("implement me")
}

func (p projectRepository) Get(ctx context.Context) ([]*domain.Project, error) {
	//TODO implement me
	panic("implement me")
}

func (p projectRepository) GetByID(ctx context.Context, id string) (*domain.Project, error) {
	//TODO implement me
	panic("implement me")
}

func (p projectRepository) Update(ctx context.Context, project *domain.Project) error {
	//TODO implement me
	panic("implement me")
}

func (p projectRepository) Delete(ctx context.Context, id string) error {
	//TODO implement me
	panic("implement me")
}
