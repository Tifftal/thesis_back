package project

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"thesis_back/internal/usecase/project"
)

type ProjectHandler struct {
	pu     *project.IProjectUseCase
	logger *zap.Logger
}

func NewProjectHandler(pu *project.IProjectUseCase, logger *zap.Logger) *ProjectHandler {
	return &ProjectHandler{
		pu:     pu,
		logger: logger.Named("ProjectHandler"),
	}
}

func (h *ProjectHandler) Create(c *gin.Context) {

}

func (h *ProjectHandler) Update(c *gin.Context) {

}

func (h *ProjectHandler) Delete(c *gin.Context) {

}

func (h *ProjectHandler) Get(c *gin.Context) {

}

func (h *ProjectHandler) GetByID(c *gin.Context) {

}
