package project

import (
	"context"
	"go.uber.org/zap"
	"thesis_back/internal/domain"
	"thesis_back/internal/repository/project"
)

type projectUseCase struct {
	repo   *project.IProjectRepository
	logger *zap.Logger
}

type IProjectUseCase interface {
	Create(ctx context.Context, project *domain.Project) (*domain.Project, error)
	Update(ctx context.Context, project *domain.Project) (*domain.Project, error)
	Delete(ctx context.Context, project *domain.Project) error
	Get(ctx context.Context, project *domain.Project) (*domain.Project, error)
	GetByID(ctx context.Context, project *domain.Project) (*domain.Project, error)
}

func NewProjectUseCase(repo *project.IProjectRepository, logger *zap.Logger) IProjectUseCase {
	return &projectUseCase{
		repo:   repo,
		logger: logger.Named("ProjectUseCase"),
	}
}

func (p projectUseCase) Create(ctx context.Context, project *domain.Project) (*domain.Project, error) {
	//TODO implement me
	panic("implement me")
}

func (p projectUseCase) Update(ctx context.Context, project *domain.Project) (*domain.Project, error) {
	//TODO implement me
	panic("implement me")
}

func (p projectUseCase) Delete(ctx context.Context, project *domain.Project) error {
	//TODO implement me
	panic("implement me")
}

func (p projectUseCase) Get(ctx context.Context, project *domain.Project) (*domain.Project, error) {
	//TODO implement me
	panic("implement me")
}

func (p projectUseCase) GetByID(ctx context.Context, project *domain.Project) (*domain.Project, error) {
	//TODO implement me
	panic("implement me")
}
