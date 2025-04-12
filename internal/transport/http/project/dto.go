package project

type CreateProjectDTO struct {
	Name string `json:"name" binding:"required"`
}
