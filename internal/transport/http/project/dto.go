package project

import (
	"thesis_back/internal/domain"
	image_dto "thesis_back/internal/transport/http/image"
	"time"
)

type CreateProjectDTO struct {
	Name string `json:"name" binding:"required"`
}

type UpdateProjectDTO struct {
	Name string `json:"name" binding:"required"`
}

type ErrorResponse struct {
	Message string `json:"message"`
}

type ProjectResponse struct {
	ID        uint                      `json:"id"`
	Name      string                    `json:"name"`
	CreatedAt time.Time                 `json:"createdAt"`
	UpdatedAt time.Time                 `json:"updatedAt"`
	Images    []image_dto.ImageResponse `json:"images,omitempty"`
}

func ToProjectResponse(project *domain.Project) ProjectResponse {
	images := make([]image_dto.ImageResponse, len(project.Images))
	for i, image := range project.Images {
		images[i] = image_dto.ToImageResponse(&image)
	}

	return ProjectResponse{
		ID:        project.ID,
		Name:      project.Name,
		CreatedAt: project.CreatedAt,
		UpdatedAt: project.UpdatedAt,
		Images:    images,
	}
}

func ToProjectsResponse(projects []*domain.Project) []ProjectResponse {
	projectsResponse := make([]ProjectResponse, len(projects))

	for i, project := range projects {
		projectsResponse[i] = ProjectResponse{
			ID:        project.ID,
			Name:      project.Name,
			CreatedAt: project.CreatedAt,
			UpdatedAt: project.UpdatedAt,
		}
	}

	return projectsResponse
}
