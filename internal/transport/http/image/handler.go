package image

import (
	"errors"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"net/http"
	"strconv"
	"thesis_back/internal/domain"
	"thesis_back/internal/usecase/image"
)

type ImageHandler struct {
	iu     image.IImageUseCase
	logger *zap.Logger
}

func NewImageHandler(iu image.IImageUseCase, logger *zap.Logger) *ImageHandler {
	return &ImageHandler{
		iu:     iu,
		logger: logger.Named("ImageHandler"),
	}
}

// Create UploadImage godoc
// @Summary Загрузка изображения
// @Description Загрузка изображения в MinIO
// @Tags Images
// @Accept multipart/form-data
// @Produce json
// @Param projectID formData integer true "ID проекта"
// @Param name formData string true "Название изображения"
// @Param image formData file true "Файл изображения"
// @Security BearerAuth
// @Success 200 {object} ImageResponse
// @Failure 400 {object} ErrorResponse
// @Router /image [post]
func (h *ImageHandler) Create(c *gin.Context) {
	var req AddImageDTO
	if err := c.ShouldBind(&req); err != nil {
		h.logger.Warn("Validation error", zap.Error(err))
		c.JSON(http.StatusBadRequest, gin.H{"error": domain.ErrInvalidRequestBody.Error()})
		return
	}

	image, err := h.iu.UploadImage(c.Request.Context(), req.Name, req.ImageFile, req.ProjectID)
	if err != nil {
		h.logger.Warn("Upload error", zap.Error(err))
		c.JSON(errorStatusCode(err), ErrorResponse{Message: err.Error()})
		return
	}

	response := ToImageResponse(image)
	c.JSON(http.StatusCreated, response)
}

// Update godoc
// @Summary Обновить название изображения
// @Tags Images
// @Security BearerAuth
// @Produce json
// @Param id path int true "ID изображения"
// @Param input body UpdateImageDTO true "Название картинки"
// @Success 200 {object} ImageResponse
// @Failure 400 {object} ErrorResponse
// @Router /image/{id} [put]
func (h *ImageHandler) Update(c *gin.Context) {
	strId := c.Param("id")
	id, err := strconv.Atoi(strId)
	if err != nil {
		h.logger.Warn("Validation error", zap.Error(err))
		c.JSON(http.StatusBadRequest, gin.H{"error": domain.ErrInvalidRequestBody.Error()})
		return
	}

	var req UpdateImageDTO
	if err := c.ShouldBindJSON(&req); err != nil {
		h.logger.Warn("Validation error", zap.Error(err))
		c.JSON(http.StatusBadRequest, gin.H{"error": domain.ErrInvalidRequestBody.Error()})
		return
	}

	image, err := h.iu.Update(c.Request.Context(), req.Name, uint(id))
	if err != nil {
		h.logger.Warn("Update error", zap.Error(err))
		c.JSON(errorStatusCode(err), ErrorResponse{Message: err.Error()})
		return
	}

	response := ToImageResponse(image)
	c.JSON(http.StatusOK, response)
}

// Delete godoc
// @Summary Удалить изображение
// @Tags Images
// @Security BearerAuth
// @Produce json
// @Param id path int true "ID изображения"
// @Success 200
// @Failure 400 {object} ErrorResponse
// @Failure 404 {object} ErrorResponse
// @Router /image/{id} [delete]
func (h *ImageHandler) Delete(c *gin.Context) {
	strId := c.Param("id")
	id, err := strconv.Atoi(strId)
	if err != nil {
		h.logger.Warn("Validation error", zap.Error(err))
		c.JSON(http.StatusBadRequest, gin.H{"error": domain.ErrInvalidRequestBody.Error()})
		return
	}

	if err := h.iu.Delete(c.Request.Context(), uint(id)); err != nil {
		h.logger.Warn("Delete error", zap.Error(err))
		c.JSON(errorStatusCode(err), ErrorResponse{Message: domain.ErrInvalidRequestBody.Error()})
		return
	}

	c.JSON(http.StatusOK, nil)
}

func errorStatusCode(err error) int {
	switch {
	case errors.Is(err, domain.ErrImageNotUploaded):
		return http.StatusBadRequest
	case errors.Is(err, domain.ErrImageNotFound):
		return http.StatusNotFound
	default:
		return http.StatusInternalServerError
	}
}
