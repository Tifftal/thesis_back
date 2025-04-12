package project

import (
	"errors"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"net/http"
	"strconv"
	"thesis_back/internal/domain"
	_ "thesis_back/internal/transport/http/image"
	"thesis_back/internal/usecase/project"
)

type ProjectHandler struct {
	pu     project.IProjectUseCase
	logger *zap.Logger
}

func NewProjectHandler(pu project.IProjectUseCase, logger *zap.Logger) *ProjectHandler {
	return &ProjectHandler{
		pu:     pu,
		logger: logger.Named("ProjectHandler"),
	}
}

// Create godoc
// @Summary Создать новый проект
// @Tags Project
// @Security BearerAuth
// @Produce json
// @Param input body CreateProjectDTO true "Название проекта"
// @Success 201 {object} ProjectResponse
// @Failure 400 {object} ErrorResponse
// @Router /project [post]
func (h *ProjectHandler) Create(c *gin.Context) {
	userID := c.GetUint("userID")

	var req CreateProjectDTO
	if err := c.ShouldBindJSON(&req); err != nil {
		h.logger.Warn("Validation error", zap.Error(err))
		c.JSON(http.StatusBadRequest, ErrorResponse{Message: domain.ErrInvalidRequestBody.Error()})
		return
	}

	project, err := h.pu.Create(c.Request.Context(), &domain.Project{
		Name:   req.Name,
		UserID: userID,
	})
	if err != nil {
		h.logger.Error("Create project error", zap.Error(err))
		c.JSON(errorStatusCode(err), ErrorResponse{Message: err.Error()})
		return
	}

	response := ToProjectResponse(project)

	c.JSON(http.StatusCreated, response)
}

// Update godoc
// @Summary Обновить проект
// @Tags Project
// @Security BearerAuth
// @Produce json
// @Param id path int true "ID проекта"
// @Param input body UpdateProjectDTO true "Название проекта"
// @Success 200 {object} ProjectResponse
// @Failure 400 {object} ErrorResponse
// @Router /project/{id} [put]
func (h *ProjectHandler) Update(c *gin.Context) {
	strId := c.Param("id")
	id, err := strconv.Atoi(strId)
	if err != nil {
		h.logger.Warn("Validation error", zap.Error(err))
		c.JSON(http.StatusBadRequest, ErrorResponse{Message: domain.ErrInvalidRequestBody.Error()})
		return
	}

	var req UpdateProjectDTO
	if err := c.ShouldBindJSON(&req); err != nil {
		h.logger.Warn("Validation error", zap.Error(err))
		c.JSON(http.StatusBadRequest, ErrorResponse{Message: domain.ErrInvalidRequestBody.Error()})
		return
	}

	project, err := h.pu.Update(c.Request.Context(), &domain.Project{
		Name: req.Name,
		ID:   uint(id),
	})
	if err != nil {
		h.logger.Error("Update project error", zap.Error(err))
		c.JSON(errorStatusCode(err), ErrorResponse{Message: err.Error()})
		return
	}

	response := ToProjectResponse(project)
	c.JSON(http.StatusOK, response)
}

// Delete godoc
// @Summary Удалить проект
// @Tags Project
// @Security BearerAuth
// @Produce json
// @Param id path int true "ID проекта"
// @Success 200
// @Failure 400 {object} ErrorResponse
// @Failure 404 {object} ErrorResponse
// @Router /project/{id} [delete]
func (h *ProjectHandler) Delete(c *gin.Context) {
	strId := c.Param("id")
	id, err := strconv.Atoi(strId)
	if err != nil {
		h.logger.Warn("Validation error", zap.Error(err))
		c.JSON(errorStatusCode(err), ErrorResponse{Message: domain.ErrInvalidRequestBody.Error()})
		return
	}

	if err := h.pu.Delete(c.Request.Context(), uint(id)); err != nil {
		h.logger.Error("Delete project error", zap.Error(err))
		c.JSON(errorStatusCode(err), ErrorResponse{Message: domain.ErrInvalidRequestBody.Error()})
		return
	}

	c.JSON(http.StatusOK, nil)
}

// Get godoc
// @Summary Получить все проекты
// @Tags project
// @Security BearerAuth
// @Produce json
// @Success 200 {object} ProjectResponse[]
// @Failure 400 {object} ErrorResponse
// @Router /project [get]
func (h *ProjectHandler) Get(c *gin.Context) {
	projects, err := h.pu.Get(c.Request.Context())
	if err != nil {
		h.logger.Error("Get projects error", zap.Error(err))
		c.JSON(errorStatusCode(err), ErrorResponse{Message: err.Error()})
		return
	}

	response := ToProjectsResponse(projects)

	c.JSON(http.StatusOK, response)
}

// GetByID Get godoc
// @Summary Получить проект по ID
// @Tags Project
// @Security BearerAuth
// @Produce json
// @Param id path int true "ID проекта"
// @Success 200 {object} ProjectResponse
// @Failure 400 {object} ErrorResponse
// @Router /project/{id} [get]
func (h *ProjectHandler) GetByID(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		h.logger.Warn("Validation error", zap.Error(err))
		c.JSON(http.StatusBadRequest, ErrorResponse{Message: domain.ErrInvalidRequestBody.Error()})
		return
	}

	project, err := h.pu.GetByID(c.Request.Context(), uint(id))
	if err != nil {
		h.logger.Error("Get project error", zap.Error(err))
		c.JSON(errorStatusCode(err), ErrorResponse{Message: domain.ErrInvalidRequestBody.Error()})
		return
	}

	response := ToProjectResponse(project)
	c.JSON(http.StatusOK, response)
}

func errorStatusCode(err error) int {
	switch {
	case errors.Is(err, domain.ErrProjectNotFound):
		return http.StatusNotFound
	case errors.Is(err, domain.ErrInvalidRequestBody):
		return http.StatusBadRequest
	default:
		return http.StatusInternalServerError
	}
}
