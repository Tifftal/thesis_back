package image

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"thesis_back/internal/usecase/image"
)

type ImageHandler struct {
	iu     *image.IImageUseCase
	logger *zap.Logger
}

func NewImageHandler(iu *image.IImageUseCase, logger *zap.Logger) *ImageHandler {
	return &ImageHandler{
		iu:     iu,
		logger: logger.Named("ImageHandler"),
	}
}

func (h *ImageHandler) Create(c *gin.Context) {

}

func (h *ImageHandler) Get(c *gin.Context) {

}

func (h *ImageHandler) GetByID(c *gin.Context) {

}

func (h *ImageHandler) Update(c *gin.Context) {

}

func (h *ImageHandler) Delete(c *gin.Context) {
	
}
