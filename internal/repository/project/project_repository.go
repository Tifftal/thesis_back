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
	Create(ctx context.Context, project *domain.Project) error
	Get(ctx context.Context) ([]*domain.Project, error)
	GetByID(ctx context.Context, id uint) (*domain.Project, error)
	Update(ctx context.Context, project *domain.Project) error
	Delete(ctx context.Context, id uint) error
}

func NewProjectRepository(db *gorm.DB) IProjectRepository {
	return &projectRepository{
		db: db,
	}
}

func (p *projectRepository) Create(ctx context.Context, project *domain.Project) error {
	err := p.db.Create(project).Error
	if err != nil {
		return err
	}

	return nil
}

func (p *projectRepository) GetByID(ctx context.Context, id uint) (*domain.Project, error) {
	var project domain.Project

	err := p.db.
		Preload("Images"). // Предзагрузка изображений
		Where("id = ?", id).
		First(&project).
		Error

	if err != nil {
		return nil, err
	}

	return &project, nil
}

func (p projectRepository) Get(ctx context.Context) ([]*domain.Project, error) {
	var projects []*domain.Project

	err := p.db.Find(&projects).Error
	if err != nil {
		return nil, err
	}

	return projects, nil
}

func (p projectRepository) Update(ctx context.Context, project *domain.Project) error {
	err := p.db.Model(&domain.Project{}).Where("id = ?", project.ID).Update("name", project.Name).Error
	if err != nil {
		return err
	}

	return nil
}

func (p projectRepository) Delete(ctx context.Context, id uint) error {
	err := p.db.Where("id = ?", id).Delete(&domain.Project{}).Error
	if err != nil {
		return err
	}

	return nil
}
