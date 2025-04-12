package project

import (
	"context"
	"go.uber.org/zap"
	"thesis_back/internal/domain"
	"thesis_back/internal/repository/project"
)

type projectUseCase struct {
	repo   project.IProjectRepository
	logger *zap.Logger
}

type IProjectUseCase interface {
	Create(ctx context.Context, project *domain.Project) (*domain.Project, error)
	Update(ctx context.Context, project *domain.Project) (*domain.Project, error)
	Delete(ctx context.Context, id uint) error
	Get(ctx context.Context) ([]*domain.Project, error)
	GetByID(ctx context.Context, id uint) (*domain.Project, error)
}

func NewProjectUseCase(repo project.IProjectRepository, logger *zap.Logger) IProjectUseCase {
	return &projectUseCase{
		repo:   repo,
		logger: logger.Named("ProjectUseCase"),
	}
}

func (p projectUseCase) Create(ctx context.Context, project *domain.Project) (*domain.Project, error) {
	if err := p.repo.Create(ctx, project); err != nil {
		return nil, err
	}

	return project, nil
}

func (p projectUseCase) Update(ctx context.Context, project *domain.Project) (*domain.Project, error) {
	if err := p.repo.Update(ctx, project); err != nil {
		return nil, err
	}

	return project, nil
}

func (p projectUseCase) Delete(ctx context.Context, id uint) error {
	if err := p.repo.Delete(ctx, id); err != nil {
		return err
	}

	return nil
}

func (p projectUseCase) Get(ctx context.Context) ([]*domain.Project, error) {
	projects, err := p.repo.Get(ctx)
	if err != nil {
		return nil, err
	}

	return projects, nil
}

func (p projectUseCase) GetByID(ctx context.Context, id uint) (*domain.Project, error) {
	project, err := p.repo.GetByID(ctx, id)
	if err != nil {
		return nil, err
	}

	return project, nil
}
