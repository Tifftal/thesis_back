package layer

import (
	"encoding/json"
	"errors"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"net/http"
	"strconv"
	"thesis_back/internal/domain"
	"thesis_back/internal/usecase/layer"
)

type LayerHandler struct {
	lu     layer.ILayerUseCase
	logger *zap.Logger
}

func NewLayerHandler(lu layer.ILayerUseCase, logger *zap.Logger) *LayerHandler {
	return &LayerHandler{
		lu:     lu,
		logger: logger.Named("LayerHandler"),
	}
}

// Create godoc
// @Summary Создать новый слой
// @Tags Layers
// @Security BearerAuth
// @Produce json
// @Param input body CreateLayerDTO true "Название слоя"
// @Success 201 {object} LayerResponse
// @Failure 400 {object} ErrorResponse
// @Router /layer [post]
func (h *LayerHandler) Create(c *gin.Context) {
	var req CreateLayerDTO
	if err := c.ShouldBindJSON(&req); err != nil {
		h.logger.Warn("Binding failed", zap.Error(err))
		c.JSON(http.StatusBadRequest, ErrorResponse{Message: domain.ErrInvalidCredentials.Error()})
		return
	}

	layer, err := h.lu.Create(c.Request.Context(), &domain.Layer{
		ImageID: req.ImageID,
		Name:    req.Name,
	})
	if err != nil {
		h.logger.Warn("Layer creation failed", zap.Error(err))
		c.JSON(errorStatusCode(err), ErrorResponse{
			Message: err.Error(),
		})
	}

	response := ToLayerResponse(layer)
	c.JSON(http.StatusCreated, response)
}

// Update godoc
// @Summary Обновить слой
// @Tags Layers
// @Security BearerAuth
// @Produce json
// @Param id path int true "ID слоя"
// @Param input body UpdateLayerDTO true "Название слоя"
// @Success 200 {object} LayerResponse
// @Failure 400 {object} ErrorResponse
// @Router /layer/{id} [put]
func (h *LayerHandler) Update(c *gin.Context) {
	strId := c.Param("id")
	id, err := strconv.Atoi(strId)
	if err != nil {
		h.logger.Warn("Layer id conversion failed", zap.Error(err))
		c.JSON(http.StatusBadRequest, ErrorResponse{Message: domain.ErrInvalidRequestBody.Error()})
		return
	}

	var req UpdateLayerDTO
	if err := c.ShouldBindJSON(&req); err != nil {
		h.logger.Warn("Binding failed", zap.Error(err))
		c.JSON(http.StatusBadRequest, ErrorResponse{Message: domain.ErrInvalidRequestBody.Error()})
		return
	}

	measurementsJSON, err := json.Marshal(req.Measurements)
	if err != nil {
		c.JSON(400, gin.H{"error": "Invalid measurements format"})
		return
	}

	layer, err := h.lu.Update(c.Request.Context(), &domain.Layer{
		ID:           uint(id),
		Name:         req.Name,
		Measurements: measurementsJSON,
	})
	if err != nil {
		h.logger.Warn("Layer update failed", zap.Error(err))
		c.JSON(errorStatusCode(err), ErrorResponse{Message: err.Error()})
		return
	}

	response := ToLayerResponse(layer)
	c.JSON(http.StatusOK, response)
}

// Delete godoc
// @Summary Удалить слой
// @Tags Layers
// @Security BearerAuth
// @Produce json
// @Param id path int true "ID слоя"
// @Success 200
// @Failure 400 {object} ErrorResponse
// @Failure 404 {object} ErrorResponse
// @Router /layer/{id} [delete]
func (h *LayerHandler) Delete(c *gin.Context) {
	strId := c.Param("id")
	id, err := strconv.Atoi(strId)
	if err != nil {
		h.logger.Warn("Layer id conversion failed", zap.Error(err))
		c.JSON(http.StatusBadRequest, ErrorResponse{Message: domain.ErrInvalidRequestBody.Error()})
		return
	}

	if err := h.lu.Delete(c.Request.Context(), uint(id)); err != nil {
		h.logger.Warn("Layer delete failed", zap.Error(err))
		c.JSON(errorStatusCode(err), ErrorResponse{
			Message: domain.ErrInvalidRequestBody.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, nil)
}

func errorStatusCode(err error) int {
	switch {
	case errors.Is(err, domain.ErrLayerNotFound):
		return http.StatusNotFound
	case errors.Is(err, domain.ErrInvalidRequestBody):
		return http.StatusBadRequest
	default:
		return http.StatusInternalServerError
	}
}
